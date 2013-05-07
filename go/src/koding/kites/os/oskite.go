package main

import (
	"errors"
	"io/ioutil"
	"koding/kites/os/ldapserver"
	"koding/tools/config"
	"koding/tools/db"
	"koding/tools/dnode"
	"koding/tools/kite"
	"koding/tools/lifecycle"
	"koding/tools/log"
	"koding/tools/utils"
	"koding/virt"
	"labix.org/v2/mgo/bson"
	"net"
	"strings"
	"sync"
	"time"
)

type VMInfo struct {
	vmId          bson.ObjectId
	channels      map[*kite.Channel]bool
	timeout       *time.Timer
	totalCpuUsage int

	State       string `json:"state"`
	CpuUsage    int    `json:"cpuUsage"`
	CpuShares   int    `json:"cpuShares"`
	MemoryUsage int    `json:"memoryUsage"`
	MemoryLimit int    `json:"memoryLimit"`
}

var infos = make(map[bson.ObjectId]*VMInfo)
var infosMutex sync.Mutex

func main() {
	lifecycle.Startup("kite.os", true)
	virt.LoadTemplates(config.Current.ProjectRoot + "/go/templates")

	dirs, err := ioutil.ReadDir("/var/lib/lxc")
	if err != nil {
		panic(err)
	}
	for _, dir := range dirs {
		if strings.HasPrefix(dir.Name(), "vm-") {
			vm := virt.VM{Id: bson.ObjectIdHex(dir.Name()[3:])}
			if err := vm.Unprepare(); err != nil {
				log.Warn(err.Error())
			}
		}
	}

	go ldapserver.Listen()
	go LimiterLoop()
	k := kite.New("os")

	k.LoadBalancer = func(correlationName string, username string, deadService string) string {
		if deadService != "" {
			if _, err := db.VMs.UpdateAll(bson.M{"hostKite": deadService}, bson.M{"$set": bson.M{"hostKite": nil}}); err != nil {
				log.LogError(err, 0)
			}
		}

		var vm *virt.VM
		if bson.IsObjectIdHex(correlationName) {
			db.VMs.FindId(bson.ObjectIdHex(correlationName)).One(&vm)
		}
		if vm == nil {
			if err := db.VMs.Find(bson.M{"name": correlationName}).One(&vm); err != nil {
				return k.ServiceUniqueName
			}
		}

		if vm.HostKite == "" {
			return k.ServiceUniqueName
		}
		return vm.HostKite
	}

	registerVmMethod(k, "vm.start", false, func(args *dnode.Partial, channel *kite.Channel, user *virt.User, vm *virt.VM, vos *virt.VOS) (interface{}, error) {
		userEntry := vm.GetUserEntry(user)
		if userEntry == nil || !userEntry.Sudo {
			return nil, errors.New("Permission denied.")
		}

		return vm.Start()
	})

	registerVmMethod(k, "vm.shutdown", false, func(args *dnode.Partial, channel *kite.Channel, user *virt.User, vm *virt.VM, vos *virt.VOS) (interface{}, error) {
		userEntry := vm.GetUserEntry(user)
		if userEntry == nil || !userEntry.Sudo {
			return nil, errors.New("Permission denied.")
		}

		return vm.Shutdown()
	})

	registerVmMethod(k, "vm.stop", false, func(args *dnode.Partial, channel *kite.Channel, user *virt.User, vm *virt.VM, vos *virt.VOS) (interface{}, error) {
		userEntry := vm.GetUserEntry(user)
		if userEntry == nil || !userEntry.Sudo {
			return nil, errors.New("Permission denied.")
		}

		return vm.Stop()
	})

	registerVmMethod(k, "vm.reinitialize", false, func(args *dnode.Partial, channel *kite.Channel, user *virt.User, vm *virt.VM, vos *virt.VOS) (interface{}, error) {
		userEntry := vm.GetUserEntry(user)
		if userEntry == nil || !userEntry.Sudo {
			return nil, errors.New("Permission denied.")
		}

		vm.Prepare(getUsers(vm), true)
		return vm.Start()
	})

	registerVmMethod(k, "vm.info", false, func(args *dnode.Partial, channel *kite.Channel, user *virt.User, vm *virt.VM, vos *virt.VOS) (interface{}, error) {
		userEntry := vm.GetUserEntry(user)
		if userEntry == nil {
			return nil, errors.New("Permission denied.")
		}

		info := infos[vm.Id]
		info.State = vm.GetState()
		return info, nil
	})

	registerVmMethod(k, "vm.createSnapshot", false, func(args *dnode.Partial, channel *kite.Channel, user *virt.User, vm *virt.VM, vos *virt.VOS) (interface{}, error) {
		userEntry := vm.GetUserEntry(user)
		if userEntry == nil || !userEntry.Sudo {
			return nil, errors.New("Permission denied.")
		}

		snippet := virt.VM{
			Id:         bson.NewObjectId(),
			SnapshotOf: vm.Id,
		}

		if err := vm.CreateConsistentSnapshot(snippet.Id.Hex()); err != nil {
			return nil, err
		}

		if err := db.VMs.Insert(snippet); err != nil {
			return nil, err
		}

		return snippet.Id.Hex(), nil
	})

	registerVmMethod(k, "spawn", true, func(args *dnode.Partial, channel *kite.Channel, user *virt.User, vm *virt.VM, vos *virt.VOS) (interface{}, error) {
		userEntry := vm.GetUserEntry(user)
		if userEntry == nil {
			return nil, errors.New("Permission denied.")
		}

		var command []string
		if args.Unmarshal(&command) != nil {
			return nil, &kite.ArgumentError{Expected: "array of strings"}
		}
		return vm.AttachCommand(user.Uid, "", command...).CombinedOutput()
	})

	registerVmMethod(k, "exec", true, func(args *dnode.Partial, channel *kite.Channel, user *virt.User, vm *virt.VM, vos *virt.VOS) (interface{}, error) {
		userEntry := vm.GetUserEntry(user)
		if userEntry == nil {
			return nil, errors.New("Permission denied.")
		}

		var line string
		if args.Unmarshal(&line) != nil {
			return nil, &kite.ArgumentError{Expected: "string"}
		}
		return vm.AttachCommand(user.Uid, "", "/bin/bash", "-c", line).CombinedOutput()
	})

	registerFileSystemMethods(k)
	registerWebtermMethods(k)
	registerAppMethods(k)

	k.Run()
}

func registerVmMethod(k *kite.Kite, method string, concurrent bool, callback func(*dnode.Partial, *kite.Channel, *virt.User, *virt.VM, *virt.VOS) (interface{}, error)) {
	k.Handle(method, concurrent, func(args *dnode.Partial, channel *kite.Channel) (interface{}, error) {
		var user virt.User
		if err := db.Users.Find(bson.M{"username": channel.Username}).One(&user); err != nil {
			panic(err)
		}
		if user.Uid < virt.UserIdOffset {
			panic("User with too low uid.")
		}

		vm, _ := channel.KiteData.(*virt.VM)
		if vm == nil {
			if bson.IsObjectIdHex(channel.CorrelationName) {
				db.VMs.FindId(bson.ObjectIdHex(channel.CorrelationName)).One(&vm)
			}
			if vm == nil {
				if err := db.VMs.Find(bson.M{"name": channel.CorrelationName}).One(&vm); err != nil {
					return nil, errors.New("There is no VM with name/id '" + channel.CorrelationName + "'.")
				}
			}

			if vm.HostKite != k.ServiceUniqueName {
				if err := db.VMs.Update(bson.M{"_id": vm.Id, "hostKite": nil}, bson.M{"$set": bson.M{"hostKite": k.ServiceUniqueName}}); err != nil {
					return nil, &kite.WrongChannelError{}
				}
				vm.HostKite = k.ServiceUniqueName
			}

			if vm.SnapshotOf != "" {
				var err error
				vm, err = vm.CreateTemporaryVM()
				if err != nil {
					return nil, err
				}
			}

			channel.KiteData = vm
		}

		infosMutex.Lock()
		info, isExistingState := infos[vm.Id]
		if !isExistingState {
			info = newInfo(vm)
			infos[vm.Id] = info
		}
		if !info.channels[channel] {
			info.channels[channel] = true
			if info.timeout != nil {
				info.timeout.Stop()
				info.timeout = nil
			}

			channel.OnDisconnect(func() {
				infosMutex.Lock()
				defer infosMutex.Unlock()

				delete(info.channels, channel)
				if len(info.channels) == 0 {
					info.startTimeout()
				}
			})
		}
		infosMutex.Unlock()

		if !isExistingState {
			if vm.IP == nil {
				ipInt := db.NextCounterValue("vm_ip")
				ip := net.IPv4(byte(ipInt>>24), byte(ipInt>>16), byte(ipInt>>8), byte(ipInt))
				if !vm.IsTemporary() {
					if err := db.VMs.Update(bson.M{"_id": vm.Id, "ip": nil}, bson.M{"$set": bson.M{"ip": ip}}); err != nil {
						panic(err)
					}
				}
				vm.IP = ip
			}
			if vm.LdapPassword == "" {
				ldapPassword := utils.RandomString()
				if !vm.IsTemporary() {
					if err := db.VMs.Update(bson.M{"_id": vm.Id}, bson.M{"$set": bson.M{"ldapPassword": ldapPassword}}); err != nil {
						panic(err)
					}
				}
				vm.LdapPassword = ldapPassword
			}

			vm.Prepare(getUsers(vm), false)
			if out, err := vm.Start(); err != nil {
				log.Err("Could not start VM.", err, out)
			}
			if out, err := vm.WaitForState("RUNNING", time.Second); err != nil {
				log.Warn("Waiting for VM startup failed.", err, out)
			}
		}

		return callback(args, channel, &user, vm, vm.OS(&user))
	})
}

func getUsers(vm *virt.VM) []virt.User {
	users := make([]virt.User, len(vm.Users))
	for i, entry := range vm.Users {
		if err := db.Users.FindId(entry.Id).One(&users[i]); err != nil {
			panic(err)
		}
		if users[i].Uid == 0 {
			panic("User with uid 0.")
		}
	}
	return users
}

func newInfo(vm *virt.VM) *VMInfo {
	return &VMInfo{
		vmId:          vm.Id,
		channels:      make(map[*kite.Channel]bool),
		totalCpuUsage: utils.MaxInt,
		CpuShares:     1000,
	}
}

func (info *VMInfo) startTimeout() {
	info.timeout = time.AfterFunc(10*time.Minute, func() {
		infosMutex.Lock()
		defer infosMutex.Unlock()

		if len(info.channels) != 0 {
			return
		}

		var vm virt.VM
		if err := db.VMs.FindId(info.vmId).One(&vm); err != nil {
			log.Err("Could not find VM for shutdown.", err)
		}
		if out, err := vm.Shutdown(); err != nil {
			log.Err("Could not shutdown VM.", err, out)
		}

		if err := vm.Unprepare(); err != nil {
			log.Warn(err.Error())
		}

		if !vm.IsTemporary() {
			if err := db.VMs.Update(bson.M{"_id": vm.Id}, bson.M{"$set": bson.M{"hostKite": nil}}); err != nil {
				log.LogError(err, 0)
			}
		}

		if vm.IsTemporary() {
			if err := vm.Destroy(); err != nil {
				log.Warn(err.Error())
			}
		}

		delete(infos, vm.Id)
	})
}

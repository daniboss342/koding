// Code generated by go-bindata.
// sources:
// provision.sh
// DO NOT EDIT!

package metadata

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _provisionSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x55\x7f\x6f\xdb\x36\x10\xfd\x5b\xfc\x14\x57\x27\xd8\x5a\x6c\x14\xed\x00\x41\x80\x0e\x0b\xb0\xc5\x42\x11\xa4\x75\x81\xfc\x00\x0a\x6c\x83\x40\x4b\x27\x8b\x33\x25\x72\x24\xe5\xd8\xc9\xfc\xdd\x07\x4a\xb2\xad\xc8\x4a\x36\xcc\xff\x99\xf7\xde\xbb\xe3\xf1\xdd\xe9\xe4\x1d\x9b\x8b\x92\xcd\xb9\xcd\x09\x39\x81\x1b\x95\x8a\x72\x01\x5a\x59\x47\xb5\x51\x2b\x61\x85\x2a\xc1\x26\x46\x68\x07\x06\xad\x56\xa5\x15\x73\x89\x90\x29\x03\x29\x6a\xa9\x36\x1e\xcf\x61\x29\x05\x96\x0e\x2c\x9a\x95\x48\x30\x24\x27\xe4\x04\xae\x94\xde\x18\xb1\xc8\x1d\xbc\xbf\xfa\x00\x67\xe3\xc9\x19\x3d\x1b\x4f\x2e\x76\x49\xae\xcb\x24\xfc\x11\xb8\x94\x50\x83\xac\xd7\x47\xb3\xc2\x34\x24\xc4\xa2\x03\x8a\x95\x02\x2d\x34\x66\x5c\x48\x5f\xdc\x7d\xee\xf3\x4a\xa9\x1e\x3d\x7d\xc5\x8d\xe0\x73\x89\x16\xb8\x41\xd0\xdc\x5a\x4c\x61\x25\x38\x38\x34\x86\x67\xca\x14\xdf\xdb\x3d\x08\xe6\x52\x25\xcb\x90\xe0\x5a\x2b\xe3\xe0\xe6\xeb\xf4\x7a\xf6\x29\x7e\xb8\x8b\x6e\x67\xbf\x7c\x89\x7e\x1e\x9d\x9e\x3e\xf7\xce\x3e\xd2\xd3\xe7\x15\x37\xe1\xb2\x2e\x36\xe6\x49\xa2\xaa\xd2\xc5\xda\xa8\x4c\x48\x8c\x4b\x91\x2c\x4b\x5e\xe0\x76\x3b\xda\x8b\x7e\xbe\x8e\x66\xf7\xf1\xc3\xed\xe7\x46\x6f\xff\xb7\x27\xd5\xb4\x2a\xae\x8c\xec\x90\xef\xae\x6e\xa3\x68\xb6\x27\x1f\xfe\xf6\xc8\x36\x31\x88\x65\x8f\xfc\x70\x17\xc5\xd1\x97\x5f\xa3\xe9\x34\x9a\xd6\xf4\xee\x41\x4f\xa0\xb2\x18\x63\x31\xc7\x34\xc5\xd4\x4b\x10\x67\xb8\x86\x11\x26\xb9\x82\xf8\x66\x1a\x4f\xbf\xce\xa2\x18\xfe\x06\x87\x08\x94\x03\x5b\x71\xc3\xa4\x5a\xb0\x44\xaa\x2a\xa5\xa2\x14\x8e\xaa\xca\xe9\xca\x85\x52\x2d\x46\x10\x7d\xbb\xbe\x27\x44\x94\xd6\x71\x29\xdb\xea\xde\x7f\x80\x67\x12\x24\x95\x91\x40\xa9\x54\x09\x77\xde\x44\x94\x5a\x21\xbd\x47\x28\xb5\xb9\x7a\xa4\x68\x8c\x32\x40\xa9\x41\x67\x36\x70\x0e\x2f\xaf\xbd\x1d\x01\x6d\x13\x01\x73\x85\x66\x8d\x74\xe8\xb8\x09\x17\x4f\x04\x00\xc0\x71\x03\xf4\x0a\x18\xd0\x75\x36\x84\x21\xc1\xc2\xa8\x4a\xf3\x34\x05\x4a\x33\x65\x12\x84\x06\x40\x82\xca\xa2\x29\x54\xea\x2f\x48\x3f\xb5\xa7\x30\x60\x02\xdf\x9f\x40\x64\xf0\x1b\xbc\x03\xba\x06\x56\x59\x53\x8f\x4b\xcb\xf8\xe3\x27\x70\xb9\xd7\x0b\x64\x09\xd4\x66\xc0\x94\x76\x6c\x29\x1c\xb2\xe6\x91\xd9\xae\xd3\x5d\x56\x4f\x85\x04\x99\xe8\x64\xc1\x26\x6e\x73\x6e\x90\x39\x34\x85\x28\x33\xd5\xc9\x54\x2c\x53\x61\x80\xea\x0e\xec\xbf\xe4\xef\xe9\x0d\xe4\xe8\xa8\x1c\x07\xd9\xb7\xc1\xd3\xf5\x71\xf1\xde\x2f\xa6\x1a\x68\xd1\xa1\xf0\x97\x10\x12\x04\x49\xee\x1f\x63\x7c\x71\x7e\x3e\x18\x54\x8f\x25\x18\xa5\xdc\xc7\x5d\x03\xfb\x98\x4c\x90\x2d\x21\x05\x17\xad\xf5\x6a\x33\x8f\x26\x67\x17\xe1\x38\x1c\x87\x13\x18\x7a\x58\xb8\xbc\x04\x86\x2e\x61\xb9\xb2\xce\x12\x12\x38\x55\x25\xf9\xc1\xef\x4d\x03\xbd\xc9\xe1\x77\x12\x40\xfd\x7b\x7b\x18\x9a\x3e\xbc\x03\xb7\xd1\x3b\xa7\xc1\x77\x97\x2c\xc5\x15\x2b\x2b\x29\xf7\x8d\xa8\x9b\xd5\x1b\xd2\x2d\x50\xfc\x0b\x26\x9d\x76\x05\x2f\x67\x8a\x04\x01\x4a\x91\x35\xe2\x9b\xaa\x18\x54\x0e\x7c\xa0\xe5\x01\xa5\xdc\xda\xaa\xc0\x0d\x5a\x18\xd0\xe0\xda\xd1\x05\xba\x61\x9d\x5d\x70\xaf\xb5\xe9\x4a\x58\x1c\xac\x2f\x13\xad\x1b\x0e\x2f\xdd\xb3\xe3\xff\x5b\x0a\x87\x45\x7a\xb4\x14\xda\x47\x5a\x3c\x91\x60\xf1\x24\x34\x50\x9a\x62\xa2\x0a\x6d\xd0\xda\xfd\xd0\x53\x6a\x5d\xaa\x8e\x28\x70\x79\x3c\x2e\xfb\x32\x6b\x43\xfe\xb0\x7e\x15\x41\x5a\x5b\xd2\xdb\xe1\xb5\xe1\x3d\x13\xf4\xc9\xcd\xe1\xb0\xc1\x82\x7f\xf5\xd6\x2b\xa5\x00\x2d\xd0\xf1\x94\x3b\x4e\xfd\x56\x7b\xa5\x9c\x03\xc8\x7f\xbd\x5a\x23\x8b\x39\x6b\x3e\x0a\x6c\x17\x0d\xff\xb4\xaa\x04\x53\x95\xbb\x99\xf6\x4b\xaf\x87\xf5\x59\x68\x0d\xb6\x79\xc7\xaf\x6f\xc1\xba\x13\xea\x0d\xf7\xf6\x55\xff\x09\x00\x00\xff\xff\xbf\x46\xc9\xfe\x93\x08\x00\x00")

func provisionShBytes() ([]byte, error) {
	return bindataRead(
		_provisionSh,
		"provision.sh",
	)
}

func provisionSh() (*asset, error) {
	bytes, err := provisionShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "provision.sh", size: 2195, mode: os.FileMode(420), modTime: time.Unix(1470666525, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"provision.sh": provisionSh,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"provision.sh": {provisionSh, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

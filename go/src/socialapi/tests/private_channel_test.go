package main

import (
	"encoding/json"
	mongomodels "koding/db/models"
	"koding/db/mongodb/modelhelper"
	"math/rand"
	"socialapi/config"
	"socialapi/models"
	"socialapi/request"
	"socialapi/rest"
	"strconv"
	"testing"

	"github.com/koding/runner"
	"labix.org/v2/mgo/bson"

	. "github.com/smartystreets/goconvey/convey"
)

func CreatePrivateChannelUser(nickname string) {
	_, err := modelhelper.GetAccount(nickname)
	if err == nil {
		return
	}

	if err != modelhelper.ErrNotFound {
		panic(err)
	}

	acc := new(mongomodels.Account)
	acc.Id = bson.NewObjectId()
	acc.Profile.Nickname = nickname

	modelhelper.CreateAccount(acc)

	a := models.NewAccount()
	a.Nick = nickname
	a.OldId = acc.Id.Hex()
	a.Create()
}

func TestPrivateMesssages(t *testing.T) {
	r := runner.New("test")
	if err := r.Init(); err != nil {
		t.Fatalf("couldnt start bongo %s", err.Error())
	}
	defer r.Close()

	appConfig := config.MustRead(r.Conf.Path)
	modelhelper.Initialize(appConfig.Mongo)
	defer modelhelper.Close()

	CreatePrivateChannelUser("devrim")
	CreatePrivateChannelUser("sinan")
	CreatePrivateChannelUser("chris")

	Convey("while testing private channels", t, func() {
		account := models.NewAccount()
		account.OldId = AccountOldId.Hex()
		account, err := rest.CreateAccount(account)
		So(err, ShouldBeNil)
		So(account, ShouldNotBeNil)

		recipient := models.NewAccount()
		recipient.OldId = AccountOldId2.Hex()
		recipient, err = rest.CreateAccount(recipient)
		So(err, ShouldBeNil)
		So(recipient, ShouldNotBeNil)

		recipient2 := models.NewAccount()
		recipient2.OldId = AccountOldId3.Hex()
		recipient2, err = rest.CreateAccount(recipient2)
		So(err, ShouldBeNil)
		So(recipient2, ShouldNotBeNil)

		groupName := "testgroup" + strconv.FormatInt(rand.Int63(), 10)

		Convey("one can send private message to one person", func() {
			pmr := models.PrivateChannelRequest{}
			pmr.AccountId = account.Id
			pmr.Body = "this is a body message for private message @chris @devrim @sinan"
			pmr.GroupName = groupName
			pmr.Recipients = []string{"chris", "devrim", "sinan"}
			cmc, err := rest.SendPrivateChannelRequest(pmr)
			So(err, ShouldBeNil)
			So(cmc, ShouldNotBeNil)

		})

		Convey("0 recipient should not fail", func() {
			pmr := models.PrivateChannelRequest{}
			pmr.AccountId = account.Id
			pmr.Body = "this is a body for private message"
			pmr.GroupName = groupName
			pmr.Recipients = []string{}

			cmc, err := rest.SendPrivateChannelRequest(pmr)
			So(err, ShouldBeNil)
			So(cmc, ShouldNotBeNil)

		})
		Convey("if group name is nil, should not fail to create PM", func() {
			pmr := models.PrivateChannelRequest{}
			pmr.AccountId = account.Id
			pmr.Body = "this is a body for private message @chris @devrim @sinan"
			pmr.GroupName = ""
			pmr.Recipients = []string{"chris", "devrim", "sinan"}

			cmc, err := rest.SendPrivateChannelRequest(pmr)
			So(err, ShouldBeNil)
			So(cmc, ShouldNotBeNil)
		})

		Convey("if sender is not defined should fail to create PM", func() {
			pmr := models.PrivateChannelRequest{}
			pmr.AccountId = 0
			pmr.Body = "this is a body for private message"
			pmr.GroupName = ""
			pmr.Recipients = []string{}

			cmc, err := rest.SendPrivateChannelRequest(pmr)
			So(err, ShouldNotBeNil)
			So(cmc, ShouldBeNil)
		})

		Convey("one can send private message to multiple person", func() {
			pmr := models.PrivateChannelRequest{}
			pmr.AccountId = account.Id
			pmr.Body = "this is a body for private message @sinan"
			pmr.GroupName = groupName
			pmr.Recipients = []string{"sinan"}
			cmc, err := rest.SendPrivateChannelRequest(pmr)
			So(err, ShouldBeNil)
			So(cmc, ShouldNotBeNil)

		})
		Convey("private message response should have created channel", func() {
			pmr := models.PrivateChannelRequest{}
			pmr.AccountId = account.Id
			pmr.Body = "this is a body for private message @devrim @sinan"
			pmr.GroupName = groupName
			pmr.Recipients = []string{"devrim", "sinan"}

			cmc, err := rest.SendPrivateChannelRequest(pmr)
			So(err, ShouldBeNil)
			So(cmc, ShouldNotBeNil)
			So(cmc.Channel.TypeConstant, ShouldEqual, models.Channel_TYPE_PRIVATE_MESSAGE)
			So(cmc.Channel.Id, ShouldBeGreaterThan, 0)
			So(cmc.Channel.GroupName, ShouldEqual, groupName)
			So(cmc.Channel.PrivacyConstant, ShouldEqual, models.Channel_PRIVACY_PRIVATE)

		})

		Convey("private message response should have participant status data", func() {
			pmr := models.PrivateChannelRequest{}
			pmr.AccountId = account.Id
			pmr.Body = "this is a body for private message @chris @devrim @sinan"
			pmr.GroupName = groupName
			pmr.Recipients = []string{"chris", "devrim", "sinan"}

			cmc, err := rest.SendPrivateChannelRequest(pmr)
			So(err, ShouldBeNil)
			So(cmc, ShouldNotBeNil)
			So(cmc.IsParticipant, ShouldBeTrue)
		})

		Convey("private message response should have participant count", func() {
			pmr := models.PrivateChannelRequest{}
			pmr.AccountId = account.Id
			pmr.Body = "this is a body for @sinan private message @devrim"
			pmr.GroupName = groupName
			pmr.Recipients = []string{"devrim", "sinan"}
			cmc, err := rest.SendPrivateChannelRequest(pmr)
			So(err, ShouldBeNil)
			So(cmc, ShouldNotBeNil)
			So(cmc.ParticipantCount, ShouldEqual, 3)
		})

		Convey("private message response should have participant preview", func() {
			pmr := models.PrivateChannelRequest{}
			pmr.AccountId = account.Id
			pmr.Body = "this is @chris a body for @devrim private message"
			pmr.GroupName = groupName
			pmr.Recipients = []string{"chris", "devrim"}
			cmc, err := rest.SendPrivateChannelRequest(pmr)
			So(err, ShouldBeNil)
			So(cmc, ShouldNotBeNil)
			So(len(cmc.ParticipantsPreview), ShouldEqual, 3)
		})

		Convey("private message response should have last Message", func() {
			body := "hi @devrim this is a body for private message also for @chris"
			pmr := models.PrivateChannelRequest{}
			pmr.AccountId = account.Id
			pmr.Body = body
			pmr.GroupName = groupName
			pmr.Recipients = []string{"chris", "devrim"}
			cmc, err := rest.SendPrivateChannelRequest(pmr)
			So(err, ShouldBeNil)
			So(cmc, ShouldNotBeNil)
			So(cmc.LastMessage.Message.Body, ShouldEqual, body)
		})

		Convey("private message should be listed by all recipients", func() {
			// use a different group name
			// in order not to interfere with another request
			groupName := "testgroup" + strconv.FormatInt(rand.Int63(), 10)

			body := "hi @devrim this is a body for private message also for @chris"
			pmr := models.PrivateChannelRequest{}
			pmr.AccountId = account.Id
			pmr.Body = body
			pmr.GroupName = groupName
			pmr.Recipients = []string{"chris", "devrim"}
			cmc, err := rest.SendPrivateChannelRequest(pmr)
			So(err, ShouldBeNil)
			So(cmc, ShouldNotBeNil)

			pm, err := rest.GetPrivateChannels(&request.Query{AccountId: account.Id, GroupName: groupName})
			So(err, ShouldBeNil)
			So(pm, ShouldNotBeNil)
			So(len(pm), ShouldNotEqual, 0)
			So(pm[0], ShouldNotBeNil)
			So(pm[0].Channel.TypeConstant, ShouldEqual, models.Channel_TYPE_PRIVATE_MESSAGE)
			So(pm[0].Channel.Id, ShouldEqual, cmc.Channel.Id)
			So(pm[0].Channel.GroupName, ShouldEqual, cmc.Channel.GroupName)
			So(pm[0].LastMessage.Message.Body, ShouldEqual, cmc.LastMessage.Message.Body)
			So(pm[0].Channel.PrivacyConstant, ShouldEqual, models.Channel_PRIVACY_PRIVATE)
			So(len(pm[0].ParticipantsPreview), ShouldEqual, 3)
			So(pm[0].IsParticipant, ShouldBeTrue)

		})

		Convey("user should be able to search private messages via purpose field", func() {
			groupName := "testgroup" + strconv.FormatInt(rand.Int63(), 10)

			pmr := models.PrivateChannelRequest{}
			pmr.AccountId = account.Id
			pmr.Body = "search private messages"
			pmr.GroupName = groupName
			pmr.Recipients = []string{"chris", "devrim"}
			pmr.Purpose = "test me up"

			cmc, err := rest.SendPrivateChannelRequest(pmr)
			So(err, ShouldBeNil)

			query := request.Query{AccountId: account.Id, GroupName: groupName}
			_, err = rest.SearchPrivateChannels(&query)
			So(err, ShouldNotBeNil)

			query.Name = "test"
			pm, err := rest.SearchPrivateChannels(&query)
			So(err, ShouldBeNil)
			So(pm, ShouldNotBeNil)
			So(len(pm), ShouldNotEqual, 0)
			So(pm[0], ShouldNotBeNil)
			So(pm[0].Channel.TypeConstant, ShouldEqual, models.Channel_TYPE_PRIVATE_MESSAGE)
			So(pm[0].Channel.Id, ShouldEqual, cmc.Channel.Id)
			So(pm[0].Channel.GroupName, ShouldEqual, cmc.Channel.GroupName)
			So(pm[0].LastMessage.Message.Body, ShouldEqual, cmc.LastMessage.Message.Body)
			So(pm[0].Channel.PrivacyConstant, ShouldEqual, models.Channel_PRIVACY_PRIVATE)
			So(pm[0].IsParticipant, ShouldBeTrue)

		})

		Convey("user join activity should be listed by recipients", func() {
			groupName := "testgroup" + strconv.FormatInt(rand.Int63(), 10)

			pmr := models.PrivateChannelRequest{}
			pmr.AccountId = account.Id
			pmr.Body = "test private message participants"
			pmr.GroupName = groupName
			pmr.Recipients = []string{"chris", "devrim"}

			cc, err := rest.SendPrivateChannelRequest(pmr)

			So(err, ShouldBeNil)
			So(cc, ShouldNotBeNil)

			ses, err := models.FetchOrCreateSession(account.Nick, groupName)
			So(err, ShouldBeNil)
			So(ses, ShouldNotBeNil)

			history, err := rest.GetHistory(
				cc.Channel.Id,
				&request.Query{
					AccountId: account.Id,
				},
				ses.ClientId,
			)

			So(err, ShouldBeNil)
			So(history, ShouldNotBeNil)
			So(len(history.MessageList), ShouldEqual, 2)

			// add participant
			_, err = rest.AddChannelParticipant(cc.Channel.Id, account.Id, recipient.Id)
			So(err, ShouldBeNil)

			history, err = rest.GetHistory(
				cc.Channel.Id,
				&request.Query{
					AccountId: account.Id,
				},
				ses.ClientId,
			)

			So(err, ShouldBeNil)
			So(history, ShouldNotBeNil)
			So(len(history.MessageList), ShouldEqual, 3)

			So(history.MessageList[0].Message, ShouldNotBeNil)
			So(history.MessageList[0].Message.TypeConstant, ShouldEqual, models.ChannelMessage_TYPE_SYSTEM)
			So(history.MessageList[0].Message.Payload, ShouldNotBeNil)
			addedBy, ok := history.MessageList[0].Message.Payload["addedBy"]
			So(ok, ShouldBeTrue)
			So(*addedBy, ShouldEqual, account.OldId)

			systemType, ok := history.MessageList[0].Message.Payload["systemType"]
			So(ok, ShouldBeTrue)
			So(*systemType, ShouldEqual, models.PrivateMessageSystem_TYPE_JOIN)

			// try to add same participant
			_, err = rest.AddChannelParticipant(cc.Channel.Id, account.Id, recipient.Id)
			So(err, ShouldBeNil)

			history, err = rest.GetHistory(
				cc.Channel.Id,
				&request.Query{
					AccountId: account.Id,
				},
				ses.ClientId,
			)

			So(err, ShouldBeNil)
			So(history, ShouldNotBeNil)
			So(len(history.MessageList), ShouldEqual, 3)

		})

		Convey("user should not be able to edit join messages", func() {
			groupName := "testgroup" + strconv.FormatInt(rand.Int63(), 10)

			pmr := models.PrivateChannelRequest{}
			pmr.AccountId = account.Id
			pmr.Body = "test private message participants again"
			pmr.GroupName = groupName
			pmr.Recipients = []string{"chris"}

			cc, err := rest.SendPrivateChannelRequest(pmr)
			So(err, ShouldBeNil)
			So(cc, ShouldNotBeNil)

			_, err = rest.AddChannelParticipant(cc.Channel.Id, account.Id, recipient.Id)
			So(err, ShouldBeNil)

			ses, err := models.FetchOrCreateSession(account.Nick, groupName)
			So(err, ShouldBeNil)
			So(ses, ShouldNotBeNil)

			history, err := rest.GetHistory(
				cc.Channel.Id,
				&request.Query{
					AccountId: account.Id,
				},
				ses.ClientId,
			)

			So(err, ShouldBeNil)
			So(history, ShouldNotBeNil)
			So(len(history.MessageList), ShouldEqual, 3)

			joinMessage := history.MessageList[0].Message
			So(joinMessage, ShouldNotBeNil)

			_, err = rest.UpdatePost(joinMessage)
			So(err, ShouldNotBeNil)
		})

		Convey("first chat message should include initial participants", func() {
			groupName := "testgroup" + strconv.FormatInt(rand.Int63(), 10)

			pmr := models.PrivateChannelRequest{}
			pmr.AccountId = account.Id
			pmr.Body = "test initial participation message"
			pmr.GroupName = groupName
			pmr.Recipients = []string{"chris", "devrim"}

			cc, err := rest.SendPrivateChannelRequest(pmr)
			So(err, ShouldBeNil)
			So(cc, ShouldNotBeNil)

			ses, err := models.FetchOrCreateSession(account.Nick, groupName)
			So(err, ShouldBeNil)
			So(ses, ShouldNotBeNil)

			history, err := rest.GetHistory(
				cc.Channel.Id,
				&request.Query{
					AccountId: account.Id,
				},
				ses.ClientId,
			)

			So(err, ShouldBeNil)
			So(history, ShouldNotBeNil)
			So(len(history.MessageList), ShouldEqual, 2)

			joinMessage := history.MessageList[1].Message
			So(joinMessage.TypeConstant, ShouldEqual, models.ChannelMessage_TYPE_SYSTEM)
			So(joinMessage.Payload, ShouldNotBeNil)
			initialParticipants, ok := joinMessage.Payload["initialParticipants"]
			So(ok, ShouldBeTrue)

			systemType, ok := history.MessageList[1].Message.Payload["systemType"]
			So(ok, ShouldBeTrue)
			So(*systemType, ShouldEqual, models.PrivateMessageSystem_TYPE_JOIN)

			participants := make([]string, 0)
			err = json.Unmarshal([]byte(*initialParticipants), &participants)
			So(err, ShouldBeNil)
			So(len(participants), ShouldEqual, 2)
			So(participants, ShouldContain, "chris")
			// So(*addedBy, ShouldEqual, account.OldId)

		})

		Convey("targetted account should be able to list private message channel of himself", nil)

	})
}

package vk

import (
	"context"
	"fmt"
	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/SevereCloud/vksdk/v2/events"
	"log"
	"strconv"
	"strings"
	"unicode/utf8"
)

const (
	msgStart       = "привет"
	msgChangeGroup = "изменить"
)

func (b *Bot) handleNewMessages() {
	b.bot.MessageNew(func(_ context.Context, obj events.MessageNewObject) {
		nickname, err := b.getUserName(obj.Message.FromID)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("%s: %s", nickname, obj.Message.Text)

		if utf8.RuneCountInString(obj.Message.Text) > 0 {
			err = b.handleNewTextMessage(obj)
			if err != nil {
				log.Fatal(err)
			}
		}
	})
}

func (b *Bot) handleNewTextMessage(obj events.MessageNewObject) error {
	var answer string

	switch strings.ToLower(obj.Message.Text) {
	case msgStart:
		answer = b.messages.StartWithoutGroup
	case msgChangeGroup:
		answer = b.messages.ChangeGroup
	default:
		answer = b.messages.GroupNotSelected
	}

	err := b.sendMessage(obj.Message.FromID, answer)
	return err
}

func (b *Bot) getUserName(userID int) (string, error) {
	p := params.UsersGetBuilder{Params: map[string]interface{}{}}
	p.UserIDs([]string{strconv.Itoa(userID)})

	response, err := b.bot.VK.UsersGet(p.Params)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s %s", response[0].FirstName, response[0].LastName), nil
}

func (b *Bot) sendMessage(peerID int, text string) error {
	p := params.NewMessagesSendBuilder()
	p.Message(text)
	p.RandomID(0)
	p.UserID(peerID)

	_, err := b.bot.VK.MessagesSend(p.Params)
	return err
}

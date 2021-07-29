package vk

import (
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/longpoll-bot"
	"github.com/tmrrwnxtsn/ulstu-schedule-bot/pkg/config"
	"log"
)

type Bot struct {
	bot      *longpoll.LongPoll
	messages config.Messages
}

func NewBot(token string, messages config.Messages) (*Bot, error) {
	vk := api.NewVK(token)

	group, err := vk.GroupsGetByID(nil)
	if err != nil {
		return nil, err
	}

	lp, err := longpoll.NewLongPoll(vk, group[0].ID)
	if err != nil {
		return nil, err
	}

	return &Bot{bot: lp, messages: messages}, nil
}

func (b *Bot) RunPolling() error {
	b.handleNewMessages()

	log.Println("Start polling!")
	err := b.bot.Run()
	return err
}

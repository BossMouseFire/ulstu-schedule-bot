package main

import (
	"github.com/tmrrwnxtsn/ulstu-schedule-vk-bot/internal/config"
	"github.com/tmrrwnxtsn/ulstu-schedule-vk-bot/internal/vk"
	"log"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	bot, err := vk.NewBot(cfg.Token, cfg.Messages)
	if err != nil {
		log.Fatal(err)
	}

	err = bot.RunPolling()
	if err != nil {
		log.Fatal(err)
	}
}

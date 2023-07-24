package main

import (
	"flag"
	"log"
	"telegaBot/internal/app/commands"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	tf := flag.String("t", "", "токен")
	flag.Parse()

	token := *tf

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}
	//u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}
	commander := commands.NewCommander(bot)
	for update := range updates {

		commander.HandleUpdate(update)

	}
}

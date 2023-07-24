package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"telegaBot/internal/service/DataBases"
)

func (c *Commander) horoscope(inputMessage *tgbotapi.Message) {
	Zodiac := DataBases.GetUserZodiac(inputMessage.Chat.ID)
	if Zodiac == "" {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Введите дату в формате ДД-ММ например (30-03)")
		c.bot.Send(msg)
		DataBases.UpdatePosition(inputMessage.Chat.ID, 2)
	} else {
		//text := ChatGPT.GetanswerGpt(fmt.Sprintf("гороскоп для %d на сегодня", Zodiac))
		text := DataBases.Gethoroscope(Zodiac)
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, text)
		c.bot.Send(msg)
		DataBases.UpdatePosition(inputMessage.Chat.ID, 1)
	}
}

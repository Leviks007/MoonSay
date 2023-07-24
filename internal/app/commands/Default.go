package commands

import (
	"telegaBot/internal/service/ChatGPT"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *Commander) Default(inputMessage *tgbotapi.Message) {
	texst := ChatGPT.GetanswerGpt(inputMessage.Text)

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, texst)
	//msg := tgbotapi.NewMessage(572178665, inputMessage.Text+texst)

	c.bot.Send(msg)
}

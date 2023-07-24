package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"telegaBot/internal/service/DataBases"
)

func (c *Commander) ball8question(inputMessage *tgbotapi.Message, selectball int) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Задай свой вопрос Луне 🌚")
	//var numericKeyboard = tgbotapi.NewReplyKeyboard()
	//msg.ReplyMarkup = numericKeyboard
	c.bot.Send(msg)
	if selectball == 1 {

		DataBases.UpdatePosition(inputMessage.Chat.ID, 31)
	} else if selectball == 2 {
		DataBases.UpdatePosition(inputMessage.Chat.ID, 32)
	}
}

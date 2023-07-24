package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"telegaBot/internal/service/DataBases"
)

func (c *Commander) mainMenu(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Выбери из клавиатуры")
	var numericKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Гороскоп 🔮")),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Шар предсказаний 🎱")),
	)

	msg.ReplyMarkup = numericKeyboard

	c.bot.Send(msg)
	DataBases.UpdatePosition(inputMessage.Chat.ID, 1)
}

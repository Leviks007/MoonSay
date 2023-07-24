package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"telegaBot/internal/service/DataBases"
)

func (c *Commander) ball8select(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Могу погадать двумя вариантами\n\n📖 <b>По цитатам:</b> ты задаешь свой вопрос - я даю ответ в виде цитаты великого человека. Что с ней делать ты решаешь сам. \n\n💭 <b>С конкретным ответом:</b> ты также задаешь вопрос, на который можно ответить да/нет. А я соединяюсь с космосом и присылаю тебе ответ. \n\nВыбирай! \n")
	msg.ParseMode = "HTML"
	var numericKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("По цитатам 📖")),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("С конкретным ответом 💭")),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Назад в главное меню ◀️")),
	)

	msg.ReplyMarkup = numericKeyboard

	c.bot.Send(msg)
	DataBases.UpdatePosition(inputMessage.Chat.ID, 3)
}

package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"telegaBot/internal/service/DataBases"
)

func (c *Commander) starter(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "👋🏻 Привет, астрологический путник! \n\nЯ бот <b>«Говорит Луна»</b>. Помогу тебе узнать, что необычного тебя ждет в этот день, а также предскажу будущее насколько это возможно. Вот, что я умею:\n\n🔮 <b>Ежедневный астрологический прогноз</b> (Гороскоп)\nНа основе твоей даты рождения я просчитаю твой знак зодиака и напишу тебе сегодняшнее пожелание на день. \n\n🎱 <b>Задай вопрос и получи ответ</b> (Шар предсказаний)\nВо мне есть рандомайзер, который поможет тебе получить случайный ответ на твой вопрос. А ты решишь, что с ним дальше делать.")
	msg.ParseMode = "HTML"
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

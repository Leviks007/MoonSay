package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"math/rand"
	"telegaBot/internal/service/DataBases"
	"time"
)

func (c *Commander) ball8(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, ball8Predictions())

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

func ball8Predictions() string {
	// Инициализируем генератор случайных чисел
	rand.Seed(time.Now().UnixNano())

	// Определяем список предсказаний
	predictions := []string{
		"Сегодня будет хороший день!",
		"Вам скоро повезет в делах!",
		"Вы встретите новых друзей!",
		"Скоро вы найдете то, что ищете!",
		"Вам нужно проявить больше терпения.",
		"Ваш труд будет вознагражден!",
	}

	// Генерируем случайное число, чтобы выбрать предсказание
	randomIndex := rand.Intn(len(predictions))

	// Выводим предсказание
	return predictions[randomIndex]
}

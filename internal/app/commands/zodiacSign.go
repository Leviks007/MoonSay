package commands

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"telegaBot/internal/service/DataBases"
)

func (c *Commander) zodiacSign(inputMessage *tgbotapi.Message) {
	// Разбить дату рождения на отдельные числа
	var day, month int
	fmt.Sscanf(inputMessage.Text, "%d-%d", &day, &month)
	// Вызовите функцию, чтобы определить знак зодиака
	zodiacSign := getZodiacSign(day, month)
	// Выведите знак зодиака
	fmt.Println("Знак зодиака:", zodiacSign)
	DataBases.UpdateZodiacSign(inputMessage.Chat.ID, zodiacSign)
	//text := ChatGPT.GetanswerGpt(fmt.Sprintf("гороскоп для %d на сегодня", Zodiac))
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "по гороскопу ты "+zodiacSign+" вот твой гороскоп")
	c.bot.Send(msg)
	c.horoscope(inputMessage)
}

// Функция getZodiacSign принимает два параметра - день и месяц - и возвращает строку с названием знака зодиака
func getZodiacSign(day, month int) string {
	// Массив с датами начала знаков зодиака для каждого месяца
	zodiacStartDates := [12]int{20, 19, 21, 20, 21, 21, 23, 23, 23, 23, 22, 22}
	// Массив со строками, содержащими названия знаков зодиака
	zodiacSigns := [12]string{"Водолей", "Рыбы", "Овен", "Телец", "Близнецы", "Рак", "Лев", "Дева", "Весы", "Скорпион", "Стрелец", "Козерог"}
	// Определить индекс знака зодиака в массиве zodiacSigns
	zodiacIndex := month - 1
	if day < zodiacStartDates[zodiacIndex] {
		zodiacIndex = (zodiacIndex - 1) % 12
	}
	// Вернуть название знака зодиака по индексу
	return zodiacSigns[zodiacIndex]
}

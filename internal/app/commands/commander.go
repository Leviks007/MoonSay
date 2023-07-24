package commands

import (
	"encoding/json"
	"fmt"
	"telegaBot/internal/service/DataBases"
	"time"
)
import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

type CommandData struct {
	Offset int `json:"offset"`
}
type Commander struct {
	bot *tgbotapi.BotAPI
}

func NewCommander(bot *tgbotapi.BotAPI,

) *Commander {
	return &Commander{
		bot: bot,
	}

}

func (c *Commander) HandleUpdate(update tgbotapi.Update) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			fmt.Printf("recover from panic: %v", panicValue)
		}
	}()

	if update.CallbackQuery != nil {
		parsedData := CommandData{}
		json.Unmarshal([]byte(update.CallbackQuery.Data), &parsedData)
		msg := tgbotapi.NewMessage(
			update.CallbackQuery.Message.Chat.ID,
			fmt.Sprintf("Command: %+v\n", parsedData),
		)
		c.bot.Send(msg)
		return
	}
	if update.Message == nil {
		return
	}

	userData := DataBases.GetUserData(update.Message.Chat.ID)
	if userData.ID == 0 {
		DataBases.AddUserId(update.Message.Chat.ID)
		userData.ID = update.Message.Chat.ID
		userData.Position = 0
		c.starter(update.Message)
		DataBases.UpdatePosition(update.Message.Chat.ID, 1)
	} else {

		if update.Message.Command() == "start" || update.Message.Command() == "help" {
			c.starter(update.Message)
			DataBases.UpdatePosition(update.Message.Chat.ID, 1)
		} else if update.Message.Text == "Назад в главное меню ◀️" {
			c.mainMenu(update.Message)
		} else {

			switch userData.Position {
			case 1:
				switch update.Message.Text {
				case "Гороскоп 🔮":
					c.horoscope(update.Message)
				case "Шар предсказаний 🎱":
					c.ball8select(update.Message)
				default:
					c.mainMenu(update.Message)
				}

			case 2:
				if isDateValid(update.Message.Text) {
					c.zodiacSign(update.Message)
				} else {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Введите дату в формате ДД-MM например (03-30)")
					c.bot.Send(msg)
				}
			case 3:
				switch update.Message.Text {
				case "По цитатам 📖":
					c.ball8question(update.Message, 2)
				case "С конкретным ответом 💭":
					c.ball8question(update.Message, 1)
				//case "Назад в главное меню ◀️":
				//	c.mainMenu(update.Message)
				default:
					c.ball8select(update.Message)
				}
			case 31:
				c.ball8(update.Message)
			case 32:
				c.quotes(update.Message)
			}
		}
	}
}

//DataBases.UpdatePosition(update.Message.Chat.ID, 0)

func isDateValid(dateString string) bool {
	_, err := time.Parse("02-01", dateString)
	return err == nil
}

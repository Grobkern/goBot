package main

import (
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("669872325:AAFU0Fn6QHXnoU12LYi7CxxXem2GF8eemDA")
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Pizdec.Fuckin'g golang nahui %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			switch update.Message.Command() {
			case "help":
				msg.Text = "try /ban and /f"
			case "ban":
				t := time.Date(2001, time.September, 9, 1, 46, 40, 0, time.UTC)
				k := tgbotapi.ChatMemberConfig{
					ChatID: update.Message.Chat.ID,
					UserID: update.Message.ReplyToMessage.From.ID,
				} //update.Message.ReplyToMessage.From.ID
				bot.KickChatMember(tgbotapi.KickChatMemberConfig{k, t.Unix()})
				//tgbotapi.KickChatMemberConfig{tgbotapi.ChatMemberConfig{chatID, "@grobkernux", " ", bot.GetChatMember(tgbotapi.ChatConfigWithUser{chatID, "@grobkernux", update.Message.From.ID})}, t.Unix()}
				//if update.Message.From.Id.IsAdmin(update.Message.From.ID)
				//bot.GetChatAdministrators(tgbotapi.ChatConfig{update.Message.Chat.ID, ""})
			case "stable":
				uid := update.Message.From.ID
				roll := tgbotapi.ChatMember{tgbotapi.User{update.Message.From.ID, update.Message.From.FirstName}, update.Message.From.ID}.IsAdministrator()
				if roll == true {
					messid := update.Message.MessageID
					msg.Text = "I'm ready"
				}

			default:
				msg.Text = "I don't know that command"
			}
			bot.Send(msg)
		}
	}
}

package main

import (
	"log"
	"net/http"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type kek struct {
	Ok     bool `json:"ok"`
	Result struct {
		User struct {
			ID           int    `json:"id"`
			IsBot        bool   `json:"is_bot"`
			FirstName    string `json:"first_name"`
			Username     string `json:"username"`
			LanguageCode string `json:"language_code"`
		} `json:"user"`
		Status string `json:"status"`
	} `json:"result"`
}

func main() {
	var (
		replyID    int
		strreplyID string
	)
	var (
		chatID    int64
		strchatID string
	)
	var (
		userID    int
		struserID string
	)
	var (
		stableID    int
		strstableID string
	)
	bot, err := tgbotapi.NewBotAPI("669872325:AAFU0Fn6QHXnoU12LYi7CxxXem2GF8eemDA")
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Connection complete %s", bot.Self.UserName)
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
				checkAdmin(strchatID, struserID, &chatID, &userID)
			case "savestab":
				stableID = update.Message.MessageID
				chatID = update.Message.Chat.ID

			case "stable":
				strstableID = strconv.Itoa(stableID)
				chatID = update.Message.Chat.ID
				strchatID = strconv.FormatInt(chatID, 10)
				chatID2 := update.Message.Chat.ID
				strchatID2 := strconv.FormatInt(chatID, 10)
				http.Get("https://api.telegram.org/bot669872325:AAFU0Fn6QHXnoU12LYi7CxxXem2GF8eemDA/forwardMessage?chat_id=" + strchatID2 + "&from_chat_id=" + strchatID + "&message_id=" + strstableID)
			default:

			}
			bot.Send(msg)
		}
	}
}

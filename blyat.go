package main

import (
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

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
		gayRand int
		counter int
	)
	var (
		replyID int
		//strreplyID string
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
	var token string = "669872325:AAFU0Fn6QHXnoU12LYi7CxxXem2GF8eemDA"
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
				replyID = update.Message.ReplyToMessage.From.ID
				chatID = update.Message.Chat.ID
				userID = update.Message.From.ID
				checkAdmin(&strchatID, &struserID, &chatID, &userID, &replyID)
			case "ping":
				message := tgbotapi.NewMessage(update.Message.Chat.ID, "Fuck you")
				bot.Send(message)
			case "savestab":
				stableID = update.Message.MessageID
				chatID = update.Message.Chat.ID
			case "f":
				chatID = update.Message.Chat.ID
				strchatID = strconv.FormatInt(chatID, 10)
				http.Get("https://api.telegram.org/bot" + token + "/sendSticker?chat_id=" + strchatID + "&sticker=CAADAgADsgADTptkAm1WnTBWvUfiAg")
			case "gay":
				gayRand = rand.Intn(100)
				if counter == 5 {
					msg.Text = "Wait"
					bot.Send(msg)
					time.Sleep(10 * time.Second)
					counter = 0
				}
				msg.Text = "You are gay with chance:" + strconv.Itoa(gayRand) + "%"
				counter++
			case "Foxed":
				msg.Text = "http://qiwi.me/f0x1d"
			case "stable":
				strstableID = strconv.Itoa(stableID)
				chatID = update.Message.Chat.ID
				strchatID = strconv.FormatInt(chatID, 10)
				chatID2 := update.Message.Chat.ID
				strchatID2 := strconv.FormatInt(chatID2, 10)
				http.Get("https://api.telegram.org/bot669872325:AAFU0Fn6QHXnoU12LYi7CxxXem2GF8eemDA/forwardMessage?chat_id=" + strchatID2 + "&from_chat_id=" + strchatID + "&message_id=" + strstableID)
			default:

			}
			bot.Send(msg)
		}
	}
}

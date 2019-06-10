package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Result struct {
	status string "json:'status'"
}
type kek struct {
	result Result
}

func Kick(chatid int64, userid int) {
	//bot, err := tgbotapi.NewBotAPI("669872325:AAFU0Fn6QHXnoU12LYi7CxxXem2GF8eemDA")
	token := "669872325:AAFU0Fn6QHXnoU12LYi7CxxXem2GF8eemDA"
	/*if err != nil {
		log.Panic(err)
	}
	t := time.Date(2001, time.September, 9, 1, 46, 40, 0, time.UTC)
	k := tgbotapi.ChatMemberConfig{
		ChatID: chatid,
		UserID: userid,
	}
	bot.KickChatMember(tgbotapi.KickChatMemberConfig{k, t.Unix()})
	*/
	usid := string(userid)
	http.Get("https://api.telegram.org/bot" + token + "/kickChatMember?chat_id=@grobkernux&user_id=" + usid)
}

func main() {
	bot, err := tgbotapi.NewBotAPI("669872325:AAFU0Fn6QHXnoU12LYi7CxxXem2GF8eemDA")
	token := "669872325:AAFU0Fn6QHXnoU12LYi7CxxXem2GF8eemDA"
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
				t := time.Date(2001, time.September, 9, 1, 46, 40, 0, time.UTC)
				k := tgbotapi.ChatMemberConfig{
					ChatID: update.Message.Chat.ID,
					UserID: update.Message.ReplyToMessage.From.ID,
				}
				chid := update.Message.Chat.ID
				uid := update.Message.From.ID
				uid_string := string(uid)
				bot.KickChatMember(tgbotapi.KickChatMemberConfig{k, t.Unix()})
				api_get, err := http.Get("https://api.telegram.org/bot" + token + "/getChatMember?chat_id=@grobkernux&user_id" + uid_string)
				if err != nil {
					log.Println(err)
				}
				var app = kek{}
				//var api map[string]interface{}
				json.NewDecoder(api_get.Body).Decode(&app)
				switch string(app.result.status) {
				case "admin":
					Kick(chid, uid)
				case "creator":
					Kick(chid, uid)
				default:
					msg.Text = "You are not admin"
				}

			case "stable":
			//	uid := update.Message.From.ID
			//	roll := tgbotapi.ChatMember{tgbotapi.User{update.Message.From.ID, update.Message.From.FirstName}, update.Message.From.ID}.IsAdministrator()
			case "dev":
			default:
				msg.Text = "I don't know that command"
			}
			bot.Send(msg)
		}
	}
}

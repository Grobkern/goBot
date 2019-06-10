package main

import (
	"encoding/json"
	"log"
	"net/http"

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

func Kick(chatid int64, userid int) {
	var token string = "669872325:AAFU0Fn6QHXnoU12LYi7CxxXem2GF8eemDA"
	usid := string(userid)
	http.Get("https://api.telegram.org/bot" + token + "/kickChatMember?chat_id=@grobkernux&user_id=" + usid)
}
func adminCheck(_chid *int64, _uid *int, _token *string) {
	apiGet, err := http.Get("https://api.telegram.org/bot" + *_token + "/getChatMember?chat_id=@grobkernux&user_id" + string(*_uid))
	if err != nil {
		log.Println(err)
	}
	var app = kek{}
	//var api map[string]interface{}
	json.NewDecoder(apiGet.Body).Decode(&app)
	switch string(app.Result.Status) {
	case "admin":
		Kick(*_chid, *_uid)
	case "creator":
		Kick(*_chid, *_uid)
	default:
	}
}
func main() {
	bot, err := tgbotapi.NewBotAPI("669872325:AAFU0Fn6QHXnoU12LYi7CxxXem2GF8eemDA")
	//token := "669872325:AAFU0Fn6QHXnoU12LYi7CxxXem2GF8eemDA"
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
				chid := update.Message.Chat.ID
				uid := update.Message.From.ID
				//uidString := string(uid)
				apiGet, err := http.Get("https://api.telegram.org/bot669872325:AAFU0Fn6QHXnoU12LYi7CxxXem2GF8eemDA/getChatMember?chat_id=@grobkernux&user_id=404334300")
				//apiGet, err := http.Get("https://api.telegram.org/bot" + token + "/getChatMember?chat_id=@grobkernux&user_id=" + uidString)
				if err != nil {
					log.Println(err)
				}
				log.Print(apiGet.Body)
				app := kek{}
				//var api map[string]interface{}
				json.NewDecoder(apiGet.Body).Decode(&app)
				log.Print(app)
				switch string(app.Result.Status) {
				case "administrator":
					Kick(chid, uid)
				case "creator":
					Kick(chid, uid)
				default:
					//msg.Text = "You are not admin"
				}

			case "stable":

			case "dev":
			default:

			}
			bot.Send(msg)
		}
	}
}

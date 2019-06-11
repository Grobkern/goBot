package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type unsplash struct {
	ID             string      `json:"id"`
	CreatedAt      string      `json:"created_at"`
	UpdatedAt      string      `json:"updated_at"`
	Width          int         `json:"width"`
	Height         int         `json:"height"`
	Color          string      `json:"color"`
	Description    interface{} `json:"description"`
	AltDescription interface{} `json:"alt_description"`
	Urls           struct {
		Raw     string `json:"raw"`
		Full    string `json:"full"`
		Regular string `json:"regular"`
		Small   string `json:"small"`
		Thumb   string `json:"thumb"`
	} `json:"urls"`
	Links struct {
		Self             string `json:"self"`
		HTML             string `json:"html"`
		Download         string `json:"download"`
		DownloadLocation string `json:"download_location"`
	} `json:"links"`
	Categories             []interface{} `json:"categories"`
	Sponsored              bool          `json:"sponsored"`
	SponsoredBy            interface{}   `json:"sponsored_by"`
	SponsoredImpressionsID interface{}   `json:"sponsored_impressions_id"`
	Likes                  int           `json:"likes"`
	LikedByUser            bool          `json:"liked_by_user"`
	CurrentUserCollections []interface{} `json:"current_user_collections"`
	User                   struct {
		ID              string      `json:"id"`
		UpdatedAt       string      `json:"updated_at"`
		Username        string      `json:"username"`
		Name            string      `json:"name"`
		FirstName       string      `json:"first_name"`
		LastName        string      `json:"last_name"`
		TwitterUsername interface{} `json:"twitter_username"`
		PortfolioURL    interface{} `json:"portfolio_url"`
		Bio             string      `json:"bio"`
		Location        string      `json:"location"`
		Links           struct {
			Self      string `json:"self"`
			HTML      string `json:"html"`
			Photos    string `json:"photos"`
			Likes     string `json:"likes"`
			Portfolio string `json:"portfolio"`
			Following string `json:"following"`
			Followers string `json:"followers"`
		} `json:"links"`
		ProfileImage struct {
			Small  string `json:"small"`
			Medium string `json:"medium"`
			Large  string `json:"large"`
		} `json:"profile_image"`
		InstagramUsername string `json:"instagram_username"`
		TotalCollections  int    `json:"total_collections"`
		TotalLikes        int    `json:"total_likes"`
		TotalPhotos       int    `json:"total_photos"`
		AcceptedTos       bool   `json:"accepted_tos"`
	} `json:"user"`
	Exif struct {
		Make         string `json:"make"`
		Model        string `json:"model"`
		ExposureTime string `json:"exposure_time"`
		Aperture     string `json:"aperture"`
		FocalLength  string `json:"focal_length"`
		Iso          int    `json:"iso"`
	} `json:"exif"`
	Views     int `json:"views"`
	Downloads int `json:"downloads"`
}
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
	var text string
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
			case "8":
				gayRand = rand.Intn(4)
				switch gayRand {
				case 0:
					msg.Text = "Мой ответ - 'да'"
				case 1:
					msg.Text = "Скорее всего да"
				case 2:
					msg.Text = "Скорее всего нет"
				case 3:
					msg.Text = "Мой ответ-'нет'"
				}
			case "unsplash":
				httpGet, err := http.Get("https://api.unsplash.com/photos/random?client_id=1435c8eaadfbeacd502ec854e73123059456f3a601722e790c009bd40fdfe15b")
				if err != nil {
					log.Print(err)
				}
				var photos = unsplash{}
				json.NewDecoder(httpGet.Body).Decode(&photos)
				text = photos.Links.Download
				msg.Text = text
			case "shrug":
				msg.Text = "¯\\_(ツ)_/¯'"
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

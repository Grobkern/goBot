package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func checkAdmin(_strchatID *string, _struserID *string, _chatID *int64, _userID *int, _replyID *int) {
	var token = "669872325:AAFU0Fn6QHXnoU12LYi7CxxXem2GF8eemDA"
	*_struserID = strconv.Itoa(*_userID)
	log.Print(*_struserID)
	*_strchatID = strconv.FormatInt(*_chatID, 10)
	log.Print(*_strchatID)
	apiGet, err := http.Get("https://api.telegram.org/bot" + token + "/getChatMember?chat_id=" + *_strchatID + "&user_id=" + *_struserID)
	errcheck(&err)
	var app = kek{}
	json.NewDecoder(apiGet.Body).Decode(&app)
	switch app.Result.Status {
	case "administrator":
		kick(*_replyID, *_chatID)
	case "creator":
		kick(*_replyID, *_chatID)
	default:
	}
}

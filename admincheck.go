package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func checkAdmin(_strchatID string, _struserID string, _chatID *int64, _userID *int) {
	var token string = "669872325:AAFU0Fn6QHXnoU12LYi7CxxXem2GF8eemDA"
	var adminRequest string = "https://api.telegram.org/bot669872325:AAFU0Fn6QHXnoU12LYi7CxxXem2GF8eemDA/getChatMember?chat_id=" + _strchatID + "&user_id=" + _struserID
	apiGet, err := http.Get(adminRequest)
	if err != nil {
		log.Print(err)
	}
	var app = kek{}
	json.NewDecoder(apiGet.Body).Decode(&app)
	switch app.Result.Status {
	case "administrator":
		kick(*_chatID, *_userID)
	case "owner":
		kick(*_chatID, *_userID)
	default:

	}
}

package main

import (
	"log"
	"net/http"
	"strconv"
)

func kick(_userID int, _chatID int64) {
	log.Print(_userID)
	usid := strconv.Itoa(_userID)
	basechatid := strconv.FormatInt(_chatID, 10)
	var kickRequest string = "https://api.telegram.org/bot669872325:AAFU0Fn6QHXnoU12LYi7CxxXem2GF8eemDA/kickChatMember?chat_id=" + basechatid + "&user_id=" + usid
	log.Print(usid)
	http.Get(kickRequest)
}

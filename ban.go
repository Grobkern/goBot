package main

import (
	"log"
	"net/http"
	"strconv"
)

func kick(userid int, chatid int64) {
	log.Print(userid)
	usid := strconv.Itoa(userid)
	basechatid := strconv.FormatInt(chatid, 10)
	var kickRequest string = "https://api.telegram.org/bot669872325:AAFU0Fn6QHXnoU12LYi7CxxXem2GF8eemDA/kickChatMember?chat_id=" + basechatid + "&user_id=" + usid
	log.Print(usid)
	http.Get(kickRequest)
}

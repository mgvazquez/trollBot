package main

import (
	"github.com/nlopes/slack"
	"strings"
	"math/rand"
	"time"
)

func randomResponse(rtas ...string) (string)  {
	rand.Seed(time.Now().Unix())
	return ""+rtas[rand.Intn(len(rtas))]+""
}

//func stringMatch(msg string, subs ...string) (Match bool, String string){
func stringMatch(msg string, subs ...string) (Match bool){
	//Match := false
	//String := ""
	for _, sub := range subs {
		if strings.Contains(strings.ToLower(msg), strings.ToLower(sub)) {
			Match = true
			//String = sub
			break
		}
	}
	//return Match, String
	return Match
}

func listenAndRespond(rtm *slack.RTM, msg *slack.MessageEvent) {
	botID := &rtm.GetInfo().User.ID

	//TODO: La idea es agregar aca otros handlers para aceptar comandos directos nombrando al bot

	if !stringMatch(msg.User, *botID, "USLACKBOT") {
		go edgewise(rtm,msg)
	}

}
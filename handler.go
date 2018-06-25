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

func stringMatch(msg string, subs ...string) (bool){
	_match := false
	for _, sub := range subs {
		if strings.Contains(strings.ToLower(msg), strings.ToLower(sub)) {
			_match = true
			break
		}
	}
	return _match
}

func listenAndRespond(rtm *slack.RTM, msg *slack.MessageEvent) {
	botID := &rtm.GetInfo().User.ID
	//usrID := &msg.User

	//TODO: La idea es agregar aca otros handlers para aceptar comandos directos nombrando al bot

	if !stringMatch(msg.User, *botID, "USLACKBOT") {
		go edgewise(rtm,msg)
	}

}
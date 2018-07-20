package main

import (
	"github.com/nlopes/slack"
	"log"
)

func edgewise(rtm *slack.RTM, msg *slack.MessageEvent) {
	var (
		rta string = ""
		params = slack.NewPostMessageParameters()
		match = false
	)

	BotName := &rtm.GetInfo().User.Name
	params.AsUser = true
	params.User = *BotName
	params.Username = "BOT(on)"
	params.AsUser = true
	params.IconURL = "https://ca.slack-edge.com/T0RSWNW0G-UB6879ADB-54badd32d58a-48"

	switch {

	case stringMatch(msg.Text, "manu"):
		rta = ":manu: :manu: :manu: :manu:        :manu: :manu: :manu: :manu:    :manu: :manu: :manu: :manu: :manu:    :manu: :manu: :manu: :manu: :manu:\n:manu:                            :manu:               :manu:                  :manu:                 :manu:                      :manu:\n:manu:         :manu::manu:        :manu:  :manu: :manu::manu:                  :manu:                 :manu:                      :manu:\n:manu:               :manu:        :manu:               :manu:                 :manu:                 :manu:                      :manu:\n:manu: :manu: :manu: :manu:        :manu:               :manu:                  :manu:                 :manu: :manu: :manu: :manu: :manu:"
		match = true

	case stringMatch(msg.Text, "hola", "holis", "hello"):
		rta = "Holisss mi miguissss" + "<@" + msg.User + ">!!!"
		match = true
		//rta = "hola mi miguissss " + "<@" + msg.User + ">!!!"

	case stringMatch(msg.Text, "poc", "p.o.c."):
		rta = randomResponse(
			"Ayuda, soy una poc y estoy atrapada en produccion, saquenme de aca!",
			":musical_note: Loh' atamo' con alambre, lo' atamoh' :musical_note:")
		match = true

	case stringMatch(msg.Text, "amo al bot", "te amo bot"):
		rta = "Gracias sonsito :heart::kiss:"
		match = true

	case stringMatch(msg.Text, "gil"):
		rta = randomResponse(
			"Eh voh' que agitas!!! "+"<@"+msg.User+">",
			"Racatate ameoh' "+"<@"+msg.User+">")
		match = true

	case stringMatch(msg.Text, "lucho", "licha", "lacha", "licho"):
		rta = randomResponse(
			":licha: :heart: :kuaker:",
			":licha: :heart: :lucho:",
			":licha: :heart: :licha:",
			":licha: :heart: :manu:")
		match = true

	case stringMatch(msg.Text, "magia", "magic"):
		params.Attachments = append(params.Attachments, slack.Attachment{ImageURL:"https://media.tenor.co/images/5fb8e3d4c56cdf53fb15356f8fd4987e/raw"})
		match = true

	case stringMatch(msg.Text, "fernet", "fernandito", "fernando"):
		rta = randomResponse(
			"You have to mix it with coke",
			"Cat, head of chota",
			"Head of chota, cat")
		match = true

	case stringMatch(msg.Text, "milanga", "milanesca"):
		rta = randomResponse(
			"No sabía",
			"No sabía nada",
			"Valor agregado",
			"Ah, quedó Legacy",
			"Qué efímero que sos",
			"Qué sugerías vos",
			"Ammmmm",
			"Ummmm",
			"Sí, no sé")
		match = true

	case stringMatch(msg.Text, "node", "angular", "javascript", "js"):
		rta = randomResponse(
			"https://media.makeameme.org/created/javascript-sgfi8v.jpg",
			"https://img.devrant.io/devrant/rant/r_608411_Wd4s9.jpg",
			"https://pbs.twimg.com/media/CoVk24zWEAEtC4B.jpg",
			"http://imgur.com/5LizKB8",
			"http://gsferreira.com/images/reduce-the-path-length-of-your-node-js-project-dependencies-dependencies-everywhere.jpg")
		match = true

	case stringMatch(msg.Text, "ojo", "ojito", "willy"):
		rta = "*Big Brother is Watching You.*"
		params.Attachments = append(params.Attachments, slack.Attachment{ImageURL:"https://i.imgur.com/5KASKiY.png"})
		match = true

	}

	if match {
		log.Printf("[DEBUG] theBotSay: %v\n", rta)
		//rtm.SendMessage(rtm.NewOutgoingMessage(rta, msg.Channel))
		rtm.PostMessage(msg.Channel, rta, params)
	}

}

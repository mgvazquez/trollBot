package main

import (
	"github.com/nlopes/slack"
	"log"
)

func edgewise(rtm *slack.RTM, msg *slack.MessageEvent) {
	var (
		rta string = ""
	)

	switch {

	case stringMatch(msg.Text, "manu"):
		rta = ":manu: :manu: :manu: :manu:        :manu: :manu: :manu: :manu:    :manu: :manu: :manu: :manu: :manu:    :manu: :manu: :manu: :manu: :manu:\n:manu:                            :manu:               :manu:                  :manu:                 :manu:                      :manu:\n:manu:         :manu::manu:        :manu:  :manu: :manu::manu:                  :manu:                 :manu:                      :manu:\n:manu:               :manu:        :manu:               :manu:                 :manu:                 :manu:                      :manu:\n:manu: :manu: :manu: :manu:        :manu:               :manu:                  :manu:                 :manu: :manu: :manu: :manu: :manu:"

	case stringMatch(msg.Text, "hola", "holis", "hello"):
		rta = "Holisss mi miguissss" + "<@" + msg.User + ">!!!"
		//rta = "hola mi miguissss " + "<@" + msg.User + ">!!!"

	case stringMatch(msg.Text, "poc", "p.o.c."):
		rta = randomResponse(
			"Ayuda, soy una poc y estoy atrapada en produccion, saquenme de aca!",
			":musical_note: Loh' atamo' con alambre, lo' atamoh' :musical_note:")

	case stringMatch(msg.Text, "amo al bot", "te amo bot"):
		rta = "Gracias sonsito :heart::kiss:"

	case stringMatch(msg.Text, "gil"):
		rta = randomResponse(
			"Eh voh' que agitas!!! "+"<@"+msg.User+">",
			"Racatate ameoh' "+"<@"+msg.User+">")

	case stringMatch(msg.Text, "lucho", "licha", "lacha", "licho"):
		rta = randomResponse(
			":licha: :heart: :kuaker:",
			":licha: :heart: :lucho:",
			":licha: :heart: :licha:",
			":licha: :heart: :manu:")

	case stringMatch(msg.Text, "magia", "magic"):
		rta = "https://media.tenor.co/images/5fb8e3d4c56cdf53fb15356f8fd4987e/raw"

	case stringMatch(msg.Text, "fernet", "fernandito", "fernando"):
		rta = randomResponse(
			"You have to mix it with coke",
			"Cat, head of chota",
			"Head of chota, cat")

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

	case stringMatch(msg.Text, "node", "angular", "javascript", "js"):
		rta = randomResponse(
			"https://media.makeameme.org/created/javascript-sgfi8v.jpg",
			"https://img.devrant.io/devrant/rant/r_608411_Wd4s9.jpg",
			"https://pbs.twimg.com/media/CoVk24zWEAEtC4B.jpg",
			"http://imgur.com/5LizKB8",
			"http://gsferreira.com/images/reduce-the-path-length-of-your-node-js-project-dependencies-dependencies-everywhere.jpg")

	case stringMatch(msg.Text, "ojo", "ojito", "willy"):
		rta = "Big Brother is Watching You\nhttps://i.imgur.com/5KASKiY.png"

}

	if rta != "" {
		log.Printf("[DEBUG] theBotSay: %v\n", rta)
		rtm.SendMessage(rtm.NewOutgoingMessage(rta, msg.Channel))

		//params := slack.NewPostMessageParameters()
		//params.Username = "BOT(on)"
		//params.IconURL = "https://ca.slack-edge.com/T0RSWNW0G-UB6879ADB-54badd32d58a-48"
		//params.Attachments = append(params.Attachments, slack.Attachment{ImageURL:"https://i.imgur.com/5LizKB8.png"})
		//rtm.PostMessage(msg.Channel, "", params)

	}

}

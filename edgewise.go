package main

import (
	"github.com/nlopes/slack"
	"log"
)

func edgewise(rtm *slack.RTM, msg *slack.MessageEvent) {

	var (
		rta    string = ""
		img    string = ""
		params        = slack.NewPostMessageParameters()
		match         = false
	)

	BotName := &rtm.GetInfo().User.Name
	params.EscapeText = false
	params.User = *BotName
	params.Username = "BOT(on)"
	params.AsUser = true
	params.Markdown = true
	params.IconURL = "https://ca.slack-edge.com/T0RSWNW0G-UB6879ADB-54badd32d58a-48"

	switch {

	case stringMatch(msg.Text, "manu"):
		rta = ":manu: :manu: :manu: :manu:        :manu: :manu: :manu: :manu:    :manu: :manu: :manu: :manu: :manu:    :manu: :manu: :manu: :manu: :manu:\n:manu:                            :manu:               :manu:                  :manu:                 :manu:                      :manu:\n:manu:         :manu::manu:        :manu:  :manu: :manu::manu:                  :manu:                 :manu:                      :manu:\n:manu:               :manu:        :manu:               :manu:                 :manu:                 :manu:                      :manu:\n:manu: :manu: :manu: :manu:        :manu:               :manu:                  :manu:                 :manu: :manu: :manu: :manu: :manu:"
		match = true

	case stringMatch(msg.Text, "hola", "holis", "hello"):
		rta = "Holisss mi miguissss " + "<@" + msg.User + ">!!!"
		match = true

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
		rta = "<@" + msg.User + ">, " + "deja de tirar tanta magia!"
		params.Attachments = append(params.Attachments, slack.Attachment{ImageURL: "https://media.tenor.co/images/5fb8e3d4c56cdf53fb15356f8fd4987e/raw"})
		match = true

	case stringMatch(msg.Text, "fernet", "fernandito", "fernando"):
		rta = randomResponse(
			"You have to mix it with coke",
			"Cat, head of chota",
			"Head of chota, cat")
		match = true

	case stringMatch(msg.Text, "milanga", "milanesa"):
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

	case stringMatch(msg.Text, "proxy"):
		rta = randomResponse(
			"Nah Nah Nah"+" <@"+msg.User+">, "+"otro proxy más?!?!",
			":kuaker:, alguien dijo PROXY?!?! :trollface:")
		match = true

	case stringMatch(msg.Text, "java"):
		img = randomResponse(
			"https://scontent-eze1-1.xx.fbcdn.net/v/t1.0-9/17103571_258039787985510_4543998202758018301_n.jpg?_nc_cat=0&oh=8f6bcf31d076e8025c7651cae16c5ea0&oe=5BEED0DA",
			"https://scontent-eze1-1.xx.fbcdn.net/v/t1.0-9/16711881_249105438878945_4294265477870217883_n.jpg?_nc_cat=0&oh=84eff3784bebffeff4c1e4dc03d93127&oe=5C34C56B",
			"https://scontent-eze1-1.xx.fbcdn.net/v/t1.0-9/17191448_259605207828968_1304022088973812803_n.jpg?_nc_cat=0&oh=f722acd747ecd7ff2eb6f1cb67c00161&oe=5C39F704")
		params.Attachments = append(params.Attachments, slack.Attachment{ImageURL: img})
		match = true

	case stringMatch(msg.Text, "css", "style-sheet"):
		img = randomResponse(
			"https://i.imgur.com/Q3cUg29.gif",
			"https://i.imgur.com/kv7Eods.png")
		params.Attachments = append(params.Attachments, slack.Attachment{ImageURL: img})
		match = true

	case stringMatch(msg.Text, ".net"):
		img = randomResponse(
			"https://scontent-eze1-1.xx.fbcdn.net/v/t1.0-9/22310395_361813147608173_1565090941484148499_n.jpg?_nc_cat=0&oh=c395effab9a8eae2c295f778ce29cbb6&oe=5BF043F4")
		params.Attachments = append(params.Attachments, slack.Attachment{ImageURL: img})
		match = true

	case stringMatch(msg.Text, "node", "angular", "javascript", "js"):
		img = randomResponse(
			"https://media.makeameme.org/created/javascript-sgfi8v.jpg",
			"https://img.devrant.io/devrant/rant/r_608411_Wd4s9.jpg",
			"https://pbs.twimg.com/media/CoVk24zWEAEtC4B.jpg",
			"https://scontent-eze1-1.xx.fbcdn.net/v/t1.0-9/21752493_353251048464383_8517848822595315187_n.jpg?_nc_cat=0&oh=082b5650978abf14e57f58f1ef074296&oe=5C2799E2",
			"https://i.imgur.com/5LizKB8.png",
			"https://scontent-eze1-1.xx.fbcdn.net/v/t1.0-9/15541620_218053788650777_1713605560900849432_n.jpg?_nc_cat=0&oh=466e0a28fc468b33aa1688b6a871f207&oe=5C21EA83",
			"http://gsferreira.com/images/reduce-the-path-length-of-your-node-js-project-dependencies-dependencies-everywhere.jpg")
		params.Attachments = append(params.Attachments, slack.Attachment{ImageURL: img})
		match = true

	case stringMatch(msg.Text, "chrome", "browser"):
		img = randomResponse(
			"https://scontent-eze1-1.xx.fbcdn.net/v/t1.0-9/26195536_397789600677194_997368045716396157_n.png?_nc_cat=0&oh=7f7a08c2b6a09bd3414e7f49398e8066&oe=5C3834A1",
			"https://scontent-eze1-1.xx.fbcdn.net/v/t31.0-8/23799936_378361682619986_7479532257492459262_o.jpg?_nc_cat=0&oh=d6f028930e946e2cde667d148cc5b3e7&oe=5C2FA87B",
			"https://scontent-eze1-1.xx.fbcdn.net/v/t1.0-9/22528485_365306187258869_2768680746694188101_n.png?_nc_cat=0&oh=d979f339594c2c0793a454fa5e27dff3&oe=5C321A3B",
			"https://scontent-eze1-1.xx.fbcdn.net/v/t1.0-9/23130696_370944810028340_6625228377914658080_n.png?_nc_cat=0&oh=3f7f0d6e6d9f2bb2d1325f91c4f4f6e4&oe=5C1FD66E")
		params.Attachments = append(params.Attachments, slack.Attachment{ImageURL: img})
		match = true

	case stringMatch(msg.Text, "php"):
		rta = "Ay!, que divertido que es bardear a PHP."
		img = randomResponse(
			"https://scontent-eze1-1.xx.fbcdn.net/v/t1.0-9/18740765_304542476668574_872412301619614653_n.png?_nc_cat=0&oh=5b8eb9a9826662429339a57bc9e94bd6&oe=5C27AA03",
			"https://scontent-eze1-1.xx.fbcdn.net/v/t1.0-9/21766449_356108781511943_2077038482413584533_n.png?_nc_cat=0&oh=4b736a2d81a33775a8f9135b2ee3c10c&oe=5C21CB6B",
			"https://scontent-eze1-1.xx.fbcdn.net/v/t1.0-9/30124133_434089967047157_870403579823885059_n.jpg?_nc_cat=0&oh=71c825feae6d72cdb8c4f6bf7a02d6fe&oe=5C313279")
		params.Attachments = append(params.Attachments, slack.Attachment{ImageURL: img})
		match = true

	case stringMatch(msg.Text, "ojo", "ojito", "willy"):
		rta = "*Big Brother is Watching You.*"
		params.Attachments = append(params.Attachments, slack.Attachment{ImageURL: "https://i.imgur.com/5KASKiY.png"})
		match = true

	}

	if match {
		log.Printf("[DEBUG] theBotSay: %v - %v\n", rta, img)
		//rtm.SendMessage(rtm.NewOutgoingMessage(rta, msg.Channel))
		rtm.PostMessage(msg.Channel, rta, params)
	}

}

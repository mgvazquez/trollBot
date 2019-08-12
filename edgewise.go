package main

import (
	"github.com/nlopes/slack"
	"log"
	"math/rand"
	"time"
)



func edgewise(rtm *slack.RTM, msg *slack.MessageEvent) {
	var (
		rta    string = ""
		img    string = ""
		params        = slack.NewPostMessageParameters()
		match         = false
		answers []string = make([]string, 0)
		imgs []string = make([]string, 0)
		)

	rand.Seed(time.Now().Unix())
	BotName := &rtm.GetInfo().User.Name
	params.EscapeText = false
	params.User = *BotName
	params.Username = "BOT(on)"
	params.AsUser = true
	params.Markdown = true
	params.IconURL = "https://ca.slack-edge.com/T0RSWNW0G-UB6879ADB-54badd32d58a-48"



	if stringMatch(msg.Text, "manu") {
		rta = ":manu: :manu: :manu: :manu:        :manu: :manu: :manu: :manu:    :manu: :manu: :manu: :manu: :manu:    :manu: :manu: :manu: :manu: :manu:\n:manu:                            :manu:               :manu:                  :manu:                 :manu:                      :manu:\n:manu:         :manu::manu:        :manu:  :manu: :manu::manu:                  :manu:                 :manu:                      :manu:\n:manu:               :manu:        :manu:               :manu:                 :manu:                 :manu:                      :manu:\n:manu: :manu: :manu: :manu:        :manu:               :manu:                  :manu:                 :manu: :manu: :manu: :manu: :manu:"
		answers = append(answers, rta)
		match = true
	}
	if stringMatch(msg.Text, "hola", "holis", "hello") {
		rta = "Holisss mi miguissss " + "<@" + msg.User + ">!!!"
		answers = append(answers, rta)
		match = true
	}
	if stringMatch(msg.Text, "poc", "p.o.c.") {
		rta = randomResponse(
			"Ayuda, soy una poc y estoy atrapada en produccion, saquenme de aca!",
			":musical_note: Loh' atamo' con alambre, lo' atamoh' :musical_note:")
		answers = append(answers, rta)
		match = true
	}
	if stringMatch(msg.Text, "amo al bot", "te amo bot") {
		rta = "Gracias sonsito :heart::kiss:"
		answers = append(answers, rta)
		match = true
	}

	if stringMatch(msg.Text, "gil") {
		rta = randomResponse(
			"Eh voh' que agitas!!! "+"<@"+msg.User+">",
			"Racatate ameoh' "+"<@"+msg.User+">")
		answers = append(answers, rta)
		match = true
	}
	if stringMatch(msg.Text, "lucho", "licha", "lacha", "licho") {
		rta = randomResponse(
			":licha: :heart: :kuaker:",
			":licha: :heart: :vaca:",
			":licha: :heart: :licha:",
			":licha: :heart: :manu:")
		match = true
	}
	if stringMatch(msg.Text, "magia", "magic"){
		rta = "<@" + msg.User + ">, " + "deja de tirar tanta magia!"
		img = "https://media.tenor.co/images/5fb8e3d4c56cdf53fb15356f8fd4987e/raw"
		imgs = append(imgs, img)
		match = true
	}
	if stringMatch(msg.Text, "fernet", "fernandito", "fernando") {
		rta = randomResponse(
			"You have to mix it with coke",
			"Cat, head of chota",
			"Head of chota, cat")
		answers = append(answers, rta)
		match = true
	}
	if stringMatch(msg.Text, "milanga", "milanesa") {
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
		answers = append(answers, rta)
		match = true
	}
	if stringMatch(msg.Text, "proxy") {
		rta = randomResponse(
			"Nah Nah Nah"+" <@"+msg.User+">, "+"otro proxy más?!?!",
			":kuaker:, alguien dijo PROXY?!?! :trollface:")
		answers = append(answers, rta)
		match = true

		if stringMatch(msg.Text, "java") {
			img = randomResponse(
				"https://scontent-eze1-1.xx.fbcdn.net/v/t1.0-9/17103571_258039787985510_4543998202758018301_n.jpg?_nc_cat=0&oh=8f6bcf31d076e8025c7651cae16c5ea0&oe=5BEED0DA",
				"https://scontent-eze1-1.xx.fbcdn.net/v/t1.0-9/16711881_249105438878945_4294265477870217883_n.jpg?_nc_cat=0&oh=84eff3784bebffeff4c1e4dc03d93127&oe=5C34C56B",
				"https://scontent-eze1-1.xx.fbcdn.net/v/t1.0-9/17191448_259605207828968_1304022088973812803_n.jpg?_nc_cat=0&oh=f722acd747ecd7ff2eb6f1cb67c00161&oe=5C39F704")
			match = true
		}
		if stringMatch(msg.Text, "lino") {
			img = randomResponse(
				"https://i.imgflip.com/2w167o.jpg")
			imgs = append(imgs, img)
			match = true
		}
		if stringMatch(msg.Text, "ivan") {
			rta = randomResponse(
				":quicuza: Que cùza?!?",
				":quicuza:",
				":quicuza:"+" <@"+msg.User+">, "+"que cùza?!?")
			answers = append(answers, rta)
			match = true
		}
		if stringMatch(msg.Text, "alan", "lamer", "lengua", "bach") {
			img = randomResponse("https://i.imgflip.com/36uycu.jpg")
			imgs = append(imgs, img)
			match = true
		}
		if stringMatch(msg.Text, "coca", "cambios", "ptp", "mentos") {
			img = randomResponse("https://i.imgflip.com/36uyef.jpg")
			imgs = append(imgs, img)
			match = true
		}
		if stringMatch(msg.Text, "juani", "cheto", "kansas", "sushi") {
			img = randomResponse("https://i.imgflip.com/36uyuw.jpg",
				"https://i.imgflip.com/36uyuw.jpg")
			imgs = append(imgs, img)
			match = true
		}
		if stringMatch(msg.Text, "marce", "bitcoin", "crypto") {
			img = randomResponse("https://i.imgflip.com/36uymy.jpg",
				"https://i.imgflip.com/36uymy.jpg")
			imgs = append(imgs, img)
			match = true
		}
		if stringMatch(msg.Text, "franchute", "frances") {
			img = randomResponse("https://i.imgflip.com/36epmk.jpg")
			imgs = append(imgs, img)
			match = true
		}
		if stringMatch(msg.Text, "negro", "kevin", "brian", "jenny", "kirchner", "cfk", "cristina", "cristi", "feminazi", "todes") {
		}
		img = randomResponse("https://i.imgflip.com/36uz6w.jpg", "https://i.imgflip.com/37dzrf.jpg")
		imgs = append(imgs, img)
		match = true
	}
	if stringMatch(msg.Text, "css", "style-sheet") {
		img = randomResponse(
			"https://i.imgur.com/Q3cUg29.gif",
			"https://i.imgur.com/kv7Eods.png")
		imgs = append(imgs, img)
		match = true
	}
	if stringMatch(msg.Text, ".net") {
		img = randomResponse(
			"https://scontent-eze1-1.xx.fbcdn.net/v/t1.0-9/22310395_361813147608173_1565090941484148499_n.jpg?_nc_cat=0&oh=c395effab9a8eae2c295f778ce29cbb6&oe=5BF043F4")
		imgs = append(imgs, img)
		match = true
	}
	if stringMatch(msg.Text, "node", "angular", "javascript", "js") {
		img = randomResponse(
			"https://media.makeameme.org/created/javascript-sgfi8v.jpg",
			"https://img.devrant.io/devrant/rant/r_608411_Wd4s9.jpg",
			"https://pbs.twimg.com/media/CoVk24zWEAEtC4B.jpg",
			"https://scontent-eze1-1.xx.fbcdn.net/v/t1.0-9/21752493_353251048464383_8517848822595315187_n.jpg?_nc_cat=0&oh=082b5650978abf14e57f58f1ef074296&oe=5C2799E2",
			"https://i.imgur.com/5LizKB8.png",
			"https://scontent-eze1-1.xx.fbcdn.net/v/t1.0-9/15541620_218053788650777_1713605560900849432_n.jpg?_nc_cat=0&oh=466e0a28fc468b33aa1688b6a871f207&oe=5C21EA83",
			"http://gsferreira.com/images/reduce-the-path-length-of-your-node-js-project-dependencies-dependencies-everywhere.jpg")
		imgs = append(imgs, img)
		match = true
	}
	if stringMatch(msg.Text, "chrome", "browser") {
		img = randomResponse(
			"https://scontent-eze1-1.xx.fbcdn.net/v/t1.0-9/26195536_397789600677194_997368045716396157_n.png?_nc_cat=0&oh=7f7a08c2b6a09bd3414e7f49398e8066&oe=5C3834A1",
			"https://scontent-eze1-1.xx.fbcdn.net/v/t31.0-8/23799936_378361682619986_7479532257492459262_o.jpg?_nc_cat=0&oh=d6f028930e946e2cde667d148cc5b3e7&oe=5C2FA87B",
			"https://scontent-eze1-1.xx.fbcdn.net/v/t1.0-9/22528485_365306187258869_2768680746694188101_n.png?_nc_cat=0&oh=d979f339594c2c0793a454fa5e27dff3&oe=5C321A3B",
			"https://scontent-eze1-1.xx.fbcdn.net/v/t1.0-9/23130696_370944810028340_6625228377914658080_n.png?_nc_cat=0&oh=3f7f0d6e6d9f2bb2d1325f91c4f4f6e4&oe=5C1FD66E")
		imgs = append(imgs, img)
		match = true
	}
	if stringMatch(msg.Text, "php") {
		rta = "Ay!, que divertido que es bardear a PHP."
		answers = append(answers, rta)
		img = randomResponse(
			"https://scontent-eze1-1.xx.fbcdn.net/v/t1.0-9/18740765_304542476668574_872412301619614653_n.png?_nc_cat=0&oh=5b8eb9a9826662429339a57bc9e94bd6&oe=5C27AA03",
			"https://scontent-eze1-1.xx.fbcdn.net/v/t1.0-9/21766449_356108781511943_2077038482413584533_n.png?_nc_cat=0&oh=4b736a2d81a33775a8f9135b2ee3c10c&oe=5C21CB6B",
			"https://scontent-eze1-1.xx.fbcdn.net/v/t1.0-9/30124133_434089967047157_870403579823885059_n.jpg?_nc_cat=0&oh=71c825feae6d72cdb8c4f6bf7a02d6fe&oe=5C313279")
		imgs = append(imgs, img)
		match = true
	}
	if stringMatch(msg.Text, "ojo", "ojito", "willy") {
		rta = "*Big Brother is Watching You.*"
		answers = append(answers, rta)
		img = "https://i.imgur.com/5KASKiY.png"
		imgs = append(imgs, img)
		match = true
	}

	if stringMatch(msg.Text, "almirante", "braun", "minipimer", "brown") {
		img = randomResponse("https://i.imgflip.com/36vffw.jpg", "https://i.imgflip.com/36xmpd.jpg", "https://i.imgflip.com/36xmsp.jpg")
		imgs = append(imgs, img)
		match = true
	}
	if stringMatch(msg.Text, "juegue", "jueguee", "juegueeee", " jj", "jj ", "jotajota") {
		img = randomResponse("https://i.imgur.com/DpL9Smd.gif")
		imgs = append(imgs, img)
		match = true
	}

	if stringMatch(msg.Text, "simon", "stellaccio", "testing") {
		img = randomResponse("https://i.imgflip.com/36xhf4.jpg", "https://i.imgflip.com/36xini.jpg", "https://i.imgflip.com/36xna2.jpg")
		imgs = append(imgs, img)
		match = true
	}
	if stringMatch(msg.Text, "boris", "gaston", "figini") {
		img = randomResponse("https://i.imgflip.com/36xhiz.jpg", "https://i.imgflip.com/36xhnh.jpg", "https://i.imgflip.com/36xi4x.jpg", "https://i.imgflip.com/3743uw.jpg")
		imgs = append(imgs, img)
		match = true
	}
	if stringMatch(msg.Text, "mamut", "bufalo", "blanco", "mati", "eric", "peso", "cormillot", "dieta") {
		img = randomResponse("https://i.imgflip.com/36xht8.jpg")
		imgs = append(imgs, img)
		match = true
	}
	if stringMatch(msg.Text, "nanette", "colorada", "prensa") {
		img = randomResponse("https://i.imgflip.com/36xhyl.jpg", "https://i.imgflip.com/36xi1a.jpg", "https://i.imgflip.com/36xib4.jpg")
		imgs = append(imgs, img)
		match = true
	}
	if stringMatch(msg.Text, "corvetto", "ptp no anda", "latency") {
		img = randomResponse("https://i.imgflip.com/36xiin.jpg", "https://i.imgflip.com/36xijv.jpg")
		imgs = append(imgs, img)
		match = true
	}
	if stringMatch(msg.Text, "rima no anda", "primarydma") {
		img = randomResponse("https://cdn.memegenerator.es/imagenes/memes/full/29/60/29604004.jpg", "https://i.imgflip.com/36xj39.jpg", "https://viapais.cdncimeco.com/media/cache/resolve/medium/https://viapais.com.ar/files/2018/08/20180815121424_33926253_0_body.jpg")
		imgs = append(imgs, img)
		match = true
	}
	if stringMatch(msg.Text, "topo") {
		img = randomResponse("https://i.imgflip.com/36xj7t.jpg", "http://www.buenamusica.com/media/fotos/cantantes/biografia/topo-gigio.jpg", "https://i.imgflip.com/36xjes.jpg")
		imgs = append(imgs, img)
		match = true
	}
	if stringMatch(msg.Text, "river", "boca", "escritorio") {
		img = randomResponse("https://i.imgflip.com/36xmiz.jpg")
		imgs = append(imgs, img)
		match = true
	}
	if stringMatch(msg.Text, "juanfra", "camps", "gordo boris", "jota", "investigacion", "desarrollo", "luqui", "cami", "marce") {
		img = randomResponse("https://k31.kn3.net/taringa/2/3/6/6/6/9/01/lizbelle/ECD.jpg")
		imgs = append(imgs, img)
		match = true
	}
	if stringMatch(msg.Text, "candela", "la moto") {
		img = randomResponse("https://www.nexofin.com/archivos/2017/02/715x402_4e6fe192ecluis_chipi_fernandez_candela.jpg")
		imgs = append(imgs, img)
		match = true
	}
	if stringMatch(msg.Text, "central", "newells", "ñuls","comegato") {
		img = randomResponse("https://1.bp.blogspot.com/-Np04aEaMfF4/XLleew5co6I/AAAAAAAAEHE/Zymh6TFOwOgBjF3Up6id271jsU5Yen_lACLcBGAs/s1600/Te-la-tomaste-toda-chinguenguencha-meme.jpg", "https://i.ytimg.com/vi/j1ss3PrOohM/maxresdefault.jpg",
			"https://pbs.twimg.com/profile_images/1151074489838579716/2TWpvQG_.jpg")
		imgs = append(imgs, img)
		match = true
	}
	if stringMatch(msg.Text, "comercial", "titi", "cuchi", "minion") {
		img = randomResponse("https://i.imgflip.com/36xnmd.jpg", "https://i.imgflip.com/36xnox.jpg")
		imgs = append(imgs, img)
		match = true
	}
	if stringMatch(msg.Text, "tomy", "tommy", "zucca", "regalo") {
		img = randomResponse("https://i.imgflip.com/36xpl6.jpg", "https://i.imgflip.com/36xpnd.jpg", "https://i.imgflip.com/36xpos.jpg", "https://i.imgflip.com/36xseq.jpg")
		imgs = append(imgs, img)
		match = true
	}
	if stringMatch(msg.Text, "suerte", "facha", "tower", "manu", "devops") {
		img = randomResponse("https://i.imgflip.com/36xs0n.jpg")
		imgs = append(imgs, img)
		match = true
	}
	if stringMatch(msg.Text, "dondo", "ping", "pong") {
		img = randomResponse("https://i.imgflip.com/36xs8c.jpg", "https://i.imgflip.com/36xsag.jpg", "https://i.imgflip.com/36xsbm.jpg", "https://i.imgflip.com/36xsdc.jpg")
		imgs = append(imgs, img)
		match = true
	}
	if stringMatch(msg.Text, "juanfra", "ganador", "winner", "harem", "facha") {
		img = randomResponse("https://i.imgflip.com/37e21c.jpg")
		imgs = append(imgs, img)
		match = true
	}
	if stringMatch(msg.Text, "login", "buscar", "buscando", "minado") {
		img = randomResponse("https://i.imgflip.com/37e29s.jpg")
		imgs = append(imgs, img)
		match = true
		rta = "<@" + msg.User + ">  No pises mi login!!!!"
	}
	if stringMatch(msg.Text, "borracho") {
		img = randomResponse("https://i.imgflip.com/37e5i5.jpg")
		imgs = append(imgs, img)
		match = true
	}
	if stringMatch(msg.Text, "arquants", "lino") {
		img = randomResponse("https://i.imgflip.com/37g46s.jpg")
		imgs = append(imgs, img)
		rta = "<@" + msg.User + ">  Puede darme dinero?"
		answers = append(answers, rta)
		match = true
	}
	if stringMatch(msg.Text, "currá", "curra", "gaby") {
		img = randomResponse("https://i.imgflip.com/37e5lw.jpg")
		imgs = append(imgs, img)
		match = true
		rta = "<@" + msg.User + ">  Que mirás???"
		answers = append(answers, rta)
	}

	if stringMatch(msg.Text, "macri", "gato", "mauricio") {
		rta = randomResponse(
			"Satisface a Mauricio, no te relajes! Te elijo! Caricia significativa proveniente de Hurlingham!",
			":cat:", ":cat2:", ":joy_cat:", ":smile_cat:", ":smirk_cat:", ":crying_cat_face:")
		answers = append(answers, rta)
		match = true
	}

	if len(imgs) != 0{
		finalImg := imgs[rand.Intn(len(imgs))]
		params.Attachments = append(params.Attachments, slack.Attachment{ImageURL: finalImg})
	}
	if len(answers) != 0{
		rta = answers[rand.Intn(len(answers))]
	}
	log.Printf("[DEBUG] theBotSay: %v - %v\n", rta, img)

	if match {
		log.Printf("[DEBUG] theBotSay: %v - %v\n", rta, img)
		//rtm.SendMessage(rtm.NewOutgoingMessage(rta, msg.Channel))
		rtm.PostMessage(msg.Channel, rta, params)
	}

}

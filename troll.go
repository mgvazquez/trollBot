package main

import (
	"os"
	"log"

	"github.com/nlopes/slack"
	"github.com/kelseyhightower/envconfig"
)

type envConfig struct {
	// BotToken is bot user token to access to slack API.
	BotToken string `envconfig:"SLACK_TOKEN" required:"true"`
}

func main() {
	os.Exit(_main(os.Args[1:]))
}

func _main(args []string) int {

	var env envConfig

	if err := envconfig.Process("", &env); err != nil {
		log.Printf("[ERROR] Failed to process env var: %s", err)
		return 1
	}

	log.Printf("[INFO] Start slack event listening")

	api := slack.New(env.BotToken)
	api.SetDebug(false)

	rtm := api.NewRTM()
	go rtm.ManageConnection()

	//TODO: Agregar un Handler que mande mensajes cuando el channel al que esta suscripto este muy calmado

	Loop:
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
				case *slack.ConnectedEvent:
					log.Println("[INFO] Connection counter:", ev.ConnectionCount)

				case *slack.MessageEvent:
					log.Printf("[INFO] Message (%v): %v\n", ev.Msg.User, ev.Msg.Text)
					listenAndRespond(rtm, ev)

				case *slack.RTMError:
					log.Printf("[ERROR] %s\n", ev.Error())

				case *slack.InvalidAuthEvent:
					log.Printf("[ERROR] Invalid credentials")
					break Loop

				default:
					log.Printf("[DEBUG] Event (%v) received.\n", msg.Type)
			}
		}
	}
	return 0
}
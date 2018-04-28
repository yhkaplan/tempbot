package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/d2r2/go-dht"
	"github.com/nlopes/slack"
)

// I wrote this for the DHT11 model, but other hardware may be compatible
const (
    sensorType = dht.DHT11
)

func main() {

	token := os.Getenv("SLACK_TOKEN")
	api := slack.New(token)
	rtm := api.NewRTM()

	go rtm.ManageConnection()

Loop:

	for {
		select {
		case msg := <-rtm.IncomingEvents:
			fmt.Println("Event received: ")
			switch ev := msg.Data.(type) {

			case *slack.ConnectedEvent:
				fmt.Println("Connection counter:", ev.ConnectionCount)

			case *slack.MessageEvent:
				fmt.Printf("Message: %v\n", ev)
				info := rtm.GetInfo()
				prefix := fmt.Sprintf("<@%s>", info.User.ID)

				if ev.User != info.User.ID && strings.HasPrefix(ev.Text, prefix) {
					respond(rtm, ev, prefix)
				}

			case *slack.RTMError:
				fmt.Printf("Error: %s\n", ev.Error())

			case *slack.InvalidAuthEvent:
				fmt.Printf("Invalid credentials")
				break Loop

			default:
				// Do nothin'
			}
		}
	}
}

// Responds to request for temp and humidity value in slack
func respond(rtm *slack.RTM, msg *slack.MessageEvent, prefix string) {
	rtm.SendMessage(rtm.NewOutgoingMessage("Checking temp", msg.Channel))

	var msgTxt string 
	temp, humid, err := getTemp()

	if err != nil {
		fmt.Println(err)
		msgTxt = fmt.Sprintf("Error occurred: %s", err)
	} else {
		//TODO: Find way to safely remove space between % sign and value for humidity
		msgTxt = fmt.Sprintf("Temperature: %v \u00B0C, Humidity: %v %%", temp, humid)
	}

	rtm.SendMessage(rtm.NewOutgoingMessage(msgTxt, msg.Channel))
}

// Gets temp and humidity from DHT11 sensor
func getTemp() (float32, float32, error) {
	temp, humid, _, err := dht.ReadDHTxxWithRetry(sensorType, 4, false, 5)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Temperature: %v, Humidity: %v", temp, humid)
	return temp, humid, err
}

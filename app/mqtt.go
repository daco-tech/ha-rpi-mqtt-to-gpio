package app

import (
	"RaspberryPiMQTTHead/general"
	"fmt"
	"log"
	"net/url"
	"strconv"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	rpio "github.com/stianeikeland/go-rpio"
)

func Connect(clientId string, uri *url.URL) mqtt.Client {
	opts := createClientOptions(clientId, uri)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	return client
}

func createClientOptions(clientId string, uri *url.URL) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", uri.Host))
	opts.SetUsername(uri.User.Username())
	password, _ := uri.User.Password()
	opts.SetPassword(password)
	opts.SetClientID(clientId)
	return opts
}

func Listen(mqttClient mqtt.Client, config general.Mqtt2Gpio, pin rpio.Pin, verbose bool) {
	//Send Status
	mqttClient.Publish(config.Mqtt.StateTopic, config.Mqtt.Qos, config.Mqtt.StateRetained, config.Mqtt.PayloadAvailable)

	log.Println("--> Setting Boot Status for Device '" + config.Name + "' to High: " + strconv.FormatBool(config.OnBootHigh))
	if config.OnBootHigh {
		pin.High()
		mqttClient.Publish(config.Mqtt.StateTopic, config.Mqtt.Qos, config.Mqtt.StateRetained, config.Mqtt.PayloadOff)
	} else {
		pin.Low()
		mqttClient.Publish(config.Mqtt.StateTopic, config.Mqtt.Qos, config.Mqtt.StateRetained, config.Mqtt.PayloadOn)
	}

	mqttClient.Subscribe(config.Mqtt.CommandTopic, config.Mqtt.Qos, func(client mqtt.Client, msg mqtt.Message) {
		if verbose {
			log.Printf("* [%s] %s\n", msg.Topic(), string(msg.Payload()))
		}
		if string(msg.Payload()) == config.Mqtt.PayloadOn {
			log.Println("Turning " + config.Name + " " + config.Mqtt.PayloadOn + "...")
			// Turn on Relay
			pin.Low()
			mqttClient.Publish(config.Mqtt.StateTopic, config.Mqtt.Qos, config.Mqtt.StateRetained, config.Mqtt.PayloadOn)
		} else {
			log.Println("Turning " + config.Name + " " + config.Mqtt.PayloadOff + "...")
			// Turn OFF Relay
			pin.High()
			mqttClient.Publish(config.Mqtt.StateTopic, config.Mqtt.Qos, config.Mqtt.StateRetained, config.Mqtt.PayloadOff)
		}
	})
}

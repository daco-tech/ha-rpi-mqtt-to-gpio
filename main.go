package main

import (
	"RaspberryPiMQTTHead/app"
	"RaspberryPiMQTTHead/general"
	"log"
	"net/url"
	"os"
	"strconv"
	"time"

	rpio "github.com/stianeikeland/go-rpio"
)

func main() {
	log.Println("Raspberry Pi MQTT HEAD for HomeAssistant")
	log.Println("Starting App...")
	log.Println("--> Loading Configs...")
	config, _ := general.LoadConfiguration("/go/bin/config.json")
	if config.Log.Verbose {
		log.Println("*** VERBOSE LOG ACTIVATED ***")
	}

	// Prepare Raspberry Pi GPIO
	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	// Unmap gpio memory when done
	defer rpio.Close()

	// Connect with MQTT Broker
	uri, err := url.Parse(config.Mqtt.Host)
	if err != nil {
		log.Fatal(err)
	}
	mqttClient := app.Connect(config.Mqtt.ClientID, uri)

	// FROM MQTT to GPIO Instructions
	for _, fromMQTT := range config.Gpio.Mqtt2Gpio {
		// Configure output gpio Pins
		if config.Log.Verbose {
			log.Println("Configure Pin " + strconv.Itoa(fromMQTT.Pin) + " for device '" + fromMQTT.Name + "' mode Output...")
		}
		pin := rpio.Pin(fromMQTT.Pin)
		pin.Output()
		//Initialize Listners
		if config.Log.Verbose {
			log.Println("Configure MQTT Listener for device '" + fromMQTT.Name + "' on topic: '" + fromMQTT.Mqtt.CommandTopic + "'...")
		}
		go app.Listen(mqttClient, fromMQTT, pin, config.Log.Verbose)
	}

	// FROM GPIO to MQTT Instructions
	states := make([]bool, len(config.Gpio.Gpio2Mqtt))
	for {
		for i, toMQTT := range config.Gpio.Gpio2Mqtt {
			mqttClient.Publish(toMQTT.Mqtt.AvailabilityTopic, toMQTT.Mqtt.Qos, toMQTT.Mqtt.AvailabilityRetained, toMQTT.Mqtt.PayloadAvailable)
			pin := rpio.Pin(toMQTT.Pin)
			pin.Input()

			res := pin.Read()
			if int(res) == 0 {
				mqttClient.Publish(toMQTT.Mqtt.StateTopic, toMQTT.Mqtt.Qos, toMQTT.Mqtt.StateRetained, toMQTT.Mqtt.PayloadOff)
				if states[i] {
					states[i] = false
					log.Println("Device '" + toMQTT.Name + "' state: " + toMQTT.Mqtt.PayloadOff)
				}

			} else {
				mqttClient.Publish(toMQTT.Mqtt.StateTopic, toMQTT.Mqtt.Qos, toMQTT.Mqtt.StateRetained, toMQTT.Mqtt.PayloadOn)
				if !states[i] {
					states[i] = true
					log.Println("Device '" + toMQTT.Name + "' state: " + toMQTT.Mqtt.PayloadOn)
				}
			}
		}
		time.Sleep(time.Duration(config.System.LoopIntervalSec) * time.Second)
	}

}

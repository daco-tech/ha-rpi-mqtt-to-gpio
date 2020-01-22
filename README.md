# Home Assistant remote integration with Raspberry Pi GPIO using MQTT
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fdaco-tech%2Fraspberrypi-mqtt-head.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fdaco-tech%2Fraspberrypi-mqtt-head?ref=badge_shield)


This is a simple Go program is to easly integrate a remote Home Assistant (Installed on a different computer than the Raspberry Pi) with Raspberry Pi's GPIO using MQTT to communicate.

This project fits my home needs so, please feel free to suggest changes and improvements, as well as contributing to this project.

## Getting started

* Make sure you have [dep](https://github.com/golang/dep) installed
* Clone this repo `git clone https://github.com/daco-tech/ha-rpi-mqtt-to-gpio.git`
* Create a configuration file with name: config.json at main.go file level with the text in the config section
* Run `make` to download dependencies and run the application
* Run `env GOOS=linux GOARCH=arm GOARM=5 go build main.go` to build the binary to run on the RaspberryPi


## Config

config.json file content example:
```
{
    "mqtt": {
        "host": "mqtt://<ip address>:1883",
        "clientId": "rpihead"
    },
    "log": {
        "verbose": true
    },
    "system": {
        "loop_interval_sec": 1
    },
    "gpio": {
        "mqtt2gpio": [
            {
                "name": "Boiler",
                "pin": 23,
                "on_boot_high": true,
                "mqtt": {
                    "qos": 0,
                    "command_topic": "home/boiler/switch/set",
                    "command_retained": true,
                    "state_topic": "home/boiler/switch/state",
                    "state_retained": true,
                    "payload_available": "online",
                    "payload_not_available": "offline",
                    "payload_on": "ON",
                    "payload_off": "OFF"
                }
            }
        ],
        "gpio2mqtt": [
            {
                "name": "Front Door",
                "pin": 27,
                "mqtt": {
                    "qos": 0,
                    "state_topic": "home/door/front/state",
                    "state_retained": true,
                    "payload_on": "closed",
                    "payload_off": "open",
                    "availability_topic": "home/door/front/status",
                    "availability_retained": true,
                    "payload_available": "online",
                    "payload_not_available": "offline"
                }
            }
        ]
    }
}
```

## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fdaco-tech%2Fraspberrypi-mqtt-head.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fdaco-tech%2Fraspberrypi-mqtt-head?ref=badge_large)
package general

type Config struct {
	Mqtt struct {
		Host     string `json:"host"`
		ClientID string `json:"clientId"`
	} `json:"mqtt"`
	Log struct {
		Verbose bool `json:"verbose"`
	} `json:"log"`
	System struct {
		LoopIntervalSec int `json:"loop_interval_sec"`
	} `json:"system"`
	Gpio struct {
		Mqtt2Gpio []struct {
			Name       string `json:"name"`
			Pin        int    `json:"pin"`
			OnBootHigh bool   `json:"on_boot_high"`
			Mqtt       struct {
				Qos                 byte   `json:"qos"`
				CommandTopic        string `json:"command_topic"`
				CommandRetained     bool   `json:"command_retained"`
				StateTopic          string `json:"state_topic"`
				StateRetained       bool   `json:"state_retained"`
				PayloadAvailable    string `json:"payload_available"`
				PayloadNotAvailable string `json:"payload_not_available"`
				PayloadOn           string `json:"payload_on"`
				PayloadOff          string `json:"payload_off"`
			} `json:"mqtt"`
		} `json:"mqtt2gpio"`
		Gpio2Mqtt []struct {
			Name string `json:"name"`
			Pin  int    `json:"pin"`
			Mqtt struct {
				Qos                  byte   `json:"qos"`
				StateTopic           string `json:"state_topic"`
				StateRetained        bool   `json:"state_retained"`
				PayloadOn            string `json:"payload_on"`
				PayloadOff           string `json:"payload_off"`
				AvailabilityTopic    string `json:"availability_topic"`
				AvailabilityRetained bool   `json:"availability_retained"`
				PayloadAvailable     string `json:"payload_available"`
				PayloadNotAvailable  string `json:"payload_not_available"`
			} `json:"mqtt"`
		} `json:"gpio2mqtt"`
	} `json:"gpio"`
}

type Mqtt2Gpio struct {
	Name       string `json:"name"`
	Pin        int    `json:"pin"`
	OnBootHigh bool   `json:"on_boot_high"`
	Mqtt       struct {
		Qos                 byte   `json:"qos"`
		CommandTopic        string `json:"command_topic"`
		CommandRetained     bool   `json:"command_retained"`
		StateTopic          string `json:"state_topic"`
		StateRetained       bool   `json:"state_retained"`
		PayloadAvailable    string `json:"payload_available"`
		PayloadNotAvailable string `json:"payload_not_available"`
		PayloadOn           string `json:"payload_on"`
		PayloadOff          string `json:"payload_off"`
	} `json:"mqtt"`
}

type Gpio2Mqtt struct {
	Name string `json:"name"`
	Pin  int    `json:"pin"`
	Mqtt struct {
		Qos                  byte   `json:"qos"`
		StateTopic           string `json:"state_topic"`
		StateRetained        bool   `json:"state_retained"`
		PayloadOn            string `json:"payload_on"`
		PayloadOff           string `json:"payload_off"`
		AvailabilityTopic    string `json:"availability_topic"`
		AvailabilityRetained bool   `json:"availability_retained"`
		PayloadAvailable     string `json:"payload_available"`
		PayloadNotAvailable  string `json:"payload_not_available"`
	} `json:"mqtt"`
}

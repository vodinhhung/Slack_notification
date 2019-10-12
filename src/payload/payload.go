package payload

import "os"

type Payload struct {
	Channel string  `json:"channel"`
	Blocks  []Block `json:"blocks"`
}

func NewPayload() Payload {
	block := Block1()

	channel, _ := os.LookupEnv("CHANNEL_ID")
	payload := Payload{
		Channel: channel,
		Blocks:  []Block{block},
	}

	return payload
}

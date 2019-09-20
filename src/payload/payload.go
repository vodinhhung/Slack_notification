package payload

type Payload struct {
    Channel string `json:"channel"`
    Blocks []Block `json:"blocks"`
}

func NewPayload() Payload {
    block := Block1();

    payload := Payload{
        Channel:    "CNBSZ3ABX",
        Blocks:     []Block{block},
    }

    return payload
}
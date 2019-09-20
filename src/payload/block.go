package payload

type Block struct {
	Type    string          `json:"type"`
	Text    *Component      `json:"text"`
	Fields  *[]Component    `json:"fields"`
}

type Component struct {
	Type    string		`json:"type"`
	Text    string    	`json:"text"`
}

func Block1() Block{
	component := Component{
		Type: "mrkdwn",
		Text: "*Type:*\nComputer (laptop)",
	}

	block := Block{
		Type:   "section",
		Text:   &component,
		Fields: nil,
	}

	return block
}

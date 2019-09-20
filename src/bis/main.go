package main

import (
	"../payload"
	"encoding/json"
	"fmt"
)

func main() {

	//url := "https://slack.com/api/chat.postMessage"
	//
	//payload := strings.NewReader("{\n\t\"channel\": \"CNBSZ3ABX\",\n\t\"blocks\": [\n\t{\n\t\t\"type\": \"section\",\n\t\t\"text\": {\n\t\t\t\"type\": \"mrkdwn\",\n\t\t\t\"text\": \"You have a new request:\\n*<fakeLink.toEmployeeProfile.com|Fred Enriquez - New device request>*\"\n\t\t}\n\t},\n\t{\n\t\t\"type\": \"section\",\n\t\t\"fields\": [\n\t\t\t{\n\t\t\t\t\"type\": \"mrkdwn\",\n\t\t\t\t\"text\": \"*Type:*\\nComputer (laptop)\"\n\t\t\t},\n\t\t\t{\n\t\t\t\t\"type\": \"mrkdwn\",\n\t\t\t\t\"text\": \"*When:*\\nSubmitted Aut 10\"\n\t\t\t},\n\t\t\t{\n\t\t\t\t\"type\": \"mrkdwn\",\n\t\t\t\t\"text\": \"*Last Update:*\\nMar 10, 2015 (3 years, 5 months)\"\n\t\t\t},\n\t\t\t{\n\t\t\t\t\"type\": \"mrkdwn\",\n\t\t\t\t\"text\": \"*Reason:*\\nAll vowel keys aren't working.\"\n\t\t\t},\n\t\t\t{\n\t\t\t\t\"type\": \"mrkdwn\",\n\t\t\t\t\"text\": \"*Specs:*\\n\\\"Cheetah Pro 15\\\" - Fast, really fast\\\"\"\n\t\t\t}\n\t\t]\n\t}\n]\n}")
	//
	//req, _ := http.NewRequest("POST", url, payload)
	//
	//req.Header.Add("content-type", "application/json")
	//req.Header.Add("authorization", "Bearer xoxp-762117304070-747132412354-756417357443-a3a4732e3ccea62405b0b1c6f6439f01")
	//req.Header.Add("cache-control", "no-cache")
	//req.Header.Add("postman-token", "421af50d-65e5-5b4c-80e0-9787af0930a7")
	//
	//res, err := http.DefaultClient.Do(req)
	//if err != nil {
	//	log.Println(err)
	//}
	//
	//defer res.Body.Close()
	//body, _ := ioutil.ReadAll(res.Body)
	//
	//fmt.Println(res)
	//fmt.Println(string(body))

	payload := payload.NewPayload()

	payloadString, _ := json.Marshal(payload)
	fmt.Println(string(payloadString))

}

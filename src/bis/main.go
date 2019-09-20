package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {

	url := "https://slack.com/api/chat.postMessage"

	payload := strings.NewReader("{\n\t\"channel\": \"CNBSZ3ABX\",\n\t\"text\": \"I am a test message http://slack.com\"\n}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", "Bearer xoxp-762117304070-747132412354-747149232594-c589fa77a55ede11f1623952b8ea0b90")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("postman-token", "421af50d-65e5-5b4c-80e0-9787af0930a7")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}

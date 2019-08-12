package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Data struct {
	Onboarding struct {
		Onboarding    string   `json:"onboarding"`
		Authorization string   `json:"authorization"`
		Body          []string `json:"body"`
		Post          string   `json:"post"`
	} `json:"onboarding"`
	Base64 string `json:"base64"`
}

var Config Data

func main() {
	data, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatalln("Failed to load config.json: ", err)
	}

	if err = json.Unmarshal(data, &Config); err != nil {
		log.Fatalln("Failed to unmarshal Json file: ", err)
	}
}

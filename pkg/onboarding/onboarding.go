package onboarding

import (
	"csv/pkg/config"
	"fmt"
	"log"

	"gopkg.in/resty.v1"
)

func OnboardRequest(id int) error {
	url := config.Config.Onboarding.Onboarding
	uri := fmt.Sprintf(url+"%d/onboard", id)
	res, err := resty.R().
		SetHeader("Authorization", config.Config.Onboarding.Authorization).
		Get(uri)
	if err != nil {
		log.Println(err)
		// log.Println(res.Status)
		return err
	}

	//res.IsError()
	//res.StatusCode
	//res.Status

	log.Println(id, res)
	return err
}

func GetAccessToken() (*resty.Response, error) {
	res, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"` + config.Config.Onboarding.Body[0] + `" : "` + config.Config.Onboarding.Body[1] + `", "` + config.Config.Onboarding.Body[2] + `": "` + config.Config.Onboarding.Body[3] + `"}`).
		Post(config.Config.Onboarding.Post)
	if err != nil {
		log.Panicln(err)
	}
	return res, err
}

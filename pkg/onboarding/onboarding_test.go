package onboarding

import (
	"log"
	"testing"
)

const id int = 1127012

func TestOnboarding(t *testing.T) {
	token, err := GetAccessToken()
	if err != nil {
		t.Error("Request failed: ", err)
	}
	log.Println(token)
}

func TestOnboardRequest(t *testing.T) {
	err := OnboardRequest(id)
	if err != nil {
		t.Error("Request failed: ", err)
		log.Panicln("Request failed: ", err)
		t.FailNow()

	}
}

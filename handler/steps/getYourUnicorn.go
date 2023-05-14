package steps

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

type GetYourUnicorn struct{}

func NewUnicorn() *GetYourUnicorn {
	return &GetYourUnicorn{}
}

type Unicorn struct {
	TailSize         int `json:"tail_size"`
	Rainbownabillity int `json:"rainbownabillity"`
	Glitteryness     int `json:"glitteryness"`
}

func (s *GetYourUnicorn) Execute(body string) (string, error) {
	var info UnicornRequestInfo
	err := json.Unmarshal([]byte(body), &info)
	if err != nil {
		return "", err
	}

	if info.Owner != "Guy" {
		return "", NewUnicornError("I'm sorry, we're out of unicorns!")
	}

	unicorn := Unicorn{
		TailSize:         randomInt(1, 10),
		Rainbownabillity: randomInt(1, 10),
		Glitteryness:     randomInt(1, 10),
	}

	return fmt.Sprintf("Here is your unicorn! %+v", unicorn), nil
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

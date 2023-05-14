package steps

import (
	"encoding/json"
	"fmt"
)

type ProvideUnicornData struct {}

func NewProvideUnicornData() *ProvideUnicornData {
	return &ProvideUnicornData{}
}

type UnicornData struct {
	UnicornName  string `json:"unicorn_name"`
	UnicornColor string `json:"unicorn_color"`
}

func (s *ProvideUnicornData) Execute(body string) (string, error) {
	var data UnicornData
	err := json.Unmarshal([]byte(body), &data)
	if err != nil {
		return "", err
	}

	fmt.Println("Providing unicorn called ", data.UnicornName, " with color ", data.UnicornColor)

	if data.UnicornName == "Sparkles" {
		return "Sparkles is a good unicorn! :)", nil
	}

	if data.UnicornColor == "Pink" {
		return "think i may have a pink Sparkle here...", nil
	}

	return "", NewUnicornError("I'm sorry Guy, we only have Sparkles... :(")
}

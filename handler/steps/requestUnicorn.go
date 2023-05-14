package steps

import (
	"encoding/json"
	"fmt"
)

type UnicornRequest struct {}

func NewUnicornRequest() *UnicornRequest  {
	return &UnicornRequest{}
}

type UnicornRequestInfo struct {
	Owner string `json:"owner"`
}

func (r *UnicornRequest) Execute(body string) (string, error) {
	var request UnicornRequestInfo 
	err := json.Unmarshal([]byte(body), &request)
	if err != nil {
		return "", err
	}

	fmt.Println("Requesting unicorn for", request.Owner)

	if request.Owner == "Guy" {
		return "I can see that you may be a good owner!", nil
	}

	return "", NewUnicornError("I'm sorry, we're out of unicorns!")
}

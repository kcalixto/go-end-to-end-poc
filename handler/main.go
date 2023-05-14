package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go-end-to-end-poc/handler/steps"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(ctx context.Context, event events.APIGatewayProxyRequest) (string, error) {
	requestUnicorn := steps.NewUnicornRequest()
	provideUnicorn := steps.NewProvideUnicornData()
	getUnicorn := steps.NewUnicorn()

	steps := map[string]steps.Steps{
		"request": requestUnicorn,
		"provide": provideUnicorn,
		"get":     getUnicorn,
	}

	var request map[string]interface{}
	err := json.Unmarshal([]byte(event.Body), &request)
	if err != nil {
		panic(err)
	}

	step := request["step"].(string)

	res, err := steps[step].Execute(event.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	return res, nil
}

func main() {
	if env, ok := os.LookupEnv("NODE_ENV"); ok && env == "local" {
		payloadSuccess, err := json.Marshal(map[string]string{
			"step":          "request",
			"owner":         "Guy",
			"unicorn_name":  "Sparkles",
			"unicorn_color": "Pink",
		})
		if err != nil {
			panic(err)
		}
		payloadError, err := json.Marshal(map[string]string{
			"step":          "request",
			"owner":         "Other Guy",
			"unicorn_name":  "Darkness",
			"unicorn_color": "Black",
		})
		if err != nil {
			panic(err)
		}

		payloads := [][]byte{payloadSuccess, payloadError}

		fmt.Println(Handler(context.TODO(), events.APIGatewayProxyRequest{
			Body: string(payloads[1]),
		}))
	} else {
		lambda.Start(Handler)
	}
}

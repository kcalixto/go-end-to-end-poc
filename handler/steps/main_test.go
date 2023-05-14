package steps

import (
	"os"
	"testing"
)

func newSteps() map[string]Steps {
	requestUnicorn := NewUnicornRequest()
	provideUnicorn := NewProvideUnicornData()
	getUnicorn := NewUnicorn()

	return map[string]Steps{
		"request": requestUnicorn,
		"provide": provideUnicorn,
		"get":     getUnicorn,
	}
}

var runStep map[string]Steps

var guy = `{"owner": "Guy"}`
var otherGuy = `{"owner": "Other Guy"}`
var sparkles = `{"unicorn_name": "Sparkles", "unicorn_color": "Pink"}`
var pinky = `{"unicorn_name": "pinky", "unicorn_color": "Pink"}`
var darkness = `{"unicorn_name": "Darkness", "unicorn_color": "Black"}`

var res string
var err error

func TestMain(m *testing.M) {
	runStep = newSteps()

	os.Exit(m.Run())
}

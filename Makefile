run: clean build deploy

clean:
	rm -rf ./bin

build:
	GOOS=linux GOARCH=amd64 go build -o bin/main handler/main.go

deploy:
	sls deploy --verbose

test:
	go run handler/main.go
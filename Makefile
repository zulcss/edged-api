.PHONY: clean api all client

all: clean api client

clean:
	rm -rf bin

api:
	go build -o bin/edged apiserver/main.go

client:
	go build -o bin/stx client/main.go

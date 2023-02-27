.PHONY: clean api all

all: clean api 

clean:
	rm edged.db
	rm -rf bin

api:
	go build -o bin/edged apiserver/main.go

.PHONY: mod tidy vendor build run reflex

mod: tidy vendor

tidy:
	go mod tidy

vendor:
	go mod vendor

build: mod
	go build -o ./mycli ./main.go

run: build
	./mycli

reflex:
	which reflex || go install github.com/cespare/reflex@latest
	reflex --start-service -R '^vendor/' -g go.mod -r '\.go$$' make run

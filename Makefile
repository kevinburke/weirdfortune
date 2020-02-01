.PHONY: install clean

data:
	go-bindata -pkg weirdfortune -o lib/fortunes.go games/weirdfortunes/...

install:
	go get -u github.com/kevinburke/go-bindata/...

build:
	go build ./...

serve:
	go run ./cmd/weirdfortune-server/main.go

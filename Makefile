.PHONY: install clean

install:
	go get -u github.com/kevinburke/go-bindata/...

data:
	go-bindata -o fortunes.go games/weirdfortunes/...

build:
	go build ./...

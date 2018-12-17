.PHONY: install clean

data:
	go-bindata -o fortunes.go games/weirdfortunes/...

install:
	go get -u github.com/kevinburke/go-bindata/...

build:
	go build ./...

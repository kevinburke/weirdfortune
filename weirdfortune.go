package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	_ = flag.Bool("all", false, "Display all fortunes including NSFW ones")
	directory := flag.String("directory", "/usr/local/games", "Path to the games folder")
	flag.Parse()
	if len(flag.Args()) > 0 {
		log.Fatalf("Usage: weirdfortune -all")
	}

	f, err := os.Open(filepath.Join(*directory, "weirdfortunes", "weirdfortunes"))
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(f.Name())
	b := new(bytes.Buffer)
	b.ReadFrom(f)
	fmt.Println(b.String())
}

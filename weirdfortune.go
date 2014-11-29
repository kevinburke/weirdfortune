package main

import (
	"flag"
	"fmt"
	"log"
	"pkg/text/template"
)

func main() {
	_ = flag.Bool("all", false, "Display all fortunes including NSFW ones")
	flag.Parse()
	if len(flag.Args()) > 0 {
		log.Fatalf("Usage: weirdfortune -all")
	}

	f, err := template.ParseFiles("github.com/kevinburke/weirdfortune/games/weirdfortunes/weirdfortunes")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(f.Name())
}

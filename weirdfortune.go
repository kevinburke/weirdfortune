package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	weirdfortune "github.com/kevinburke/weirdfortune/lib"
)

func main() {
	all := flag.Bool("all", false, "Display all fortunes including NSFW ones")
	flag.Parse()
	_ = all
	if len(flag.Args()) > 0 {
		log.Fatalf("Usage: weirdfortune [-all] [-directory DIRECTORY] ")
	}

	choice, err := weirdfortune.Fortune()
	if err != nil {
		log.Fatal(err)
	}
	usertweet := strings.SplitN(choice, ": ", 2)
	if len(usertweet) < 2 {
		fmt.Printf("incorrectly formatted tweet: %v\n", choice)
	}

	out := fmt.Sprintf("\x1b[0;31m%s\x1b[0m\n"+
		"\x1b[0;34m%s\x1b[0m", usertweet[0], usertweet[1])
	fmt.Println(out)
}

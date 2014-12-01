package main

import (
	"bytes"
	"flag"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	all := flag.Bool("all", false, "Display all fortunes including NSFW ones")
	directory := flag.String("directory", "/usr/local/games", "Path to the games folder")
	flag.Parse()
	if len(flag.Args()) > 0 {
		log.Fatalf("Usage: weirdfortune [-all] [-directory DIRECTORY] ")
	}

	f, err := os.Open(filepath.Join(*directory, "weirdfortunes", "weirdfortunes"))
	if err != nil {
		log.Fatalf(err.Error())
	}
	b := new(bytes.Buffer)
	b.ReadFrom(f)

	if *all {
		b.WriteString("%\n")
		f, err = os.Open(filepath.Join(*directory, "weirdfortunes", "nsfwweirdfortunes"))
		if err != nil {
			log.Fatalf(err.Error())
		}
		b.ReadFrom(f)
	}
	// XXX: remove duplicates
	tweets := strings.Split(b.String(), "\n%\n")
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	choice := r.Intn(len(tweets))
	usertweet := strings.SplitN(tweets[choice], ": ", 2)

	os.Stdout.WriteString("\033[0;31m")
	os.Stdout.WriteString(usertweet[0])
	os.Stdout.WriteString("\033[0m")
	os.Stdout.WriteString("\n")
	os.Stdout.WriteString("\033[0;34m")
	os.Stdout.WriteString(usertweet[1])
	os.Stdout.WriteString("\033[0m")
	os.Stdout.WriteString("\n")
}

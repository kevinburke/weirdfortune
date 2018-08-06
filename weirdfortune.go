package main

import (
	"bytes"
	"crypto/rand"
	"flag"
	"fmt"
	"log"
	"math/big"
	mrand "math/rand"
	"strings"
)

func checkError(err error) {
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func main() {
	all := flag.Bool("all", false, "Display all fortunes including NSFW ones")
	flag.Parse()
	if len(flag.Args()) > 0 {
		log.Fatalf("Usage: weirdfortune [-all] [-directory DIRECTORY] ")
	}

	bits, err := Asset("games/weirdfortunes/weirdfortunes")
	checkError(err)
	if *all {
		nsfwbits, err := Asset("games/weirdfortunes/nsfwweirdfortunes")
		checkError(err)
		bits = append(bits, nsfwbits...)
	}

	// XXX: remove duplicates
	tweets := bytes.Split(bits, []byte("\n%\n"))
	i := big.NewInt(2<<32 - 1)
	n, err := rand.Int(rand.Reader, i)
	checkError(err)

	r := mrand.New(mrand.NewSource(n.Int64()))
	choice := r.Intn(len(tweets))
	usertweet := strings.SplitN(string(tweets[choice]), ": ", 2)

	out := fmt.Sprintf("\x1b[0;31m%s\x1b[0m\n"+
		"\x1b[0;34m%s\x1b[0m", usertweet[0], usertweet[1])
	fmt.Println(out)
}

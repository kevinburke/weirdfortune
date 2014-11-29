package main

import (
	"flag"
	"fmt"
	"go/build"
	"log"
	"os"
	"pkg/text/template"
)

func main() {
	_ = flag.Bool("all", false, "Display all fortunes including NSFW ones")
	flag.Parse()
	if len(flag.Args()) > 0 {
		log.Fatalf("Usage: weirdfortune -all")
	}

	fmt.Println(build.Default.GOPATH)
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(wd)
	template.Must(template.ParseFiles("games/weirdfortunes/weirdfortunes"))

}

package main

import (
	"strings"
	"testing"

	weirdfortune "github.com/kevinburke/weirdfortune/lib"
)

func TestAllFortunes(t *testing.T) {
	for i := 0; i < 10000; i++ {
		choice, err := weirdfortune.Fortune()
		if err != nil {
			t.Fatal(err)
		}
		usertweet := strings.SplitN(choice, ": ", 2)
		if len(usertweet) < 2 {
			t.Fatalf("incorrectly formatted tweet: %v", choice)
		}
	}
}

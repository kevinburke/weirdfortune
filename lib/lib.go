package weirdfortune

import (
	"bytes"
	"crypto/rand"
	"math/big"
	mrand "math/rand"
	"sync"
)

var seedOnce sync.Once
var r *mrand.Rand
var randErr error
var tweets [][]byte

func Fortune() (string, error) {
	seedOnce.Do(func() {
		i := big.NewInt(2<<32 - 1)
		n, err := rand.Int(rand.Reader, i)
		if err != nil {
			randErr = err
			return
		}
		r = mrand.New(mrand.NewSource(n.Int64()))
		bits, err := Asset("games/weirdfortunes/weirdfortunes")
		if err != nil {
			randErr = err
			return
		}
		nsfwbits, err := Asset("games/weirdfortunes/nsfwweirdfortunes")
		if err != nil {
			randErr = err
			return
		}
		bits = append(bits, nsfwbits...)
		tweets = bytes.Split(bits, []byte("\n%\n"))
	})
	if randErr != nil {
		return "", randErr
	}
	choice := r.Intn(len(tweets))
	return string(tweets[choice]), nil
}

package randstring

import (
	"math/rand"
	"time"
)

func init() {
    rand.Seed(time.Now().UnixNano())
}

const letterBytes = "+-/<[abcdefghijklmnopqrstuvwxyz](ABCDEFGHIJKLMNOPQRSTUVWXYZ){0123456789}>-_^#"

func GetRandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
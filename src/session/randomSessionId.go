package session

import "math/rand"

func InitRand() {
    rand.Seed(time.Now().UnixNano())
}

const letterBytes = "+-/<[abcdefghijklmnopqrstuvwxyz](ABCDEFGHIJKLMNOPQRSTUVWXYZ){0123456789}>-_^#"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
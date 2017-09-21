package randomizer

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

var runes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GenerateRandomString(n int) string {
	rand.Seed(time.Now().UnixNano())

	r := make([]rune, n)
	for i := range r {
		r[i] = runes[rand.Intn(len(runes))]
	}
	return string(r)
}

func GenerateUUID() (string, error) {
	f, err := os.OpenFile("/dev/urandom", os.O_RDONLY, 0)
	if err != nil {
		return "", err
	}

	b := make([]byte, 16)

	f.Read(b)
	f.Close()

	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid, nil
}

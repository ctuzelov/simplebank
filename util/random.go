package util

import (
	"math/rand"
	"strings"
	"time"
)

const alhpabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random int64
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string with length n
func RandomString(n int) string {
	var str strings.Builder
	k := len(alhpabet)

	for i := 0; i < n; i++ {
		val := alhpabet[rand.Intn(k)]
		str.WriteByte(val)
	}

	return str.String()
}

// RandomOwner generates random owner name
func RandomOwner() string {
	return RandomString(5)
}

// RandomOwner generates random money for the balance
func RandomMoney() int64 {
	return RandomInt(0, 100000)
}

// RandomOwner generates random curryncy
func RandomCurrency() string {
	currencies := []string{"TENGE", "USD", "TAD", "RUB"}
	n := len(currencies) - 1
	return currencies[rand.Intn(n)]
}

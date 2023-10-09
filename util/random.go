package util

import (
	"fmt"
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuwxyz"

// Gerenates a random number (int64) between a given min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// Generates a random string based on a given number
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// Generates a random owner name using RandomString Function
func RandomOwner() string {
	return RandomString(6)
}

// Generates a random monetary value using RandomInt Function
func RandomMoneyValue() int64 {
	return RandomInt(0, 1000)
}

// Generates a random Currency
func RandomCurrency() string {
	currencies := []string{USD, EUR, BRL}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}

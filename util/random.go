package util

import (
	"math/rand"
	"strings"
	"time"
	"unicode"
)

const alphabet = "abcdefghijklmnoprstowzxy"

func init() {
	rand.Seed(time.Now().UnixNano())
}
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomOwner() string {
	// Ensure the string length is at least 1
	if len(alphabet) == 0 {
		return ""
	}

	// First, generate a random character and make it uppercase
	firstChar := rune(alphabet[rand.Intn(len(alphabet))])
	upperFirstChar := unicode.ToUpper(firstChar)

	// Then, generate the rest of the string with the original RandomString function
	restOfString := RandomString(5) // 5 because one character is already generated

	return string(upperFirstChar) + restOfString
}
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}
func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "GBP"}
	return currencies[rand.Intn(len(currencies))]
}

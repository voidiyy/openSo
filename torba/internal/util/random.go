package util

import (
	"math/rand"
	"strings"
	"time"
)

const alpb = "abcdefghijklmnoprstqvwxyz"

func Init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func RandomFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min+1)
}

func RandomString(length int) string {
	var sb strings.Builder

	k := len(alpb)

	for i := 0; i < length; i++ {
		sb.WriteByte(alpb[rand.Intn(k)])
	}
	return sb.String()
}

func RandomIntSlice(size int, min, max int64) []int64 {
	slice := make([]int64, size)
	for i := range slice {
		slice[i] = min + rand.Int63n(max-min+1)
	}
	return slice
}

func RandomStringSlice(size int) []string {
	slice := make([]string, size)
	for i := range slice {
		slice[i] = RandomString(size)
	}
	return slice
}

func UserName() string {
	return RandomString(10)
}

func Email() string {
	return RandomString(8) + "@gmail.com"
}

func Password() string {
	return RandomString(8)
}

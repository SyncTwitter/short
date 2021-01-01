package main

import (
	"math/rand"
	"time"
)

const chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Generation(length int) string {
	b := make([]byte, length)

	currentTime := time.Now().UnixNano()
	rand.Seed(currentTime)

	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}

	return string(b)
}

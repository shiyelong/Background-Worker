package utils

import (
	"math/rand"
	"time"
)

type Captcha struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

func GenerateCaptcha() Captcha {
	rand.Seed(time.Now().UnixNano())
	digits := "0123456789"
	id := generateRandomString(10)
	content := generateRandomString(6, digits)
	return Captcha{
		ID:      id,
		Content: content,
	}
}

func generateRandomString(length int, charset ...string) string {
	var letters string
	if len(charset) > 0 {
		letters = charset[0]
	} else {
		letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	}
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

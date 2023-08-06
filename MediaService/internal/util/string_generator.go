package util

import (
	"fmt"
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func generateRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func GenerateRandomFileName(fileExt string) string {

	uniqueID := generateRandomString(10)
	timeStamp := time.Now().Format("20060102150405")
	uniqueFileName := fmt.Sprintf("%s_%s%s", timeStamp, uniqueID, fileExt)

	return uniqueFileName
}

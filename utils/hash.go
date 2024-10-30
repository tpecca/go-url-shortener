package utils

import "math/rand"

func CreateHash(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]rune, length)
	for i := range result {
		result[i] = rune(charset[rand.Intn(len(charset))])
	}
	return string(result)
}

package helper

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateRandomTransaction() string {
	// Ambil 6 digit terakhir dari timestamp untuk mendapatkan nilai yang lebih pendek
	timestamp := time.Now().Unix() % 1000000
	randomPart := rand.Int63() % 1000000
	return fmt.Sprintf("TXN-%d-%d", timestamp, randomPart)
}

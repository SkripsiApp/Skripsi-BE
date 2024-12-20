package helper

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateRandomTransaction() string {
	timestamp := time.Now().UnixNano()
	randomPart := rand.Int63()
	return fmt.Sprintf("TXN-%d-%d", timestamp, randomPart)
}

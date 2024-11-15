package email

import (
	"crypto/rand"
	"log"
	"math/big"
	"os"
	"strings"
)

func GenerateOTP(length int) (string, error) {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		b[i] = charset[num.Int64()]
	}
	return string(b), nil
}

func ContainsLowerCase(s string) bool {
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			return true
		}
	}
	return false
}

func SendOTPEmail(emailAddress string, otp string) {
	go func() {
		// Buka file template email.
		filePath := "utils/email/templates/otp.html"
		file, err := os.ReadFile(filePath)
		if err != nil {
			log.Printf("gagal membaca template email: %v", err)
			return
		}
		emailTemplate := string(file)

		emailContent := strings.Replace(emailTemplate, "{{.Otp}}", otp, -1)

		_, errEmail := SendEmailSMTPForOTP([]string{emailAddress}, emailContent, otp)
		if errEmail != nil {
			log.Printf("gagal mengirim otp: %v", errEmail)
		}
	}()
}

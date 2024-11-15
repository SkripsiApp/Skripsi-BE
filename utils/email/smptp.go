package email

import (
	"os"
	"strconv"
	"strings"

	"gopkg.in/gomail.v2"
)

func SendEmailSMTPForOTP(to []string, template string, data interface{}) (bool, error) {
	emailHost := os.Getenv("EMAIL_HOST")
	emailFrom := os.Getenv("EMAIL_FROM")
	emailPassword := os.Getenv("EMAIL_PASSWORD")
	emailPortStr := os.Getenv("EMAIL_PORT")
	emailPort, _ := strconv.Atoi(emailPortStr)

	m := gomail.NewMessage()
	m.SetHeader("From", emailFrom)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", "OTP Akun Anda")

	emailContent := strings.Replace(template, "{{.Otp}}", data.(string), -1)

	// Set HTML body
	m.SetBody("text/html", emailContent)

	d := gomail.NewDialer(emailHost, emailPort, emailFrom, emailPassword)

	// Send email
	if err := d.DialAndSend(m); err != nil {
		return false, err
	}
	return true, nil
}

package email

import (
	"log"
	"os"
	"strings"
)

func SendEmailPayment(emailAddress string, customerName string, paymentURL string) {
	go func() {
		// Buka file template email untuk konfirmasi pembayaran.
		filePath := "utils/email/templates/payment-confirmation.html"
		file, err := os.ReadFile(filePath)
		if err != nil {
			log.Printf("gagal membaca template email konfirmasi pembayaran: %v", err)
			return
		}
		emailTemplate := string(file)

		// Ganti placeholder dengan data yang sesuai
		emailContent := strings.Replace(emailTemplate, "{{.CustomerName}}", customerName, -1)
		emailContent = strings.Replace(emailContent, "{{.PaymentURL}}", paymentURL, -1)

		// Kirim email konfirmasi pembayaran
		data := map[string]string{
			"CustomerName": customerName,
			"PaymentURL":   paymentURL,
		}

		_, errEmail := SendEmailSMTPForPayment([]string{emailAddress}, emailContent, data)
		if errEmail != nil {
			log.Printf("gagal mengirim email konfirmasi pembayaran: %v", errEmail)
		}
	}()
}

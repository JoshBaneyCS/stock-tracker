package utils

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendEmail(to string, subject string, body string) error {
	from := os.Getenv("EMAIL_USER")
	pass := os.Getenv("EMAIL_PASS")
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	server := "smtp.gmail.com"
	auth := smtp.PlainAuth("", from, pass, server)

	return smtp.SendMail(server+":587", auth, from, []string{to}, []byte(msg))
}

func FormatStockAlert(symbol string, direction string, price float64) (string, string) {
	subject := fmt.Sprintf("📉 Stock Alert: %s", symbol)
	body := fmt.Sprintf("The stock %s has %s your alert threshold of $%.2f.\n\nCheck your dashboard for more details.", symbol, direction, price)
	return subject, body
}

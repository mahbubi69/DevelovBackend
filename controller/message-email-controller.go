package controller

import (
	"os"
	"strings"
)

func (s *Server) SendOtpToEmail(toEmail []string, senderName, subject, message string) error {

	body := "From: " + senderName + "\n" +
		"To: " + strings.Join(toEmail, ",") + "\n" +
		"Subject: " + subject + "\n\n" +
		message

	err := s.ConfigOtpGmail(
		os.Getenv("CONFIG_AUTH_EMAIL"),
		os.Getenv("CONFIG_AUTH_PASSWORD"),
		os.Getenv("CONFIG_SMTP_HOST"),
		os.Getenv("CONFIG_SMTP_PORT"),
		toEmail,
		[]byte(body))

	if err != nil {
		return err
	}

	return nil
}

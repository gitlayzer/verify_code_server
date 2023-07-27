package utils

import (
	"flag"
	"github.com/gitlayzer/callback_platform/config"
	"net/smtp"
	"os"
	"strings"
)

func SendVerificationCode(email, code, id string) error {
	SmtpUser := config.SmtpUser
	if config.SmtpUser == "" {
		config.SmtpUser = os.Getenv("SMTP_USER")
	}
	SmtpPass := config.SmtpPass
	if config.SmtpPass == "" {
		config.SmtpPass = os.Getenv("SMTP_PASS")
	}
	flag.Parse()
	host := config.SmtpHost + ":" + config.SmtpPort
	if config.SmtpHost == "" && config.SmtpPort == "" {
		config.SmtpHost = os.Getenv("SMTP_HOST")
		config.SmtpPort = os.Getenv("SMTP_PORT")
	}
	to := email

	body := "点击此链接验证：" + "\n" + config.Url + "/verify/" + id + "/" + code

	err := SendToMail(SmtpUser, SmtpPass, host, to, body, "html")
	return err
}

func SendToMail(user, password, host, to, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var contentType string
	if mailtype == "html" {
		contentType = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		contentType = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + ">\r\nSubject: 验证码" + "\r\n" + contentType + "\r\n\r\n" + body)
	sendTo := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, sendTo, msg)
	return err
}

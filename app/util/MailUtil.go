package util

import (
	"errors"
	"fmt"
	"idstar-idp/rest-api/app/config"
	"time"

	"gopkg.in/gomail.v2"
)

var otpExpireFormat = "Jan 02, 2006 03:04:05 PM MST"

type MailUtil struct {
	host       string
	port       int
	username   string
	password   string
	senderName string
	baseUrl    string
}

func NewMailUtil() *MailUtil {
	return &MailUtil{
		host:       config.GetConfigValue("mail.smtp.host"),
		port:       config.GetConfigIntValue("mail.smtp.port"),
		username:   config.GetConfigValue("mail.auth.user"),
		password:   config.GetConfigValue("mail.auth.password"),
		senderName: config.GetConfigValue("mail.sender.name"),
		baseUrl:    config.GetConfigValue("app.base.url"),
	}
}

func (c *MailUtil) sendMail(recipientEmail string, recipientName string, mailSubject string, mailBody string) error {
	mailer := gomail.NewMessage()
	mailer.SetAddressHeader("From", c.username, c.senderName)
	mailer.SetAddressHeader("To", recipientEmail, recipientName)
	mailer.SetHeader("Subject", mailSubject)
	mailer.SetBody("text/html", mailBody)

	dialer := gomail.NewDialer(c.host, c.port, c.username, c.password)
	err := dialer.DialAndSend(mailer)
	if err != nil {
		return errors.New(fmt.Sprint("error send email to ", recipientEmail, ", reason: ", err))
	}
	return nil
}

func (c *MailUtil) SendUserActivationMail(recipientEmail string, recipientName string, otp string, otpExpire time.Time) error {
	subject := "User Activation - IDP Go REST API"
	encodedParam := EncodeForActivationLink(recipientEmail, otp)
	activationLink := c.baseUrl + "/v1/registration/activate?go=" + encodedParam
	body := "Dear " + recipientName + ",<br><br><br>" +
		"You account has been successfully created. To complete registration process, please use this activation link:<br><br>" +
		activationLink + "<br><br>" +
		"Or use the code provided below:<br><br>" +
		"<b>" + otp + "</b><br><br>" +
		"The link and code are valid until <b>" + otpExpire.Format(otpExpireFormat) + "</b>.<br><br><br>" +
		"Thanks,<br>Go REST API team"
	return c.sendMail(recipientEmail, recipientName, subject, body)
}

func (c *MailUtil) SendChangePasswordMail(recipientEmail string, recipientName string, otp string, otpExpire time.Time) error {
	subject := "Reset Password - IDP Go REST API"
	body := "Dear " + recipientName + ",<br><br><br>" +
		"You recently requested to reset the password for your account. To proceed, please use the code provided below:<br><br>" +
		"<b>" + otp + "</b><br><br>" +
		"This code is valid until <b>" + otpExpire.Format(otpExpireFormat) + "</b>. If you did not make this request, please ignore this email.<br><br><br>" +
		"Thanks,<br>Go REST API team"
	return c.sendMail(recipientEmail, recipientName, subject, body)
}

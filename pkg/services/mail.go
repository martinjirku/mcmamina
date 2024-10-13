package services

import "net/smtp"

type MailService struct {
	password string
	account  string
}

func NewMailService(account, password string) MailService {
	return MailService{
		password: password,
		account:  account,
	}
}

func (s *MailService) Send(to, body string) error {
	from := s.account
	pass := s.password
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Hello there\n\n" +
		body
	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))
	if err != nil {
		return err
	}
	return nil
}

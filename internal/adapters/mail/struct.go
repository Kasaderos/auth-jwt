package mail

import (
	"context"
	"net/smtp"
)

type St struct {
	auth     smtp.Auth
	smtpConn string
	email    string
}

func NewClient(smtpConn, email, password string) *St {
	auth := smtp.PlainAuth("", email, password, smtpHost)
	return &St{
		auth:     auth,
		smtpConn: smtpConn,
		email:    email,
	}
}

func (u *St) SendMessage(ctx context.Context, receiverEmail, message string) error {
	// Receiver email address.
	to := []string{
		receiverEmail,
	}

	// TODO send with context
	err := smtp.SendMail(
		u.smtpConn,
		u.auth,
		u.email,
		to,
		message,
	)
	if err != nil {
		return err
	}

	return nil
}

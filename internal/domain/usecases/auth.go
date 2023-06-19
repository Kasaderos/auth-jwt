package usecases

import (
	"context"
)

// todo
var hmacSampleSecret = []byte("vicrotia-secret")

func (u *St) Auth(ctx context.Context, tokenString string) error {
}

func (u *St) SendMessage(ctx context.Context, email, message string) error {
	return u.cr.Mail.SendMessage(ctx, email, message)
}

package usecases

import (
	"birthday-bot/internal/domain/entities"
	"context"
)

// todo
var hmacSampleSecret = []byte("vicrotia-secret")

func (u *St) GetUserByEmail(ctx context.Context, email string) (*entities.UserSt, error) {
	return u.cr.User.GetByEmail(ctx, email, true)
}

func (u *St) SendActivationMessage(ctx context.Context, email, message string) error {
	return u.cr.Mail.SendMessage(ctx, email, message)
}

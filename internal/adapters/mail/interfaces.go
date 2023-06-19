package mail

import "context"

type Client interface {
	SendMessage(ctx context.Context, email string, message string) error
}

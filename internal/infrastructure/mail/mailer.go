package mail

import "context"

type Mailer interface {
	Send(ctx context.Context, message Message) error
}

type Message struct {
	From    string
	To      string
	Subject string
	Body    string
}

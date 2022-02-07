package email_provider

type EmailProvider interface {
	Send() error
}

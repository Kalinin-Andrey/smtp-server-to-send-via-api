package unisender

import (
	"smtp2api/pkg/email_provider"
)

type EmailProvider struct {
}

func New() email_provider.EmailProvider {
	return nil
}

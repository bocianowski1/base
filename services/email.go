package services

import "log/slog"

type IEmailService interface {
	Send(addrs []string, subject, body string) error
}

type EmailService struct {
	// logger repo.ILogger
	// cache  repo.ICache
}

func NewEmailService() IEmailService {
	return &EmailService{}
}

func (s *EmailService) Send(to []string, subject, body string) error {
	slog.Warn("not implemented")
	return nil
}

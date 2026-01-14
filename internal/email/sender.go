package email

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/go-mail/mail"
)

// EmailService defines the interface for sending emails
// This allows us to swap implementations (SMTP, Mailgun) without changing handler code
type EmailService interface {
	Send(to, subject, body string) error
}

// SMTPSender implements EmailService using SMTP
type SMTPSender struct {
	username string
	password string
	host     string
	port     int
}

// NewSMTPSender creates a new SMTP email service from environment variables
func NewSMTPSender() *SMTPSender {
	return &SMTPSender{
		username: os.Getenv("SMTP_USERNAME"),
		password: os.Getenv("SMTP_PASSWORD"),
		host:     os.Getenv("SMTP_HOST"),
		port:     587, // Standard SMTP port
	}
}

// Send delivers an email via SMTP
func (s *SMTPSender) Send(to, subject, body string) error {
	email := mail.NewMessage()
	email.SetHeader("To", to)
	email.SetHeader("From", s.username)
	email.SetHeader("Subject", subject)
	email.SetBody("text/plain", body)

	dialer := mail.NewDialer(s.host, s.port, s.username, s.password)
	return dialer.DialAndSend(email)
}

// MailgunSender implements EmailService using Mailgun API
type MailgunSender struct {
	domain string
	apiKey string
	from   string
}

// NewMailgunSender creates a new Mailgun email service
func NewMailgunSender() *MailgunSender {
	return &MailgunSender{
		domain: os.Getenv("MAILGUN_DOMAIN"),
		apiKey: os.Getenv("MAILGUN_API_KEY"),
		from:   os.Getenv("MAILGUN_FROM_EMAIL"),
	}
}

// Send delivers an email via Mailgun API
func (s *MailgunSender) Send(to, subject, body string) error {
	if s.domain == "" || s.apiKey == "" || s.from == "" {
		return fmt.Errorf("Mailgun configuration missing: domain=%s, apiKey set=%v, from=%s",
			s.domain, s.apiKey != "", s.from)
	}

	// Mailgun API endpoint
	apiURL := fmt.Sprintf("https://api.mailgun.net/v3/%s/messages", s.domain)

	// Prepare form data
	data := url.Values{}
	data.Set("from", s.from)
	data.Set("to", to)
	data.Set("subject", subject)
	data.Set("text", body)

	// Create request with Basic Auth
	req, _ := http.NewRequest("POST", apiURL, strings.NewReader(data.Encode()))
	auth := base64.StdEncoding.EncodeToString([]byte("api:" + s.apiKey))
	req.Header.Add("Authorization", "Basic "+auth)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send email via Mailgun: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("Mailgun returned status %d", resp.StatusCode)
	}

	return nil
}

// NewEmailService is a factory function to get the appropriate email service
// This allows for a provider switch by changing an env variable
func NewEmailService() (EmailService, error) {
	provider := os.Getenv("EMAIL_PROVIDER")
	if provider == "" {
		provider = "mailgun" // Default to Mailgun (free tier, future-proof)
	}

	switch provider {
	case "smtp":
		return NewSMTPSender(), nil
	case "mailgun":
		return NewMailgunSender(), nil
	default:
		return nil, fmt.Errorf("unknown email provider: %s", provider)
	}
}

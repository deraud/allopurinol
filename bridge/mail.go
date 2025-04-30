package bridge

import (
	"context"
	"errors"
	"fmt"
	"github.com/resend/resend-go/v2"
	"healthcare-allopurinol/entity"
	"healthcare-allopurinol/utility/mail"
	"os"
)

type MailBridgeItf interface {
	SendVerifyEmail(ctx context.Context, recipient, name, token string) error
	SendResetEmail(ctx context.Context, recipient, token string) error
	SendPharmacistEmail(ctx context.Context, recipient, password string) error
}

type MailBridgeImpl struct {
	mail *resend.Client
}

func NewMailBridge(mail *resend.Client) MailBridgeItf {
	return &MailBridgeImpl{mail: mail}
}

func sendEmail(r *MailBridgeImpl, ctx context.Context, content *entity.Mail) error {
	sender := fmt.Sprintf("%s <no-reply@%s>", os.Getenv("MAIL_SENDER"), os.Getenv("MAIL_DOMAIN"))

	email := &resend.SendEmailRequest{
		To:      []string{content.Recipient},
		From:    sender,
		Subject: content.Subject,
		Html:    content.Body,
	}

	_, err := r.mail.Emails.SendWithContext(ctx, email)
	if err != nil {
		return errors.New(fmt.Sprintf("Resend: %v", err))
	}

	return nil
}

func (r *MailBridgeImpl) SendVerifyEmail(ctx context.Context, recipient, name, token string) error {
	const message = "Please click the button below to verify your email."
	var link = fmt.Sprintf("https://digitalent.games.test.shopee.io/vm1/auth/verify?token=%s", token)

	content := &entity.Mail{
		Recipient: recipient,
		Subject:   "Verify Email",
		Body:      fmt.Sprintf(mail.AuthHtml, name, message, link, "Verify Now"),
	}

	return sendEmail(r, ctx, content)
}

func (r *MailBridgeImpl) SendResetEmail(ctx context.Context, recipient, token string) error {
	const message = "Please click the button below to reset your password."
	var link = fmt.Sprintf("https://digitalent.games.test.shopee.io/vm1/auth/reset-password?token=%s", token)

	content := &entity.Mail{
		Recipient: recipient,
		Subject:   "Reset Password",
		Body:      fmt.Sprintf(mail.AuthHtml, "Dear User", message, link, "Reset Now"),
	}

	return sendEmail(r, ctx, content)
}

func (r *MailBridgeImpl) SendPharmacistEmail(ctx context.Context, recipient, password string) error {
	const message = "Here are your account credentials:"

	content := &entity.Mail{
		Recipient: recipient,
		Subject:   "Pharmacist Credentials",
		Body:      fmt.Sprintf(mail.PharmacistHtml, message, recipient, password),
	}

	return sendEmail(r, ctx, content)
}

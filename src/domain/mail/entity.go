package mail

import (
	"time"

	"github.com/google/uuid"
)

type Mail struct {
	ID        uuid.UUID
	Subject   string
	Body      string
	SenderID  uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

type MailRecipient struct {
	ID             uuid.UUID
	MailID         uuid.UUID
	RecipientEmail string
	Type           RecipientType
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func NewMail(id uuid.UUID, subject, body string, senderID uuid.UUID, createdAt, updatedAt time.Time) *Mail {
	return &Mail{
		ID:        id,
		Subject:   subject,
		Body:      body,
		SenderID:  senderID,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

func NewMailRecipient(id uuid.UUID, mailID uuid.UUID, recipientEmail string, recipientType RecipientType, createdAt, updatedAt time.Time) *MailRecipient {
	return &MailRecipient{
		ID:             id,
		MailID:         mailID,
		RecipientEmail: recipientEmail,
		Type:           recipientType,
		CreatedAt:      createdAt,
		UpdatedAt:      updatedAt,
	}
}

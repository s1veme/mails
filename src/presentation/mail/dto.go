package mail

import (
	"time"

	"mails/src/presentation/common"

	"github.com/google/uuid"
)

type MailRequest struct {
	Subject       string   `json:"subject"`
	Body          string   `json:"body"`
	RecipientsTo  []string `json:"recipients_to"`
	RecipientsCC  []string `json:"recipients_cc"`
	RecipientsBCC []string `json:"recipients_bcc"`
}

type EmailListResponse struct {
	Items      []MailResponse    `json:"items"`
	Pagination common.Pagination `json:"pagination"`
}

type MailResponse struct {
	ID         uuid.UUID           `json:"id"`
	Subject    string              `json:"subject"`
	Body       string              `json:"body"`
	SenderID   uuid.UUID           `json:"sender_id"`
	CreatedAt  time.Time           `json:"created_at"`
	UpdatedAt  time.Time           `json:"updated_at"`
	Recipients []RecipientResponse `json:"recipients"`
}

type RecipientResponse struct {
	Email string `json:"email"`
	Type  string `json:"type"`
}

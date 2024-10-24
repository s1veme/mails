package mail

type RecipientType string

const (
	ToRecipient  RecipientType = "TO"
	CcRecipient  RecipientType = "CC"
	BccRecipient RecipientType = "BCC"
)

package folder

import (
	"time"

	"github.com/google/uuid"
)

type FolderRequest struct {
	Name string `json:"name"`
}

type FolderResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type MoveMailRequest struct {
	MailID   uuid.UUID `json:"mail_id"`
	FolderID uuid.UUID `json:"folder_id"`
}

type RemoveMailFromFolderRequest struct {
	MailID   uuid.UUID `json:"mail_id"`
	FolderID uuid.UUID `json:"folder_id"`
}

type DeleteFolderRequest struct {
	FolderID uuid.UUID `json:"folder_id"`
}

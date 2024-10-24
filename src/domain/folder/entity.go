package folder

import (
	"time"

	"github.com/google/uuid"
)

type Folder struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type MailFolder struct {
	ID       uuid.UUID
	MailID   uuid.UUID
	FolderID uuid.UUID
}

func NewFolder(id, userID uuid.UUID, name string, createdAt, updatedAt time.Time) *Folder {
	return &Folder{
		ID:        id,
		UserID:    userID,
		Name:      name,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

func NewMailFolder(id, mailID uuid.UUID, folderID uuid.UUID) *MailFolder {
	return &MailFolder{
		ID:       id,
		MailID:   mailID,
		FolderID: folderID,
	}
}

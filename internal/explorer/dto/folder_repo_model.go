package dto

//FolderDTO for Folder entity in the database.

import (
	"context"
	"database/sql"
	"remy_explorer/internal/explorer/domain"
	"time"
)

type FolderRepository interface {
	CreateFolder(ctx context.Context, folder *FolderDTO) (*int64, error)
	GetFolderByID(ctx context.Context, id int64) (*FolderDTO, error)
	GetFoldersByParentID(ctx context.Context, parentID int64) ([]*FolderDTO, error)
	UpdateFolder(ctx context.Context, folder *FolderDTO) error
	DeleteFolder(ctx context.Context, id int64) error
}

// FolderDTO is the data transfer object for the Folder entity in the database.
type FolderDTO struct {
	ID        int64         `json:"id"`
	OwnerID   string        `json:"owner_id"`
	Name      string        `json:"name"`
	ParentID  sql.NullInt64 `json:"parent_id"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

func (m *FolderDTO) ToDomain() *domain.Folder {
	return &domain.Folder{
		ID:        m.ID,
		OwnerID:   m.OwnerID,
		Name:      m.Name,
		ParentID:  m.ParentID.Int64,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func FolderToDTO(f *domain.Folder) *FolderDTO {
	return &FolderDTO{
		ID:        f.ID,
		OwnerID:   f.OwnerID,
		Name:      f.Name,
		ParentID:  sql.NullInt64{Int64: f.ParentID, Valid: true},
		CreatedAt: f.CreatedAt,
		UpdatedAt: f.UpdatedAt,
	}
}

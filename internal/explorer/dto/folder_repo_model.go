package dto

//FolderDTO for Folder entity in the database.

import (
	"context"
	"database/sql"
	"remy_explorer/internal/explorer/model"
	"strconv"
	"time"
)

type FolderRepository interface {
	CreateFolder(ctx context.Context, folder *FolderDTO) (*string, error)
	GetFolderByID(ctx context.Context, id string) (*FolderDTO, error)
	GetFoldersByParentID(ctx context.Context, parentID string) ([]*FolderDTO, error)
	UpdateFolder(ctx context.Context, folder *FolderDTO) error
	DeleteFolder(ctx context.Context, id string) error
}

// FolderDTO is the data transfer object for the Folder entity in the database.
type FolderDTO struct {
	ID        int            `json:"id"`
	OwnerID   string         `json:"owner_id"`
	Name      string         `json:"name"`
	ParentID  sql.NullString `json:"parent_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

func (m *FolderDTO) ToDomain() *model.Folder {
	return &model.Folder{
		ID:        strconv.Itoa(m.ID),
		OwnerID:   m.OwnerID,
		Name:      m.Name,
		ParentID:  m.ParentID.String,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func FolderToDTO(f *model.Folder) *FolderDTO {
	id, _ := strconv.Atoi(f.ID) //TODO: rewrite
	return &FolderDTO{
		ID:        id,
		OwnerID:   f.OwnerID,
		Name:      f.Name,
		ParentID:  sql.NullString{String: f.ParentID, Valid: true},
		CreatedAt: f.CreatedAt,
		UpdatedAt: f.UpdatedAt,
	}
}

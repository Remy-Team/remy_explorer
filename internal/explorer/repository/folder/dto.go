package folder

//DTO for Folder entity in the database.

import (
	"database/sql"
	"remy_explorer/pkg/domain/folder"
	"remy_explorer/pkg/domain/user"
	"time"
)

// FolderDTOID is the type for the ID of the FolderDTO.
type FolderDTOID int64

// FolderDTO is the data transfer object for the Folder entity in the database.
type FolderDTO struct {
	ID        FolderDTOID   `json:"id"`
	OwnerID   user.ID       `json:"owner_id"`
	Name      string        `json:"name"`
	ParentID  sql.NullInt64 `json:"parent_id"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

func (m *FolderDTO) ToDomain() *folder.Folder {
	return &folder.Folder{
		ID:        folder.FolderID(m.ID),
		OwnerID:   m.OwnerID,
		Name:      m.Name,
		ParentID:  folder.FolderID(m.ParentID.Int64),
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func ToDTO(f *folder.Folder) *FolderDTO {
	return &FolderDTO{
		ID:        FolderDTOID(f.ID),
		OwnerID:   f.OwnerID,
		Name:      f.Name,
		ParentID:  sql.NullInt64{Int64: int64(f.ParentID), Valid: true},
		CreatedAt: f.CreatedAt,
		UpdatedAt: f.UpdatedAt,
	}
}

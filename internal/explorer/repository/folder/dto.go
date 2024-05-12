package folder

//DTO for Folder entity in the database.

import (
	"database/sql"
	"remy_explorer/internal/explorer/service/folder"
	"time"
)

// DTO is the data transfer object for the Folder entity in the database.
type DTO struct {
	ID        int64         `json:"id"`
	OwnerID   string        `json:"owner_id"`
	Name      string        `json:"name"`
	ParentID  sql.NullInt64 `json:"parent_id"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

func (m *DTO) ToDomain() *folder.Folder {
	return &folder.Folder{
		ID:        m.ID,
		OwnerID:   m.OwnerID,
		Name:      m.Name,
		ParentID:  m.ParentID.Int64,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func ToDTO(f *folder.Folder) *DTO {
	return &DTO{
		ID:        f.ID,
		OwnerID:   f.OwnerID,
		Name:      f.Name,
		ParentID:  sql.NullInt64{Int64: f.ParentID, Valid: true},
		CreatedAt: f.CreatedAt,
		UpdatedAt: f.UpdatedAt,
	}
}

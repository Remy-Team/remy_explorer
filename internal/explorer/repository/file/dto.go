package file

// DTO for File entity in the database.
import (
	"database/sql"
	"remy_explorer/internal/explorer/service/file"
	"time"
)

// DTO is the data transfer object for the File entity in the database.
type DTO struct {
	ID         int64            `json:"id"`
	OwnerID    string           `json:"owner_id"`
	Name       string           `json:"name"`
	FolderID   int64            `json:"folder_id"`
	ObjectPath sql.NullString   `json:"object_path"`
	Size       int              `json:"size"`
	Type       sql.NullString   `json:"type"`
	CreatedAt  time.Time        `json:"created_at"`
	UpdatedAt  time.Time        `json:"updated_at"`
	Tags       []sql.NullString `json:"tags"`
}

func (d DTO) ToDomain() *file.File {
	tags := make([]string, 0)
	for _, tag := range d.Tags {
		tags = append(tags, tag.String)
	}
	return &file.File{
		ID:         d.ID,
		OwnerID:    d.OwnerID,
		Name:       d.Name,
		FolderID:   d.FolderID,
		ObjectPath: d.ObjectPath.String,
		Size:       d.Size,
		Type:       d.Type.String,
		CreatedAt:  d.CreatedAt,
		UpdatedAt:  d.UpdatedAt,
		Tags:       tags,
	}
}

// ToDTO converts a File to a FileDTO.
func ToDTO(f *file.File) DTO {
	tags := make([]sql.NullString, 0)
	for _, tag := range f.Tags {
		tags = append(tags, sql.NullString{String: tag, Valid: true})
	}
	return DTO{
		ID:         f.ID,
		OwnerID:    f.OwnerID,
		Name:       f.Name,
		FolderID:   f.FolderID,
		ObjectPath: sql.NullString{String: f.ObjectPath, Valid: true},
		Size:       f.Size,
		Type:       sql.NullString{String: f.Type, Valid: true},
		CreatedAt:  f.CreatedAt,
		UpdatedAt:  f.UpdatedAt,
		Tags:       tags,
	}
}

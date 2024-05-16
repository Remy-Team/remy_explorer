package dto

// FileDTO for File entity in the database.
import (
	"context"
	"database/sql"
	"remy_explorer/internal/explorer/domain"
	"time"
)

// FileRepository is the interface that defines the methods that a file repository must implement.
type FileRepository interface {
	CreateFile(ctx context.Context, file *FileDTO) (*int64, error)
	GetFileByID(ctx context.Context, id int64) (*FileDTO, error)
	GetFilesByFolderID(ctx context.Context, folderID int64) ([]*FileDTO, error)
	UpdateFile(ctx context.Context, file *FileDTO) error
	DeleteFile(ctx context.Context, id int64) error
	GetFilesByFolderIdSorted(ctx context.Context, folderID int64, sortOption *SortOption) ([]*FileDTO, error)
}

//TODO Можно добавить лимит и оффсет для пагинации файлов

// FileDTO is the data transfer object for the File entity in the database.
type FileDTO struct {
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

func (d FileDTO) ToDomain() *domain.File {
	tags := make([]string, 0)
	for _, tag := range d.Tags {
		tags = append(tags, tag.String)
	}
	return &domain.File{
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

// FileToDTO converts a File to a FileDTO.
func FileToDTO(f *domain.File) FileDTO {
	tags := make([]sql.NullString, 0)
	for _, tag := range f.Tags {
		tags = append(tags, sql.NullString{String: tag, Valid: true})
	}
	return FileDTO{
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

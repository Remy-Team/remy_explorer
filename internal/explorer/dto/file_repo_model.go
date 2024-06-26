package dto

// FileDTO for File entity in the database.
import (
	"context"
	"database/sql"
	"remy_explorer/internal/explorer/model"
	"strconv"
	"time"
)

// FileRepository is the interface that defines the methods that a file repository must implement.
type FileRepository interface {
	CreateFile(ctx context.Context, file *FileDTO) (*string, error)
	GetFileByID(ctx context.Context, id string) (*FileDTO, error)
	GetFilesByFolderID(ctx context.Context, folderID string) ([]*FileDTO, error)
	UpdateFile(ctx context.Context, file *FileDTO) error
	DeleteFile(ctx context.Context, id string) error
	GetFilesByFolderIdSorted(ctx context.Context, folderID string, sortOption *SortOption) ([]*FileDTO, error)
}

//TODO Можно добавить лимит и оффсет для пагинации файлов

// FileDTO is the data transfer object for the File entity in the database.
type FileDTO struct {
	ID         int              `json:"id"`
	OwnerID    string           `json:"owner_id"`
	Name       string           `json:"name"`
	FolderID   int              `json:"folder_id"`
	ObjectPath sql.NullString   `json:"object_path"`
	Size       int              `json:"size"`
	Type       sql.NullString   `json:"type"`
	CreatedAt  time.Time        `json:"created_at"`
	UpdatedAt  time.Time        `json:"updated_at"`
	Tags       []sql.NullString `json:"tags"`
}

func (d FileDTO) ToDomain() *model.File {
	tags := make([]string, 0)
	for _, tag := range d.Tags {
		tags = append(tags, tag.String)
	}
	return &model.File{
		ID:         strconv.Itoa(d.ID),
		OwnerID:    d.OwnerID,
		Name:       d.Name,
		FolderID:   strconv.Itoa(d.FolderID),
		ObjectPath: d.ObjectPath.String,
		Size:       d.Size,
		Type:       d.Type.String,
		CreatedAt:  d.CreatedAt,
		UpdatedAt:  d.UpdatedAt,
		Tags:       tags,
	}
}

// FileToDTO converts a File to a FileDTO.
func FileToDTO(f *model.File) FileDTO {
	tags := make([]sql.NullString, 0)
	for _, tag := range f.Tags {
		tags = append(tags, sql.NullString{String: tag, Valid: true})
	}
	id, _ := strconv.Atoi(f.ID)         //TODO: rewrite
	f_id, _ := strconv.Atoi(f.FolderID) //TODO: rewrite
	return FileDTO{
		ID:         id,
		OwnerID:    f.OwnerID,
		Name:       f.Name,
		FolderID:   f_id,
		ObjectPath: sql.NullString{String: f.ObjectPath, Valid: true},
		Size:       f.Size,
		Type:       sql.NullString{String: f.Type, Valid: true},
		CreatedAt:  f.CreatedAt,
		UpdatedAt:  f.UpdatedAt,
		Tags:       tags,
	}
}

package file

import (
	"context"
)

// Repository is the interface that defines the methods that a file repository must implement.
type Repository interface {
	CreateFile(ctx context.Context, file *DTO) (*int64, error)
	GetFileByID(ctx context.Context, id int64) (*DTO, error)
	GetFilesByFolderID(ctx context.Context, folderID int64) ([]*DTO, error)
	UpdateFile(ctx context.Context, file *DTO) error
	DeleteFile(ctx context.Context, id int64) error
	GetFilesByFolderIdSorted(ctx context.Context, folderID int64, sortOption *SortOption) ([]*DTO, error)
}

//TODO Можно доабавить лимит и оффсет для пагинации файлов

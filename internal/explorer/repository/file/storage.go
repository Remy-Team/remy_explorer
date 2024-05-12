package file

import (
	"context"
)

// Repository is the interface that defines the methods that a file repository must implement.
type Repository interface {
	CreateFile(ctx context.Context, file *DTO) (*int64, error)
	GetFileByID(ctx context.Context, id FileDTOID) (*DTO, error)
	GetFilesByFolderID(ctx context.Context, folderID FileDTOID) ([]*DTO, error)
	UpdateFile(ctx context.Context, file *DTO) error
	DeleteFile(ctx context.Context, id FileDTOID) error
	GetFilesByFolderIdSorted(ctx context.Context, folderID FileDTOID, sortOption *SortOption) ([]*DTO, error)
}

//TODO Можно доабавить лимит и оффсет для пагинации файлов

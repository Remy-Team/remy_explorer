package file

import (
	"context"
)

// Repository is the interface that defines the methods that a file repository must implement.
type Repository interface {
	CreateFile(ctx context.Context, file *FileDTO) error
	GetFileByID(ctx context.Context, id FileDTOID) (*FileDTO, error)
	GetFilesByFolderID(ctx context.Context, folderID FileDTOID) ([]*FileDTO, error)
	UpdateFile(ctx context.Context, file *FileDTO) error
	DeleteFile(ctx context.Context, id FileDTOID) error
	GetFilesByFolderIdSorted(ctx context.Context, folderID FileDTOID, sortOption *SortOption) ([]*FileDTO, error)
}

//TODO Можно доабавить лимит и оффсет для пагинации файлов

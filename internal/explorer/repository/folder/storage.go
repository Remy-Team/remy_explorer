package folder

import (
	"context"
)

type Repository interface {
	CreateFolder(ctx context.Context, folder *FolderDTO) error
	GetFolderByID(ctx context.Context, id FolderDTOID) (*FolderDTO, error)
	GetFoldersByParentID(ctx context.Context, parentID FolderDTOID) ([]*FolderDTO, error)
	UpdateFolder(ctx context.Context, folder *FolderDTO) error
	DeleteFolder(ctx context.Context, id FolderDTOID) error
}

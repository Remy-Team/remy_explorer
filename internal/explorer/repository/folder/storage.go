package folder

import (
	"context"
)

type Repository interface {
	CreateFolder(ctx context.Context, folder *DTO) (*int64, error)
	GetFolderByID(ctx context.Context, id int64) (*DTO, error)
	GetFoldersByParentID(ctx context.Context, parentID int64) ([]*DTO, error)
	UpdateFolder(ctx context.Context, folder *DTO) error
	DeleteFolder(ctx context.Context, id int64) error
}

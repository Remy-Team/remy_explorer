package folder

import (
	"context"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"remy_explorer/internal/explorer/domain"
	"remy_explorer/internal/explorer/dto"
)

type FolderService interface {
	CreateFolder(ctx context.Context, folder *domain.Folder) (*int64, error)
	GetFolderByID(ctx context.Context, id int64) (*domain.Folder, error)
	GetFoldersByParentID(ctx context.Context, parentID int64) ([]*domain.Folder, error)
	UpdateFolder(ctx context.Context, folder *domain.Folder) error
	DeleteFolder(ctx context.Context, id int64) error
}

type service struct {
	repo dto.FolderRepository
	log  log.Logger
}

func (s service) CreateFolder(ctx context.Context, f *domain.Folder) (*int64, error) {
	logger := log.With(s.log, "folder", "CreateFolder")
	folderDTO := dto.FolderToDTO(f)
	id, err := s.repo.CreateFolder(ctx, folderDTO)
	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, err
	}
	logger.Log("message", "Folder created", "id", folderDTO.ID)
	return id, nil
}

func (s service) GetFolderByID(ctx context.Context, id int64) (*domain.Folder, error) {
	logger := log.With(s.log, "folder", "GetFolderByID")
	folderDTO, err := s.repo.GetFolderByID(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, err

	}
	logger.Log("message", "Folder retrieved", "id", folderDTO.ID)
	return folderDTO.ToDomain(), nil
}

func (s service) GetFoldersByParentID(ctx context.Context, parentID int64) ([]*domain.Folder, error) {
	logger := log.With(s.log, "folder", "GetFoldersByParentID")
	folderDTOs, err := s.repo.GetFoldersByParentID(ctx, parentID)
	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, err
	}
	folders := make([]*domain.Folder, len(folderDTOs))
	for i, f := range folderDTOs {
		folders[i] = f.ToDomain()
	}
	logger.Log("message", "Folders retrieved", "count", len(folders))
	return folders, nil
}

func (s service) UpdateFolder(ctx context.Context, folder *domain.Folder) error {
	logger := log.With(s.log, "folder", "UpdateFolder")
	folderDTO := dto.FolderToDTO(folder)
	if err := s.repo.UpdateFolder(ctx, folderDTO); err != nil {
		level.Error(logger).Log("err", err)
		return err
	}
	logger.Log("message", "Folder updated", "id", folderDTO.ID)
	return nil
}

func (s service) DeleteFolder(ctx context.Context, id int64) error {
	logger := log.With(s.log, "folder", "DeleteFolder")
	if err := s.repo.DeleteFolder(ctx, id); err != nil {
		level.Error(logger).Log("err", err)
		return err
	}
	logger.Log("message", "Folder deleted", "id", id)
	return nil
}

func NewService(repo dto.FolderRepository, logger log.Logger) FolderService {
	return &service{
		repo: repo,
		log:  log.With(logger, "service", "folder"),
	}
}

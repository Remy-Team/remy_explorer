package folder

import (
	"context"
	"errors"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"remy_explorer/internal/explorer/dto"
	modelerr "remy_explorer/internal/explorer/err"
	"remy_explorer/internal/explorer/model"
)

type FolderService interface {
	CreateFolder(ctx context.Context, folder *model.Folder) (*string, error)
	GetFolderByID(ctx context.Context, id string) (*model.Folder, error)
	GetFoldersByParentID(ctx context.Context, parentID string) ([]*model.Folder, error)
	UpdateFolder(ctx context.Context, folder *model.Folder) error
	DeleteFolder(ctx context.Context, id string) error
}

type service struct {
	repo dto.FolderRepository
	log  log.Logger
}

func (s service) CreateFolder(ctx context.Context, f *model.Folder) (*string, error) {
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

func (s service) GetFolderByID(ctx context.Context, id string) (*model.Folder, error) {
	logger := log.With(s.log, "folder", "GetFolderByID")
	folderDTO, err := s.repo.GetFolderByID(ctx, id)
	if err != nil {
		var errNotFound *modelerr.NotFound
		if errors.As(err, &errNotFound) {
			level.Info(logger).Log("err", err, "msg", "file not found")
			return nil, err
		}
		level.Error(logger).Log("err", err)
		return nil, err

	}
	logger.Log("message", "Folder retrieved", "id", folderDTO.ID)
	return folderDTO.ToDomain(), nil
}

func (s service) GetFoldersByParentID(ctx context.Context, parentID string) ([]*model.Folder, error) {
	logger := log.With(s.log, "folder", "GetFoldersByParentID")
	folderDTOs, err := s.repo.GetFoldersByParentID(ctx, parentID)
	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, err
	}
	folders := make([]*model.Folder, len(folderDTOs))
	for i, f := range folderDTOs {
		folders[i] = f.ToDomain()
	}
	logger.Log("message", "Folders retrieved", "count", len(folders))
	return folders, nil
}

func (s service) UpdateFolder(ctx context.Context, folder *model.Folder) error {
	logger := log.With(s.log, "folder", "UpdateFolder")
	folderDTO := dto.FolderToDTO(folder)
	if err := s.repo.UpdateFolder(ctx, folderDTO); err != nil {
		level.Error(logger).Log("err", err)
		return err
	}
	logger.Log("message", "Folder updated", "id", folderDTO.ID)
	return nil
}

func (s service) DeleteFolder(ctx context.Context, id string) error {
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

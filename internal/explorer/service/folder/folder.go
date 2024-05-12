package folder

import (
	"context"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	folderrep "remy_explorer/internal/explorer/repository/folder"
)

type Service interface {
	CreateFolder(ctx context.Context, folder *Folder) (*int64, error)
	GetFolderByID(ctx context.Context, id int64) (*Folder, error)
	GetFoldersByParentID(ctx context.Context, parentID int64) ([]*Folder, error)
	UpdateFolder(ctx context.Context, folder *Folder) error
	DeleteFolder(ctx context.Context, id int64) error
}

type service struct {
	repo folderrep.Repository
	log  log.Logger
}

func (s service) CreateFolder(ctx context.Context, f *Folder) (*int64, error) {
	logger := log.With(s.log, "folder", "CreateFolder")
	folderDTO := folderrep.ToDTO(f)
	id, err := s.repo.CreateFolder(ctx, folderDTO)
	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, err
	}
	logger.Log("message", "Folder created", "id", folderDTO.ID)
	return id, nil
}

func (s service) GetFolderByID(ctx context.Context, id int64) (*Folder, error) {
	logger := log.With(s.log, "folder", "GetFolderByID")
	folderDTO, err := s.repo.GetFolderByID(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, err

	}
	logger.Log("message", "Folder retrieved", "id", folderDTO.ID)
	return folderDTO.ToDomain(), nil
}

func (s service) GetFoldersByParentID(ctx context.Context, parentID int64) ([]*Folder, error) {
	logger := log.With(s.log, "folder", "GetFoldersByParentID")
	folderDTOs, err := s.repo.GetFoldersByParentID(ctx, parentID)
	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, err
	}
	folders := make([]*Folder, len(folderDTOs))
	for i, f := range folderDTOs {
		folders[i] = f.ToDomain()
	}
	logger.Log("message", "Folders retrieved", "count", len(folders))
	return folders, nil
}

func (s service) UpdateFolder(ctx context.Context, folder *Folder) error {
	logger := log.With(s.log, "folder", "UpdateFolder")
	folderDTO := folderrep.ToDTO(folder)
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

func NewService(repo folderrep.Repository, logger log.Logger) Service {
	return &service{
		repo: repo,
		log:  log.With(logger, "service", "folder"),
	}
}

package service

import (
	"context"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	folder2 "remy_explorer/internal/explorer/repository/folder"
	"remy_explorer/pkg/domain/folder"
)

type Service interface {
	CreateFolder(ctx context.Context, folder *folder.Folder) error
	GetFolderByID(ctx context.Context, id folder.FolderID) (*folder.Folder, error)
	GetFoldersByParentID(ctx context.Context, parentID folder.FolderID) ([]*folder.Folder, error)
	UpdateFolder(ctx context.Context, folder *folder.Folder) error
	DeleteFolder(ctx context.Context, id folder.FolderID) error
}

type service struct {
	repo folder2.Repository
	log  log.Logger
}

func (s service) CreateFolder(ctx context.Context, folder *folder.Folder) error {
	logger := log.With(s.log, "folder", "CreateFolder")
	folderDTO := folder2.ToDTO(folder)

	if err := s.repo.CreateFolder(ctx, folderDTO); err != nil {
		level.Error(logger).Log("err", err)
		return err
	}
	logger.Log("message", "Folder created", "id", folderDTO.ID)
	return nil
}

func (s service) GetFolderByID(ctx context.Context, id folder.FolderID) (*folder.Folder, error) {
	logger := log.With(s.log, "folder", "GetFolderByID")
	folderDTO, err := s.repo.GetFolderByID(ctx, folder2.FolderDTOID(id))
	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, err

	}
	logger.Log("message", "Folder retrieved", "id", folderDTO.ID)
	return folderDTO.ToDomain(), nil
}

func (s service) GetFoldersByParentID(ctx context.Context, parentID folder.FolderID) ([]*folder.Folder, error) {
	logger := log.With(s.log, "folder", "GetFoldersByParentID")
	folderDTOs, err := s.repo.GetFoldersByParentID(ctx, folder2.FolderDTOID(parentID))
	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, err
	}
	folders := make([]*folder.Folder, len(folderDTOs))
	for i, f := range folderDTOs {
		folders[i] = f.ToDomain()
	}
	logger.Log("message", "Folders retrieved", "count", len(folders))
	return folders, nil
}

func (s service) UpdateFolder(ctx context.Context, folder *folder.Folder) error {
	logger := log.With(s.log, "folder", "UpdateFolder")
	folderDTO := folder2.ToDTO(folder)
	if err := s.repo.UpdateFolder(ctx, folderDTO); err != nil {
		level.Error(logger).Log("err", err)
		return err
	}
	logger.Log("message", "Folder updated", "id", folderDTO.ID)
	return nil
}

func (s service) DeleteFolder(ctx context.Context, id folder.FolderID) error {
	logger := log.With(s.log, "folder", "DeleteFolder")
	if err := s.repo.DeleteFolder(ctx, folder2.FolderDTOID(id)); err != nil {
		level.Error(logger).Log("err", err)
		return err
	}
	logger.Log("message", "Folder deleted", "id", id)
	return nil
}

func NewService(repo folder2.Repository, logger log.Logger) Service {
	return &service{
		repo: repo,
		log:  log.With(logger, "service", "folder"),
	}
}

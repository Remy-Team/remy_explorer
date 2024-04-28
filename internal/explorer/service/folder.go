package service

import (
	folder2 "remy_explorer/internal/explorer/repository/folder"
	"remy_explorer/pkg/domain/folder"
	"remy_explorer/pkg/logging"

	"context"
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
	log  *logging.Logger
}

func (s service) CreateFolder(ctx context.Context, folder *folder.Folder) error {

	folderDTO := folder2.ToDTO(folder)

	if err := s.repo.CreateFolder(ctx, folderDTO); err != nil {
		s.log.Error(err)
		return err
	}
	return nil
}

func (s service) GetFolderByID(ctx context.Context, id folder.FolderID) (*folder.Folder, error) {
	folderDTO, err := s.repo.GetFolderByID(ctx, folder2.FolderDTOID(id))
	if err != nil {
		s.log.Error(err)
		return nil, err

	}
	return folderDTO.ToDomain(), nil
}

func (s service) GetFoldersByParentID(ctx context.Context, parentID folder.FolderID) ([]*folder.Folder, error) {
	folderDTOs, err := s.repo.GetFoldersByParentID(ctx, folder2.FolderDTOID(parentID))
	if err != nil {
		s.log.Error(err)
		return nil, err
	}
	folders := make([]*folder.Folder, len(folderDTOs))
	for i, f := range folderDTOs {
		folders[i] = f.ToDomain()
	}
	return folders, nil
}

func (s service) UpdateFolder(ctx context.Context, folder *folder.Folder) error {
	folderDTO := folder2.ToDTO(folder)
	if err := s.repo.UpdateFolder(ctx, folderDTO); err != nil {
		s.log.Error(err)
		return err
	}
	return nil
}

func (s service) DeleteFolder(ctx context.Context, id folder.FolderID) error {
	if err := s.repo.DeleteFolder(ctx, folder2.FolderDTOID(id)); err != nil {
		s.log.Error(err)
		return err
	}
	return nil
}

func NewService(repo folder2.Repository, log *logging.Logger) Service {
	return &service{
		repo: repo,
		log:  log,
	}
}

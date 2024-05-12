package file

import (
	"context"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	filerep "remy_explorer/internal/explorer/repository/file"
)

// Service provides file operations
type Service interface {
	CreateFile(ctx context.Context, f *File) (*int64, error)
	GetFileByID(ctx context.Context, id int64) (*File, error)
	GetFilesByFolderID(ctx context.Context, parentID int64) ([]*File, error)
	UpdateFile(ctx context.Context, f *File) (bool, error)
	DeleteFile(ctx context.Context, id int64) (bool, error)
}

type service struct {
	repo filerep.Repository
	log  log.Logger
}

func (s service) CreateFile(ctx context.Context, f *File) (*int64, error) {
	logger := log.With(s.log, "folder", "UpdateFolder")
	fileDTO := filerep.ToDTO(f)
	id, err := s.repo.CreateFile(ctx, &fileDTO)
	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, err
	}
	logger.Log("message", "File created", "id", fileDTO.ID)
	return id, nil
}

func (s service) GetFileByID(ctx context.Context, id int64) (*File, error) {
	logger := log.With(s.log, "folder", "GetFolderByID")
	fileDTO, err := s.repo.GetFileByID(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, err
	}
	file := fileDTO.ToDomain()
	logger.Log("message", "File retrieved", "id", file.ID)
	return file, nil
}

func (s service) GetFilesByFolderID(ctx context.Context, parentID int64) ([]*File, error) {
	logger := log.With(s.log, "folder", "GetFoldersByParentID")
	fileDTOs, err := s.repo.GetFilesByFolderID(ctx, parentID)
	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, err
	}
	files := make([]*File, len(fileDTOs))
	for i, f := range fileDTOs {
		files[i] = f.ToDomain()
	}
	logger.Log("message", "Files retrieved", "count", len(files))
	return files, nil
}

func (s service) UpdateFile(ctx context.Context, f *File) (bool, error) {
	logger := log.With(s.log, "folder", "UpdateFolder")
	fileDTO := filerep.ToDTO(f)
	if err := s.repo.UpdateFile(ctx, &fileDTO); err != nil {
		level.Error(logger).Log("err", err)
		return false, err
	}
	logger.Log("message", "File updated", "id", fileDTO.ID)
	return true, nil
}

func (s service) DeleteFile(ctx context.Context, id int64) (bool, error) {
	logger := log.With(s.log, "folder", "DeleteFolder")
	if err := s.repo.DeleteFile(ctx, id); err != nil {
		level.Error(logger).Log("err", err)
		return false, err
	}
	logger.Log("message", "File deleted", "id", id)
	return true, nil
}

func NewService(repo filerep.Repository, logger log.Logger) Service {
	return &service{
		repo: repo,
		log:  log.With(logger, "service", "file"),
	}
}

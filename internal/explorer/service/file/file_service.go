package file

import (
	"context"
	"errors"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"remy_explorer/internal/explorer/dto"
	modelerr "remy_explorer/internal/explorer/err"
	"remy_explorer/internal/explorer/model"
)

// FileService provides file operations
type FileService interface {
	CreateFile(ctx context.Context, f *model.File) (*string, error)
	GetFileByID(ctx context.Context, id string) (*model.File, error)
	GetFilesByFolderID(ctx context.Context, parentID string) ([]*model.File, error)
	UpdateFile(ctx context.Context, f *model.File) (bool, error)
	DeleteFile(ctx context.Context, id string) (bool, error)
}

type service struct {
	repo dto.FileRepository
	log  log.Logger
}

func (s service) CreateFile(ctx context.Context, f *model.File) (*string, error) {
	logger := log.With(s.log, "folder", "UpdateFolder")
	fileDTO := dto.FileToDTO(f)
	id, err := s.repo.CreateFile(ctx, &fileDTO)
	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, err
	}
	logger.Log("message", "File created", "id", fileDTO.ID)
	return id, nil
}

func (s service) GetFileByID(ctx context.Context, id string) (*model.File, error) {
	logger := log.With(s.log, "method", "GetFileByID")
	fileDTO, err := s.repo.GetFileByID(ctx, id)
	if err != nil {
		var errNotFound *modelerr.NotFound
		if errors.As(err, &errNotFound) {
			level.Info(logger).Log("err", err, "msg", "folder not found")
			return nil, err
		}
		level.Error(logger).Log("err", err, "msg", "failed to retrieve file")
		return nil, err
	}
	file := fileDTO.ToDomain()
	logger.Log("message", "File retrieved", "id", file.ID)
	return file, nil
}

func (s service) GetFilesByFolderID(ctx context.Context, parentID string) ([]*model.File, error) {
	logger := log.With(s.log, "folder", "GetFoldersByParentID")
	fileDTOs, err := s.repo.GetFilesByFolderID(ctx, parentID)
	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, err
	}
	files := make([]*model.File, len(fileDTOs))
	for i, f := range fileDTOs {
		files[i] = f.ToDomain()
	}
	logger.Log("message", "Files retrieved", "count", len(files))
	return files, nil
}

func (s service) UpdateFile(ctx context.Context, f *model.File) (bool, error) {
	logger := log.With(s.log, "folder", "UpdateFolder")
	fileDTO := dto.FileToDTO(f)
	if err := s.repo.UpdateFile(ctx, &fileDTO); err != nil {
		level.Error(logger).Log("err", err)
		return false, err
	}
	logger.Log("message", "File updated", "id", fileDTO.ID)
	return true, nil
}

func (s service) DeleteFile(ctx context.Context, id string) (bool, error) {
	logger := log.With(s.log, "folder", "DeleteFolder")
	if err := s.repo.DeleteFile(ctx, id); err != nil {
		level.Error(logger).Log("err", err)
		return false, err
	}
	logger.Log("message", "File deleted", "id", id)
	return true, nil
}

func NewService(repo dto.FileRepository, logger log.Logger) FileService {
	return &service{
		repo: repo,
		log:  log.With(logger, "service", "file"),
	}
}

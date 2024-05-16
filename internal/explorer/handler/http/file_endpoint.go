package http

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"remy_explorer/internal/explorer/domain"
	fileSchemas "remy_explorer/internal/explorer/handler/http/schemas"
	"remy_explorer/internal/explorer/service/file"
)

func makeCreateFileEndpoint(s file.FileService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(fileSchemas.CreateFileRequest)
		f := domain.File{
			Name:       req.Name,
			FolderID:   req.FolderID,
			OwnerID:    req.OwnerID,
			Size:       req.Size,
			Type:       req.Type,
			ObjectPath: req.Path,
		}
		id, err := s.CreateFile(ctx, &f)
		return fileSchemas.CreateFileResponse{ID: *id}, err
	}
}

func makeGetFileByIDEndpoint(s file.FileService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(fileSchemas.GetFileByIDRequest)
		f, err := s.GetFileByID(ctx, req.ID)
		return fileSchemas.GetFileByIDResponse{
			ID:        f.ID,
			Name:      f.Name,
			FolderID:  f.FolderID,
			Size:      f.Size,
			Type:      f.Type,
			Path:      f.ObjectPath,
			CreatedAt: f.CreatedAt.String(),
			UpdatedAt: f.UpdatedAt.String(),
			Tags:      f.Tags,
		}, err
	}
}

func makeGetFilesByParentIDEndpoint(s file.FileService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(fileSchemas.GetFilesByFolderIDRequest)
		files, err := s.GetFilesByFolderID(ctx, req.FolderID)
		length := len(files) // Used two times
		shortFiles := make([]fileSchemas.ShortFileInfo, length)
		for i, f := range files {
			shortFiles[i] = fileSchemas.ShortFileInfo{
				ID:   f.ID,
				Name: f.Name,
				Type: f.Type,
			}
		}
		return fileSchemas.GetFilesByFolderIDResponse{
			Length: length,
			Files:  shortFiles,
		}, err
	}
}

func makeUpdateFileEndpoint(s file.FileService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(fileSchemas.UpdateFileRequest)
		f := domain.File{
			ID:       req.ID,
			Name:     req.Name,
			FolderID: req.FolderID,
		}
		ok, err := s.UpdateFile(ctx, &f)
		return fileSchemas.UpdateFileResponse{Ok: ok}, err
	}
}

func makeDeleteFileEndpoint(s file.FileService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(fileSchemas.DeleteFileRequest)
		ok, err := s.DeleteFile(ctx, req.ID)
		return fileSchemas.DeleteFileResponse{Ok: ok}, err
	}
}

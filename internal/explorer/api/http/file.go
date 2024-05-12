package http

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	file2 "remy_explorer/internal/explorer/api/http/schemas"
	"remy_explorer/internal/explorer/service/file"
)

func makeCreateFileEndpoint(s file.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(file2.CreateFileRequest)
		f := file.File{
			Name:       req.Name,
			FolderID:   req.FolderID,
			OwnerID:    req.OwnerID,
			Size:       req.Size,
			Type:       req.Type,
			ObjectPath: req.Path,
		}
		id, err := s.CreateFile(ctx, &f)
		return file2.CreateFileResponse{ID: *id}, err
	}
}

func makeGetFileByIDEndpoint(s file.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(file2.GetFileByIDRequest)
		f, err := s.GetFileByID(ctx, req.ID)
		return file2.GetFileByIDResponse{
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

func makeGetFilesByParentIDEndpoint(s file.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(file2.GetFilesByFolderIDRequest)
		files, err := s.GetFilesByFolderID(ctx, req.FolderID)
		length := len(files) // Used two times
		shortFiles := make([]file2.ShortFileInfo, length)
		for i, f := range files {
			shortFiles[i] = file2.ShortFileInfo{
				ID:   f.ID,
				Name: f.Name,
				Type: f.Type,
			}
		}
		return file2.GetFilesByFolderIDResponse{
			Length: length,
			Files:  shortFiles,
		}, err
	}
}

func makeUpdateFileEndpoint(s file.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(file2.UpdateFileRequest)
		f := file.File{
			ID:       req.ID,
			Name:     req.Name,
			FolderID: req.FolderID,
		}
		ok, err := s.UpdateFile(ctx, &f)
		return file2.UpdateFileResponse{Ok: ok}, err
	}
}

func makeDeleteFileEndpoint(s file.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(file2.DeleteFileRequest)
		ok, err := s.DeleteFile(ctx, req.ID)
		return file2.DeleteFileResponse{Ok: ok}, err
	}
}

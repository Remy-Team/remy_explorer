package file

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"remy_explorer/internal/explorer/service/file"
)

// Endpoints holds all Go kit endpoints for file operations
type Endpoints struct {
	// For files
	CreateFile         endpoint.Endpoint
	GetFileByID        endpoint.Endpoint
	GetFilesByParentID endpoint.Endpoint
	UpdateFile         endpoint.Endpoint
	DeleteFile         endpoint.Endpoint
}

// MakeEndpoints initializes all Go kit endpoints for file operations
func MakeEndpoints(s file.Service) Endpoints {
	return Endpoints{
		CreateFile:         makeCreateFileEndpoint(s),
		GetFileByID:        makeGetFileByIDEndpoint(s),
		GetFilesByParentID: makeGetFilesByParentIDEndpoint(s),
		UpdateFile:         makeUpdateFileEndpoint(s),
		DeleteFile:         makeDeleteFileEndpoint(s),
	}
}

func makeCreateFileEndpoint(s file.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateFileRequest)
		f := file.File{
			Name:       req.Name,
			FolderID:   req.FolderID,
			OwnerID:    req.OwnerID,
			Size:       req.Size,
			Type:       req.Type,
			ObjectPath: req.Path,
		}
		id, err := s.CreateFile(ctx, &f)
		return CreateFileResponse{ID: *id}, err
	}
}

func makeGetFileByIDEndpoint(s file.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetFileByIDRequest)
		f, err := s.GetFileByID(ctx, req.ID)
		return GetFileByIDResponse{
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
		req := request.(GetFilesByFolderIDRequest)
		files, err := s.GetFilesByFolderID(ctx, req.FolderID)
		length := len(files) // Used two times
		shortFiles := make([]ShortFileInfo, length)
		for i, f := range files {
			shortFiles[i] = ShortFileInfo{
				ID:   f.ID,
				Name: f.Name,
				Type: f.Type,
			}
		}
		return GetFilesByFolderIDResponse{
			length: length,
			Files:  shortFiles,
		}, err
	}
}

func makeUpdateFileEndpoint(s file.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateFileRequest)
		f := file.File{
			ID:       req.ID,
			Name:     req.Name,
			FolderID: req.FolderID,
		}
		ok, err := s.UpdateFile(ctx, &f)
		return UpdateFileResponse{Ok: ok}, err
	}
}

func makeDeleteFileEndpoint(s file.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteFileRequest)
		ok, err := s.DeleteFile(ctx, req.ID)
		return DeleteFileResponse{Ok: ok}, err
	}
}

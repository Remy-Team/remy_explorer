package api

// Write endpooints with go-kit
import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"remy_explorer/internal/explorer/service"
	"remy_explorer/pkg/domain/folder"
	"remy_explorer/pkg/domain/user"
)

// Endpoints holds all Go kit endpoints for folder operations
type Endpoints struct {
	// For folders
	CreateFolder         endpoint.Endpoint
	GetFolderByID        endpoint.Endpoint
	GetFoldersByParentID endpoint.Endpoint
	UpdateFolder         endpoint.Endpoint
	DeleteFolder         endpoint.Endpoint
}

// MakeEndpoints initializes all Go kit endpoints for folder operations
func MakeEndpoints(s service.Service) Endpoints {
	return Endpoints{
		CreateFolder:         makeCreateFolderEndpoint(s),
		GetFolderByID:        makeGetFolderByIDEndpoint(s),
		GetFoldersByParentID: makeGetFoldersByParentIDEndpoint(s),
		UpdateFolder:         makeUpdateFolderEndpoint(s),
		DeleteFolder:         makeDeleteFolderEndpoint(s),
	}
}

func makeCreateFolderEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateFolderRequest)
		f := folder.Folder{
			Name:     req.Name,
			OwnerID:  user.ID(req.OwnerID),
			ParentID: folder.FolderID(req.ParentID),
		}
		id, err := s.CreateFolder(ctx, &f)
		return CreateFolderResponse{id.ToInt64()}, err
	}
}

func makeGetFolderByIDEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetFolderByIDRequest)
		f, err := s.GetFolderByID(ctx, folder.FolderID(req.ID))
		return GetFolderByIDResponse{
			ID:        f.ID.ToInt64(),
			OwnerID:   string(f.OwnerID),
			Name:      f.Name,
			ParentID:  f.ParentID.ToInt64(),
			CreatedAt: f.CreatedAt.String(),
			UpdatedAt: f.UpdatedAt.String(),
		}, err
	}
}

func makeGetFoldersByParentIDEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetFoldersByParentIDRequest)
		folders, err := s.GetFoldersByParentID(ctx, folder.FolderID(req.ParentID))
		if err != nil {
			return nil, err
		}
		var res []GetFoldersByParentIDResponse
		for _, f := range folders {
			res = append(res, GetFoldersByParentIDResponse{
				ID:   f.ID.ToInt64(),
				Name: f.Name,
			})
		}
		return res, nil

	}
}

func makeUpdateFolderEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateFolderRequest)
		f := folder.Folder{
			ID:       folder.FolderID(req.ID),
			Name:     req.Name,
			ParentID: folder.FolderID(req.ParentID),
		}
		err := s.UpdateFolder(ctx, &f)
		return UpdateFolderResponse{Ok: true}, err // TODO: replace  true with returned value
	}

}

func makeDeleteFolderEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteFolderRequest)
		err := s.DeleteFolder(ctx, folder.FolderID(req.ID))
		return DeleteFolderResponse{Ok: true}, err //TODO: replace true with returned value
	}

}

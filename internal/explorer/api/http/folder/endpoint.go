package folder

// Write endpooints with go-kit
import (
	"context"
	"github.com/go-kit/kit/endpoint"
	folder2 "remy_explorer/internal/explorer/service/folder"
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
func MakeEndpoints(s folder2.Service) Endpoints {
	return Endpoints{
		CreateFolder:         makeCreateFolderEndpoint(s),
		GetFolderByID:        makeGetFolderByIDEndpoint(s),
		GetFoldersByParentID: makeGetFoldersByParentIDEndpoint(s),
		UpdateFolder:         makeUpdateFolderEndpoint(s),
		DeleteFolder:         makeDeleteFolderEndpoint(s),
	}
}

func makeCreateFolderEndpoint(s folder2.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateFolderRequest)
		f := folder2.Folder{
			Name:     req.Name,
			OwnerID:  req.OwnerID,
			ParentID: req.ParentID,
		}
		id, err := s.CreateFolder(ctx, &f)
		return CreateFolderResponse{*id}, err
	}
}

func makeGetFolderByIDEndpoint(s folder2.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetFolderByIDRequest)
		f, err := s.GetFolderByID(ctx, req.ID)
		return GetFolderByIDResponse{
			ID:        f.ID,
			OwnerID:   string(f.OwnerID),
			Name:      f.Name,
			ParentID:  f.ParentID,
			CreatedAt: f.CreatedAt.String(),
			UpdatedAt: f.UpdatedAt.String(),
		}, err
	}
}

func makeGetFoldersByParentIDEndpoint(s folder2.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetFoldersByParentIDRequest)
		folders, err := s.GetFoldersByParentID(ctx, req.ParentID)
		if err != nil {
			return nil, err
		}
		var res []GetFoldersByParentIDResponse
		for _, f := range folders {
			res = append(res, GetFoldersByParentIDResponse{
				ID:   f.ID,
				Name: f.Name,
			})
		}
		return res, nil

	}
}

func makeUpdateFolderEndpoint(s folder2.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateFolderRequest)
		f := folder2.Folder{
			ID:       req.ID,
			Name:     req.Name,
			ParentID: req.ParentID,
		}
		err := s.UpdateFolder(ctx, &f)
		return UpdateFolderResponse{Ok: err == nil}, err // TODO: replace true with returned value
	}

}

func makeDeleteFolderEndpoint(s folder2.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteFolderRequest)
		err := s.DeleteFolder(ctx, req.ID)
		return DeleteFolderResponse{Ok: err == nil}, err //TODO: replace true with returned value
	}

}

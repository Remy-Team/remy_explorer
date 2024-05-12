package http

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"remy_explorer/internal/explorer/api/http/schemas"
	folder2 "remy_explorer/internal/explorer/service/folder"
)

func makeCreateFolderEndpoint(s folder2.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(schemas.CreateFolderRequest)
		f := folder2.Folder{
			Name:     req.Name,
			OwnerID:  req.OwnerID,
			ParentID: req.ParentID,
		}
		id, err := s.CreateFolder(ctx, &f)
		return schemas.CreateFolderResponse{ID: *id}, err
	}
}

func makeGetFolderByIDEndpoint(s folder2.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(schemas.GetFolderByIDRequest)
		f, err := s.GetFolderByID(ctx, req.ID)
		return schemas.GetFolderByIDResponse{
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
		req := request.(schemas.GetFoldersByParentIDRequest)
		folders, err := s.GetFoldersByParentID(ctx, req.ParentID)
		if err != nil {
			return nil, err
		}
		var res []schemas.GetFoldersByParentIDResponse
		for _, f := range folders {
			res = append(res, schemas.GetFoldersByParentIDResponse{
				ID:   f.ID,
				Name: f.Name,
			})
		}
		return res, nil

	}
}

func makeUpdateFolderEndpoint(s folder2.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(schemas.UpdateFolderRequest)
		f := folder2.Folder{
			ID:       req.ID,
			Name:     req.Name,
			ParentID: req.ParentID,
		}
		err := s.UpdateFolder(ctx, &f)
		return schemas.UpdateFolderResponse{Ok: err == nil}, err // TODO: replace true with returned value
	}

}

func makeDeleteFolderEndpoint(s folder2.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(schemas.DeleteFolderRequest)
		err := s.DeleteFolder(ctx, req.ID)
		return schemas.DeleteFolderResponse{Ok: err == nil}, err //TODO: replace true with returned value
	}

}

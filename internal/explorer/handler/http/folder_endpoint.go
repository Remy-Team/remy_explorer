package http

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"remy_explorer/internal/explorer/domain"
	"remy_explorer/internal/explorer/handler/http/schemas"
)

func makeCreateFolderEndpoint(s domain.FolderService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(schemas.CreateFolderRequest)
		f := domain.Folder{
			Name:     req.Name,
			OwnerID:  req.OwnerID,
			ParentID: req.ParentID,
		}
		id, err := s.CreateFolder(ctx, &f)
		return schemas.CreateFolderResponse{ID: *id}, err
	}
}

func makeGetFolderByIDEndpoint(s domain.FolderService) endpoint.Endpoint {
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

func makeGetFoldersByParentIDEndpoint(s domain.FolderService) endpoint.Endpoint {
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

func makeUpdateFolderEndpoint(s domain.FolderService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(schemas.UpdateFolderRequest)
		f := domain.Folder{
			ID:       req.ID,
			Name:     req.Name,
			ParentID: req.ParentID,
		}
		err := s.UpdateFolder(ctx, &f)
		return schemas.UpdateFolderResponse{Ok: err == nil}, err // TODO: replace true with returned value
	}

}

func makeDeleteFolderEndpoint(s domain.FolderService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(schemas.DeleteFolderRequest)
		err := s.DeleteFolder(ctx, req.ID)
		return schemas.DeleteFolderResponse{Ok: err == nil}, err //TODO: replace true with returned value
	}

}

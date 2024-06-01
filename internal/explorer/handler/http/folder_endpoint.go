package http

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"remy_explorer/internal/explorer/domain"
	"remy_explorer/internal/explorer/handler/http/schemas"
	"remy_explorer/internal/explorer/service/folder"
)

// makeCreateFolderEndpoint creates an endpoint for creating a folder
//
//	@Summary		Create a new folder
//	@Description	Create a new folder in the system
//	@Tags			folders
//	@Accept			json
//	@Produce		json
//	@Param			body	body		schemas.CreateFolderRequest	true	"Create Folder Request"
//	@Success		200		{object}	schemas.CreateFolderResponse
//	@Failure		400		{object}	schemas.ErrorResponse
//	@Failure		500		{object}	schemas.ErrorResponse
//	@Router			/folders [post]
func makeCreateFolderEndpoint(logger log.Logger, s folder.FolderService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		level.Info(logger).Log("msg", "entering  makeCreateFolderEndpoint", "request", request)
		req, ok := request.(schemas.CreateFolderRequest)
		if !ok {
			return nil, errors.New("invalid request type")
		}
		f := domain.Folder{
			Name:     req.Name,
			OwnerID:  req.OwnerID,
			ParentID: req.ParentID,
		}
		id, err := s.CreateFolder(ctx, &f)
		return schemas.CreateFolderResponse{ID: *id}, err
	}
}

// makeGetFolderByIDEndpoint creates an endpoint for getting a folder by ID
//
//	@Summary		Get folder by ID
//	@Description	Retrieve a folder's details by its ID
//	@Tags			folders
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Folder ID"
//	@Success		200	{object}	schemas.GetFolderByIDResponse
//	@Failure		404	{object}	schemas.ErrorResponse
//	@Failure		500	{object}	schemas.ErrorResponse
//	@Router			/folders/{id} [get]
func makeGetFolderByIDEndpoint(logger log.Logger, s folder.FolderService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		level.Info(logger).Log("msg", "entering  makeGetFolderByIDEndpoint", "request", request)
		req, ok := request.(schemas.GetFolderByIDRequest)
		if !ok {
			return nil, errors.New("invalid request type")
		}
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

// makeGetFoldersByParentIDEndpoint creates an endpoint for getting folders by parent ID
//
//	@Summary		Get folders by parent ID
//	@Description	Retrieve a list of folders within a specific parent folder
//	@Tags			folders
//	@Accept			json
//	@Produce		json
//	@Param			parentID	path		string	true	"Parent Folder ID"
//	@Success		200			{array}		schemas.GetFoldersByParentIDResponse
//	@Failure		404			{object}	schemas.ErrorResponse
//	@Failure		500			{object}	schemas.ErrorResponse
//	@Router			/folders/{parentID}/subfolders [get]
func makeGetFoldersByParentIDEndpoint(logger log.Logger, s folder.FolderService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		level.Info(logger).Log("msg", "entering  makeGetFoldersByParentIDEndpoint", "request", request)
		req, ok := request.(schemas.GetFoldersByParentIDRequest)
		if !ok {
			return nil, errors.New("invalid request type")
		}
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

// makeUpdateFolderEndpoint creates an endpoint for updating a folder
//
//	@Summary		Update a folder
//	@Description	Update the details of an existing folder
//	@Tags			folders
//	@Accept			json
//	@Produce		json
//	@Param			body	body		schemas.UpdateFolderRequest	true	"Update Folder Request"
//	@Success		200		{object}	schemas.UpdateFolderResponse
//	@Failure		400		{object}	schemas.ErrorResponse
//	@Failure		404		{object}	schemas.ErrorResponse
//	@Failure		500		{object}	schemas.ErrorResponse
//	@Router			/folders [put]
func makeUpdateFolderEndpoint(logger log.Logger, s folder.FolderService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		level.Info(logger).Log("msg", "entering  makeUpdateFolderEndpoint", "request", request)
		req, ok := request.(schemas.UpdateFolderRequest)
		if !ok {
			return nil, errors.New("invalid request type")
		}
		f := domain.Folder{
			ID:       req.ID,
			Name:     req.Name,
			ParentID: req.ParentID,
		}
		err := s.UpdateFolder(ctx, &f)
		return schemas.UpdateFolderResponse{Ok: err == nil}, err
	}
}

// makeDeleteFolderEndpoint creates an endpoint for deleting a folder
//
//	@Summary		Delete a folder
//	@Description	Delete a folder by its ID
//	@Tags			folders
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Folder ID"
//	@Success		200	{object}	schemas.DeleteFolderResponse
//	@Failure		404	{object}	schemas.ErrorResponse
//	@Failure		500	{object}	schemas.ErrorResponse
//	@Router			/folders/{id} [delete]
func makeDeleteFolderEndpoint(logger log.Logger, s folder.FolderService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		level.Info(logger).Log("msg", "entering  makeDeleteFolderEndpoint", "request", request)
		req, ok := request.(schemas.DeleteFolderRequest)
		if !ok {
			return nil, errors.New("invalid request type")
		}
		err := s.DeleteFolder(ctx, req.ID)
		return schemas.DeleteFolderResponse{Ok: err == nil}, err
	}
}

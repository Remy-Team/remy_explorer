package http

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"remy_explorer/internal/explorer/domain"
	schemas "remy_explorer/internal/explorer/handler/http/schemas"
	"remy_explorer/internal/explorer/service/file"
)

// makeCreateFileEndpoint creates an endpoint for creating a file
//
//	@Summary		Create a new file
//	@Description	Create a new file in the system
//	@Tags			files
//	@Accept			json
//	@Produce		json
//	@Param			body	body		schemas.CreateFileRequest	true	"Create File Request"
//	@Success		200		{object}	schemas.CreateFileResponse
//	@Failure		400		{object} 	schemas.ErrorResponse
//	@Failure		500		{object} 	schemas.ErrorResponse
//	@Router			/files [post]
func makeCreateFileEndpoint(s file.FileService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(schemas.CreateFileRequest)
		f := domain.File{
			Name:       req.Name,
			FolderID:   req.FolderID,
			OwnerID:    req.OwnerID,
			Size:       req.Size,
			Type:       req.Type,
			ObjectPath: req.Path,
		}
		id, err := s.CreateFile(ctx, &f)
		return schemas.CreateFileResponse{ID: *id}, err
	}
}

// makeGetFileByIDEndpoint creates an endpoint for getting a file by ID
//
//	@Summary		Get file by ID
//	@Description	Retrieve a file's details by its ID
//	@Tags			files
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"File ID"
//	@Success		200	{object}	schemas.GetFileByIDResponse
//	@Failure		404	{object} 	schemas.ErrorResponse
//	@Failure		500	{object} 	schemas.ErrorResponse
//	@Router			/files/{id} [get]
func makeGetFileByIDEndpoint(s file.FileService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(schemas.GetFileByIDRequest)
		f, err := s.GetFileByID(ctx, req.ID)
		return schemas.GetFileByIDResponse{
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

// makeGetFilesByParentIDEndpoint creates an endpoint for getting files by folder ID
//
//	@Summary		Get files by folder ID
//	@Description	Retrieve a list of files in a specific folder
//	@Tags			files
//	@Accept			json
//	@Produce		json
//	@Param			folderID	path		string	true	"Folder ID"
//	@Success		200			{object}	schemas.GetFilesByFolderIDResponse
//	@Failure		404			{object} 	schemas.ErrorResponse
//	@Failure		500			{object} 	schemas.ErrorResponse
//	@Router			/folders/{folderID}/files [get]
func makeGetFilesByParentIDEndpoint(s file.FileService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(schemas.GetFilesByFolderIDRequest)
		files, err := s.GetFilesByFolderID(ctx, req.FolderID)
		length := len(files) // Used two times
		shortFiles := make([]schemas.ShortFileInfo, length)
		for i, f := range files {
			shortFiles[i] = schemas.ShortFileInfo{
				ID:   f.ID,
				Name: f.Name,
				Type: f.Type,
			}
		}
		return schemas.GetFilesByFolderIDResponse{
			Length: length,
			Files:  shortFiles,
		}, err
	}
}

// makeUpdateFileEndpoint creates an endpoint for updating a file
//
//	@Summary		Update a file
//	@Description	Update the details of an existing file
//	@Tags			files
//	@Accept			json
//	@Produce		json
//	@Param			body	body		schemas.UpdateFileRequest	true	"Update File Request"
//	@Success		200		{object}	schemas.UpdateFileResponse
//	@Failure		400		{object} 	schemas.ErrorResponse
//	@Failure		404		{object} 	schemas.ErrorResponse
//	@Failure		500		{object} 	schemas.ErrorResponse
//	@Router			/files [put]
func makeUpdateFileEndpoint(s file.FileService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(schemas.UpdateFileRequest)
		f := domain.File{
			ID:       req.ID,
			Name:     req.Name,
			FolderID: req.FolderID,
		}
		ok, err := s.UpdateFile(ctx, &f)
		return schemas.UpdateFileResponse{Ok: ok}, err
	}
}

// makeDeleteFileEndpoint creates an endpoint for deleting a file
//
//	@Summary		Delete a file
//	@Description	Delete a file by its ID
//	@Tags			files
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"File ID"
//	@Success		200	{object}	schemas.DeleteFileResponse
//	@Failure		404	{object} 	schemas.ErrorResponse
//	@Failure		500	{object} 	schemas.ErrorResponse
//	@Router			/files/{id} [delete]
func makeDeleteFileEndpoint(s file.FileService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(schemas.DeleteFileRequest)
		ok, err := s.DeleteFile(ctx, req.ID)
		return schemas.DeleteFileResponse{Ok: ok}, err
	}
}

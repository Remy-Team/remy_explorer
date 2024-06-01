package http

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"remy_explorer/internal/explorer/service/file"
	"remy_explorer/internal/explorer/service/folder"
)

// Endpoints holds all Go kit endpoints for file operations
type Endpoints struct {
	// For files
	CreateFile         endpoint.Endpoint
	GetFileByID        endpoint.Endpoint
	GetFilesByParentID endpoint.Endpoint
	UpdateFile         endpoint.Endpoint
	DeleteFile         endpoint.Endpoint
	//Folder endpoints
	CreateFolder         endpoint.Endpoint
	GetFolderByID        endpoint.Endpoint
	GetFoldersByParentID endpoint.Endpoint
	UpdateFolder         endpoint.Endpoint
	DeleteFolder         endpoint.Endpoint
}

// MakeEndpoints initializes all Go kit endpoints for file operations
func MakeEndpoints(logger log.Logger, fileS file.FileService, folderS folder.FolderService) Endpoints {
	return Endpoints{
		CreateFile:         makeCreateFileEndpoint(logger, fileS),
		GetFileByID:        makeGetFileByIDEndpoint(logger, fileS),
		GetFilesByParentID: makeGetFilesByParentIDEndpoint(logger, fileS),
		UpdateFile:         makeUpdateFileEndpoint(logger, fileS),
		DeleteFile:         makeDeleteFileEndpoint(logger, fileS),
		// Folder endpoints
		CreateFolder:         makeCreateFolderEndpoint(logger, folderS),
		GetFolderByID:        makeGetFolderByIDEndpoint(logger, folderS),
		GetFoldersByParentID: makeGetFoldersByParentIDEndpoint(logger, folderS),
		UpdateFolder:         makeUpdateFolderEndpoint(logger, folderS),
		DeleteFolder:         makeDeleteFolderEndpoint(logger, folderS),
	}
}

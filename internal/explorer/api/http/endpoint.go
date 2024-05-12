package http

import (
	"github.com/go-kit/kit/endpoint"
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
func MakeEndpoints(fileS file.Service, folderS folder.Service) Endpoints {
	return Endpoints{
		CreateFile:         makeCreateFileEndpoint(fileS),
		GetFileByID:        makeGetFileByIDEndpoint(fileS),
		GetFilesByParentID: makeGetFilesByParentIDEndpoint(fileS),
		UpdateFile:         makeUpdateFileEndpoint(fileS),
		DeleteFile:         makeDeleteFileEndpoint(fileS),
		// Folder endpoints
		CreateFolder:         makeCreateFolderEndpoint(folderS),
		GetFolderByID:        makeGetFolderByIDEndpoint(folderS),
		GetFoldersByParentID: makeGetFoldersByParentIDEndpoint(folderS),
		UpdateFolder:         makeUpdateFolderEndpoint(folderS),
		DeleteFolder:         makeDeleteFolderEndpoint(folderS),
	}
}

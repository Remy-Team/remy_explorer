package http

import (
	"github.com/go-kit/kit/endpoint"
	"remy_explorer/internal/explorer/domain"
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
func MakeEndpoints(fileS domain.FileService, folderS domain.FolderService) Endpoints {
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

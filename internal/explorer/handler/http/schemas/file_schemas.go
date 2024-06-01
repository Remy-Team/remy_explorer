package schemas

// CreateFileRequest represents the request to create a new file
type CreateFileRequest struct {
	Name     string `json:"name" validate:"required"` // Name of the file
	Type     string `json:"type"`                     // Type of the file
	FolderID string `json:"folder_id"`                // ID of the parent folder
	OwnerID  string `json:"owner_id"`                 // ID of the owner
	Path     string `json:"path"`                     // Path where the file is stored
	Size     int    `json:"size"`                     // Size of the file
}

// CreateFileResponse represents the response after creating a new file
type CreateFileResponse struct {
	ID string `json:"id"` // ID of the created file
}

// GetFileByIDRequest represents the request to get a file by its ID
type GetFileByIDRequest struct {
	ID string `json:"id" validate:"required"` // ID of the file to retrieve
}

// GetFileByIDResponse represents the response with the details of a file
type GetFileByIDResponse struct {
	ID        string   `json:"id"`         // ID of the file
	Name      string   `json:"name"`       // Name of the file
	Type      string   `json:"type"`       // Type of the file
	Size      int      `json:"size"`       // Size of the file
	FolderID  string   `json:"folder_id"`  // ID of the parent folder
	Path      string   `json:"path"`       // Path where the file is stored
	CreatedAt string   `json:"created_at"` // Timestamp when the file was created
	UpdatedAt string   `json:"updated_at"` // Timestamp when the file was last updated
	Tags      []string `json:"tags"`       // Tags associated with the file
}

// GetFilesByFolderIDRequest represents the request to get files by folder ID
type GetFilesByFolderIDRequest struct {
	FolderID string `json:"folder_id" validate:"required"` // ID of the parent folder
}

// ShortFileInfo represents a short version of file information
type ShortFileInfo struct {
	ID   string `json:"id"`   // ID of the file
	Name string `json:"name"` // Name of the file
	Type string `json:"type"` // Type of the file
}

// GetFilesByFolderIDResponse represents the response with the list of files in a folder
type GetFilesByFolderIDResponse struct {
	Length int             `json:"length"` // Number of files
	Files  []ShortFileInfo `json:"files"`  // List of files
}

// UpdateFileRequest represents the request to update a file
type UpdateFileRequest struct {
	ID       string `json:"id" validate:"required"`   // ID of the file to update
	Name     string `json:"name" validate:"required"` // New name of the file
	FolderID string `json:"folder_id"`                // New parent folder ID
}

// UpdateFileResponse represents the response after updating a file
type UpdateFileResponse struct {
	Ok bool `json:"ok"` // Indicates whether the update was successful
}

// DeleteFileRequest represents the request to delete a file
type DeleteFileRequest struct {
	ID string `json:"id" validate:"required"` // ID of the file to delete
}

// DeleteFileResponse represents the response after deleting a file
type DeleteFileResponse struct {
	Ok bool `json:"ok"` // Indicates whether the deletion was successful
}

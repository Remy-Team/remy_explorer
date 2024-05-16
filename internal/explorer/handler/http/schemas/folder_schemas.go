package schemas

// CreateFolderRequest represents the request to create a new folder
type CreateFolderRequest struct {
	Name     string `json:"name" validate:"required"` // Name of the folder
	ParentID int64  `json:"parent_id"`                // ID of the parent folder
	OwnerID  string `json:"owner_id"`                 // ID of the owner
}

// CreateFolderResponse represents the response after creating a new folder
type CreateFolderResponse struct {
	ID int64 `json:"id"` // ID of the created folder
}

// GetFolderByIDRequest represents the request to get a folder by its ID
type GetFolderByIDRequest struct {
	ID int64 `json:"id" validate:"required"` // ID of the folder to retrieve
}

// GetFolderByIDResponse represents the response with the details of a folder
type GetFolderByIDResponse struct {
	ID        int64  `json:"id"`         // ID of the folder
	OwnerID   string `json:"owner_id"`   // ID of the owner
	Name      string `json:"name"`       // Name of the folder
	ParentID  int64  `json:"parent_id"`  // ID of the parent folder
	CreatedAt string `json:"created_at"` // Timestamp when the folder was created
	UpdatedAt string `json:"updated_at"` // Timestamp when the folder was last updated
}

// GetFoldersByParentIDRequest represents the request to get folders by parent ID
type GetFoldersByParentIDRequest struct {
	ParentID int64 `json:"parent_id" validate:"required"` // ID of the parent folder
}

// GetFoldersByParentIDResponse represents the response with the list of folders within a specific parent folder
type GetFoldersByParentIDResponse struct {
	ID   int64  `json:"id"`   // ID of the folder
	Name string `json:"name"` // Name of the folder
}

// UpdateFolderRequest represents the request to update a folder
type UpdateFolderRequest struct {
	ID       int64  `json:"id" validate:"required"`   // ID of the folder to update
	Name     string `json:"name" validate:"required"` // New name of the folder
	ParentID int64  `json:"parent_id"`                // New parent folder ID
}

// UpdateFolderResponse represents the response after updating a folder
type UpdateFolderResponse struct {
	Ok bool `json:"ok"` // Indicates whether the update was successful
}

// DeleteFolderRequest represents the request to delete a folder
type DeleteFolderRequest struct {
	ID int64 `json:"id" validate:"required"` // ID of the folder to delete
}

// DeleteFolderResponse represents the response after deleting a folder
type DeleteFolderResponse struct {
	Ok bool `json:"ok"` // Indicates whether the deletion was successful
}

package folder

// Request and response schemas for Folder entity in the API
type (
	CreateFolderRequest struct {
		Name     string `json:"name" validate:"required"`
		ParentID int64  `json:"parent_id"`
		OwnerID  string `json:"owner_id"`
	}
	CreateFolderResponse struct {
		ID int64 `json:"id"`
	}
	GetFolderByIDRequest struct {
		ID int64 `json:"id" validate:"required"`
	}
	GetFolderByIDResponse struct {
		ID        int64  `json:"id"`
		OwnerID   string `json:"owner_id"`
		Name      string `json:"name"`
		ParentID  int64  `json:"parent_id"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}
	GetFoldersByParentIDRequest struct {
		ParentID int64 `json:"parent_id" validate:"required"`
	}
	GetFoldersByParentIDResponse struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	}
	UpdateFolderRequest struct {
		ID       int64  `json:"id" validate:"required"`
		Name     string `json:"name" validate:"required"`
		ParentID int64  `json:"parent_id"`
	}
	UpdateFolderResponse struct {
		Ok bool `json:"ok"`
	}
	DeleteFolderRequest struct {
		ID int64 `json:"id" validate:"required"`
	}
	DeleteFolderResponse struct {
		Ok bool `json:"ok"`
	}
)

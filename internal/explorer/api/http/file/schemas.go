package file

type (
	CreateFileRequest struct {
		Name     string `json:"name" validate:"required"`
		Type     string `json:"type"`
		FolderID int64  `json:"folder_id"`
		OwnerID  string `json:"owner_id"`
		Path     string `json:"path"`
		Size     int
	}
	CreateFileResponse struct {
		ID int64 `json:"id"`
	}
	GetFileByIDRequest struct {
		ID int64 `json:"id" validate:"required"`
	}
	GetFileByIDResponse struct {
		ID        int64    `json:"id"`
		Name      string   `json:"name"`
		Type      string   `json:"type"`
		Size      int      `json:"size"`
		FolderID  int64    `json:"folder_id"`
		Path      string   `json:"path"`
		CreatedAt string   `json:"created_at"`
		UpdatedAt string   `json:"updated_at"`
		Tags      []string `json:"tags"`
	}
	GetFilesByFolderIDRequest struct {
		FolderID int64 `json:"folder_id" validate:"required"`
	}
	ShortFileInfo struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
		Type string `json:"type"`
	}

	GetFilesByFolderIDResponse struct {
		length int
		Files  []ShortFileInfo `json:"files"`
	}
	UpdateFileRequest struct {
		ID       int64  `json:"id" validate:"required"`
		Name     string `json:"name" validate:"required"`
		FolderID int64  `json:"folder_id"`
	}
	UpdateFileResponse struct {
		Ok bool `json:"ok"`
	}
	DeleteFileRequest struct {
		ID int64 `json:"id" validate:"required"`
	}
	DeleteFileResponse struct {
		Ok bool `json:"ok"`
	}
)

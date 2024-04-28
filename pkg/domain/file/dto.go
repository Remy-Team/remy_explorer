package file

type CreateFileDTO struct {
	OwnerID    int64  `json:"owner_id"`
	Name       string `json:"name"`
	FolderID   int64  `json:"folder_id"`
	ObjectPath string `json:"object_path"`
	Size       int    `json:"size"`
	Type       string `json:"type"`
}

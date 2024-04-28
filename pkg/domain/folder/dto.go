package folder

type CreateFolderDTO struct {
	OwnerID int64  `json:"owner_id"`
	Name    string `json:"name"`
	Parent  int64  `json:"parent_id"`
}

type GetFolderByIDDTO struct {
	ID int64 `json:"id"`
}

type GetFoldersByParentIDDTO struct {
	ParentID int64 `json:"parent_id"`
}

type UpdateFolderDTO struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	ParentID int64  `json:"parent_id"`
}

type DeleteFolderDTO struct {
	ID int64 `json:"id"`
}

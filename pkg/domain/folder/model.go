// Package folder CREATE TABLE IF NOT EXISTS public.folder
// (
//
//	id         BIGSERIAL PRIMARY KEY,
//	owner_id   BIGINT                  NOT NULL,
//	name       VARCHAR(255)            NOT NULL,
//	parent_id  BIGINT,
//	created_at TIMESTAMP DEFAULT NOW() NOT NULL,
//	updated_at TIMESTAMP DEFAULT NOW() NOT NULL,
//	CONSTRAINT fk_folder_parent_id FOREIGN KEY (parent_id) REFERENCES public.folder (id)
//
// );
package folder

import (
	"remy_explorer/pkg/domain/user"
	"time"
)

type FolderID int64

func (i FolderID) ToInt64() int64 {
	return int64(i)
}

type Folder struct {
	ID        FolderID  `json:"id"`
	OwnerID   user.ID   `json:"owner_id"`
	Name      string    `json:"name"`
	ParentID  FolderID  `json:"parent"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

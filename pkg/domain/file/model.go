// Package file CREATE TABLE IF NOT EXISTS public.file
// (
//
//	id          BIGSERIAL PRIMARY KEY,
//	owner_id    BIGINT                  NOT NULL,
//	name        VARCHAR(255)            NOT NULL,
//	folder_id   BIGINT                  NOT NULL,
//	object_path VARCHAR(255)            NOT NULL,
//	size        INT                     NOT NULL,
//	type        VARCHAR(255)            NOT NULL,
//	created_at  TIMESTAMP DEFAULT NOW() NOT NULL,
//	updated_at  TIMESTAMP DEFAULT NOW() NOT NULL,
//	tags        JSONB,
//	CONSTRAINT fk_file_folder_id FOREIGN KEY (folder_id) REFERENCES public.folder (id)
//
// );
package file

import (
	"remy_explorer/pkg/domain/folder"
	"remy_explorer/pkg/domain/user"
	"time"
)

type FileID int64

type File struct {
	ID         FileID          `json:"id"`
	OwnerID    user.ID         `json:"owner"`
	Name       string          `json:"name"`
	FolderID   folder.FolderID `json:"folder"`
	ObjectPath string          `json:"object_path"`
	Size       int             `json:"size"`
	Type       string          `json:"type"`
	CreatedAt  time.Time       `json:"created_at"`
	UpdatedAt  time.Time       `json:"updated_at"`
	Tags       []string        `json:"tags"`
}

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
package domain

import (
	"context"
	"time"
)

type File struct {
	ID         int64     `json:"id"`
	OwnerID    string    `json:"owner"`
	Name       string    `json:"name"`
	FolderID   int64     `json:"folder"`
	ObjectPath string    `json:"object_path"`
	Size       int       `json:"size"`
	Type       string    `json:"type"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Tags       []string  `json:"tags"`
}

// FileService provides file operations
type FileService interface {
	CreateFile(ctx context.Context, f *File) (*int64, error)
	GetFileByID(ctx context.Context, id int64) (*File, error)
	GetFilesByFolderID(ctx context.Context, parentID int64) ([]*File, error)
	UpdateFile(ctx context.Context, f *File) (bool, error)
	DeleteFile(ctx context.Context, id int64) (bool, error)
}
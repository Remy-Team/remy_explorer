package file

import (
	"context"
	"database/sql"
	"remy_explorer/internal/config"
	"remy_explorer/pkg/logging"
	"remy_explorer/pkg/postgresql"
	"testing"
)

func prepareTest() (Repository, error) {
	logger := logging.GetLogger()
	cfg := config.GetConfigWithPath("../../../config.yml")
	postgreSQLClient, err := postgresql.New(context.TODO(), cfg.Storage, 5)
	if err != nil {
		return nil, err
	}
	r := New(postgreSQLClient, logger)
	return r, nil
}
func TestGetFilesByFolderIdSorted(t *testing.T) {
	r, err := prepareTest()
	if err != nil {
		return
	}

	sortOption, err := NewSortOption("name", "ASC")
	if err != nil {
		t.Fatalf("Failed to create sort option: %v", err)
	}

	// Assume folderID 1 has multiple files
	files, err := r.GetFilesByFolderIdSorted(context.Background(), 1, sortOption)
	if err != nil {
		t.Errorf("Failed to retrieve sorted files: %v", err)
	}
	if len(files) == 0 {
		t.Errorf("No files retrieved")
	}
}

func TestCreateFile(t *testing.T) {
	r, err := prepareTest()
	if err != nil {
		return
	}

	newFile := FileDTO{
		Name:       "Test File",
		FolderID:   1,
		OwnerID:    1,
		ObjectPath: sql.NullString{String: "path/to/file", Valid: true},
		Size:       1024,
		Type:       sql.NullString{String: "application/pdf", Valid: true},
	}

	err = r.CreateFile(context.Background(), &newFile)
	if err != nil {
		t.Errorf("Failed to create file: %v", err)
	}
	if newFile.ID == 0 {
		t.Errorf("No ID set on new file")
	}
}

func TestGetFileByID(t *testing.T) {
	r, err := prepareTest()
	if err != nil {
		return
	}
	//get id from files array by GetFilesByFolderID
	files, err := r.GetFilesByFolderID(context.Background(), 1)
	if err != nil {
		t.Errorf("Failed to retrieve files: %v", err)
	}
	fileID := files[0].ID

	// Assume we have a file with ID 1
	file, err := r.GetFileByID(context.Background(), fileID)
	if err != nil {
		t.Errorf("Failed to retrieve file: %v", err)
	}
	if file == nil {
		t.Errorf("No file found with ID %v", fileID)
	} else if file.ID != fileID {
		t.Errorf("Retrieved file has incorrect ID: got %v, want %v", file.ID, fileID)
	}
}
func TestGetFilesByFolderID(t *testing.T) {
	r, err := prepareTest()
	if err != nil {
		return
	}

	// Assume folderID 1 has multiple files
	files, err := r.GetFilesByFolderID(context.Background(), 1)
	if err != nil {
		t.Errorf("Failed to retrieve files: %v", err)
	}
	if len(files) == 0 {
		t.Errorf("No files retrieved for folder ID 1")
	}
}

func TestUpdateFile(t *testing.T) {
	r, err := prepareTest()
	if err != nil {
		return
	}

	// Assume this f is already in the database with ID 2
	f := FileDTO{
		Name:       "Test File Updated",
		FolderID:   1,
		OwnerID:    1,
		ObjectPath: sql.NullString{String: "path/to/file/updated", Valid: true},
		Size:       1024,
		Type:       sql.NullString{String: "application/pdf", Valid: true},
	}

	err = r.UpdateFile(context.Background(), &f)
	if err != nil {
		t.Errorf("Failed to update file: %v", err)
	}
}

func TestDeleteFile(t *testing.T) {
	r, err := prepareTest()
	if err != nil {
		return
	}

	// Assume we have a file with ID 3
	err = r.DeleteFile(context.Background(), 1)
	if err != nil {
		t.Errorf("Failed to delete file: %v", err)
	}
}

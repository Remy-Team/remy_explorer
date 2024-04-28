package service_test

import (
	"context"
	"database/sql"
	"remy_explorer/internal/explorer/repository/folder"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"remy_explorer/internal/config"
	"remy_explorer/pkg/logging"
	"remy_explorer/pkg/postgresql"
)

func prepareRepository(t *testing.T) folder.Repository {
	t.Helper()
	logger := logging.GetLogger()
	cfg := config.GetConfigWithPath("../../../config.yml")
	postgreSQLClient, err := postgresql.New(context.TODO(), cfg.Storage, 5)
	require.NoError(t, err)
	return folder.New(postgreSQLClient, logger)
}

func TestCreateFolder(t *testing.T) {
	repo := prepareRepository(t)

	f := folder.FolderDTO{
		Name:     "Test Folder",
		ParentID: sql.NullInt64{Int64: 1, Valid: true},
		OwnerID:  1,
	}
	err := repo.CreateFolder(context.TODO(), &f)
	require.NoError(t, err)
	assert.NotZero(t, f.ID)
}

func TestGetFolderByID(t *testing.T) {
	repo := prepareRepository(t)

	// First create a folder to retrieve
	f := folder.FolderDTO{
		Name:     "New Folder for Retrieval",
		ParentID: sql.NullInt64{Int64: 1, Valid: true},
		OwnerID:  1,
	}
	err := repo.CreateFolder(context.TODO(), &f)
	require.NoError(t, err)
	require.NotZero(t, f.ID)

	// Retrieve the folder
	retrievedFolder, err := repo.GetFolderByID(context.TODO(), f.ID)
	require.NoError(t, err)
	assert.Equal(t, f.Name, retrievedFolder.Name)
	assert.Equal(t, f.ParentID, retrievedFolder.ParentID)
}

func TestGetFoldersByParentID(t *testing.T) {
	repo := prepareRepository(t)

	// Assume ParentID 1 has multiple folders; ensure at least one exists
	f := folder.FolderDTO{
		Name:     "Child Folder",
		ParentID: sql.NullInt64{Int64: 1, Valid: true},
		OwnerID:  1,
	}
	err := repo.CreateFolder(context.TODO(), &f)
	require.NoError(t, err)
	require.NotZero(t, f.ID)

	folders, err := repo.GetFoldersByParentID(context.TODO(), 1)
	require.NoError(t, err)
	assert.NotEmpty(t, folders)
}

func TestUpdateFolder(t *testing.T) {
	repo := prepareRepository(t)

	// Create a folder to update
	f := folder.FolderDTO{
		Name:     "Folder to Update",
		ParentID: sql.NullInt64{Int64: 1, Valid: true},
		OwnerID:  1,
	}
	err := repo.CreateFolder(context.TODO(), &f)
	require.NoError(t, err)
	require.NotZero(t, f.ID)

	// Update the folder
	f.Name = "Updated Name"
	err = repo.UpdateFolder(context.TODO(), &f)
	require.NoError(t, err)

	// Retrieve to verify update
	updatedFolder, err := repo.GetFolderByID(context.TODO(), f.ID)
	require.NoError(t, err)
	assert.Equal(t, "Updated Name", updatedFolder.Name)
}

func TestDeleteFolder(t *testing.T) {
	repo := prepareRepository(t)

	// Create a folder to delete
	f := folder.FolderDTO{
		Name:     "Folder to Delete",
		ParentID: sql.NullInt64{Int64: 1, Valid: true},
		OwnerID:  1,
	}
	err := repo.CreateFolder(context.TODO(), &f)
	require.NoError(t, err)
	require.NotZero(t, f.ID)

	// Delete the folder
	err = repo.DeleteFolder(context.TODO(), f.ID)
	require.NoError(t, err)

	// Try to retrieve the deleted folder
	_, err = repo.GetFolderByID(context.TODO(), f.ID)
	assert.Error(t, err)
}

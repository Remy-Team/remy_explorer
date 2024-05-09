package api

import (
	"context"
	"remy_explorer/internal/config"
	folder2 "remy_explorer/internal/explorer/repository/folder"
	"remy_explorer/pkg/domain/folder"
	"remy_explorer/pkg/logging"
	"remy_explorer/pkg/postgresql"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("Starting explorer service")
	cfg := config.GetConfig()
	logger.Info("Configuration loaded", cfg)

	logger.Info("Connecting to the database")
	postgreSQLClient, err := postgresql.New(context.TODO(), cfg.Storage, 5)
	if err != nil {
		logger.Fatalf("%v", err)
	}
	r := folder2.New(postgreSQLClient, logger)

	// Test the repository methods by creating a folder and retrieving it.
	f := folder.Folder{
		Name:     "Test",
		ParentId: 1,
		OwnerID:  1,
	}
	err = r.CreateFolder(context.TODO(), &f)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info("Folder created", f)

	// Test the GetFolderByID method.
	q, err := r.GetFolderByID(context.TODO(), f.ID)
	if err != nil {
		logger.Fatal(err)

	}
	logger.Info("Folder retrieved", q)

}

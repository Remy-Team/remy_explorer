package file

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
	"remy_explorer/pkg/logging"
	"remy_explorer/pkg/postgresql"
)

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func (r repository) GetFilesByFolderIdSorted(ctx context.Context, folderID FileDTOID, sortOption *SortOption) ([]*FileDTO, error) {

	q := fmt.Sprintf("SELECT id, owner_id, name, folder_id, object_path, size, type, created_at, updated_at, tags FROM public.file WHERE folder_id = $1 ORDER BY %s %s", sortOption.Field, sortOption.Order)
	rows, err := r.client.Query(ctx, q, folderID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf("SQL Error: %s, Details: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())
			r.logger.Error(newErr)
			return nil, nil
		}
		return nil, err

	}
	defer rows.Close()
	result := make([]*FileDTO, 0)
	for rows.Next() {
		var f FileDTO
		if err := rows.Scan(&f.ID, &f.OwnerID, &f.Name, &f.FolderID, &f.ObjectPath, &f.Size, &f.Type, &f.CreatedAt, &f.UpdatedAt, &f.Tags); err != nil {
			return nil, err
		}
		result = append(result, &f)

	}
	return result, nil
}

// CreateFile creates a new file in the database.
func (r repository) CreateFile(ctx context.Context, file *FileDTO) error {
	q := `INSERT INTO public.file (name, folder_id, owner_id, size, type, object_path) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	if err := r.client.QueryRow(ctx, q, file.Name, file.FolderID, file.OwnerID, file.Size, file.Type, file.ObjectPath).Scan(&file.ID); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf("SQL Error: %s, Details: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())
			r.logger.Error(newErr)
			return nil
		}
		return err
	}
	return nil
}

// GetFileByID retrieves a file by its ID.
func (r repository) GetFileByID(ctx context.Context, id FileDTOID) (*FileDTO, error) {
	q := `SELECT id, owner_id, name, folder_id, object_path, size, type, created_at, updated_at, tags FROM public.file WHERE id = $1`
	var f FileDTO
	if err := r.client.QueryRow(ctx, q, id).Scan(&f.ID, &f.OwnerID, &f.Name, &f.FolderID, &f.ObjectPath, &f.Size, &f.Type, &f.CreatedAt, &f.UpdatedAt, &f.Tags); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf("SQL Error: %s, Details: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())
			r.logger.Error(newErr)
			return nil, nil
		}
		return nil, err
	}
	return &f, nil
}

// GetFilesByFolderID retrieves all files with a given folder ID.
func (r repository) GetFilesByFolderID(ctx context.Context, folderID FileDTOID) ([]*FileDTO, error) {
	q := `SELECT id, owner_id, name, folder_id, object_path, size, type, created_at, updated_at, tags FROM public.file WHERE folder_id = $1`
	rows, err := r.client.Query(ctx, q, folderID)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	var files []*FileDTO
	for rows.Next() {
		var f FileDTO
		if err := rows.Scan(&f.ID, &f.OwnerID, &f.Name, &f.FolderID, &f.ObjectPath, &f.Size, &f.Type, &f.CreatedAt, &f.UpdatedAt, &f.Tags); err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}
		files = append(files, &f)
	}
	return files, nil
}

// UpdateFile updates a file in the database.
func (r repository) UpdateFile(ctx context.Context, file *FileDTO) error {
	q := `UPDATE public.file SET name = $1, folder_id = $2, object_path = $3, size = $4, type = $5, updated_at = $6, tags = $7 WHERE id = $8`
	if _, err := r.client.Exec(ctx, q, file.Name, file.FolderID, file.ObjectPath, file.Size, file.Type, file.UpdatedAt, file.Tags, file.ID); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf("SQL Error: %s, Details: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())
			r.logger.Error(newErr)
			return nil
		}
		return err
	}
	return nil
}

// DeleteFile deletes a file from the database.
func (r repository) DeleteFile(ctx context.Context, id FileDTOID) error {
	q := `DELETE FROM public.file WHERE id = $1`
	if _, err := r.client.Exec(ctx, q, id); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf("SQL Error: %s, Details: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())
			r.logger.Error(newErr)
			return nil
		}
		return err
	}
	return nil
}

// New creates a new repository.
func New(client postgresql.Client, logger *logging.Logger) Repository {
	return repository{
		client: client,
		logger: logger,
	}
}

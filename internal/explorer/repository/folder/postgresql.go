package folder

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

// CreateFolder creates a new folder in the database.
func (r repository) CreateFolder(ctx context.Context, folder *FolderDTO) error {
	q := `INSERT INTO public.folder (name, parent_id, owner_id) VALUES ($1, $2, $3) RETURNING id`
	if err := r.client.QueryRow(ctx, q, folder.Name, folder.ParentID, folder.OwnerID).Scan(&folder.ID); err != nil {
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

// GetFolderByID retrieves a folder by its ID.
func (r repository) GetFolderByID(ctx context.Context, id FolderDTOID) (*FolderDTO, error) {
	q := `SELECT id, owner_id, name, parent_id, created_at, updated_at FROM public.folder WHERE id = $1`
	var f FolderDTO
	if err := r.client.QueryRow(ctx, q, id).Scan(&f.ID, &f.OwnerID, &f.Name, &f.ParentID, &f.CreatedAt, &f.UpdatedAt); err != nil {
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

// GetFoldersByParentID retrieves all folders with a given parent ID.
func (r repository) GetFoldersByParentID(ctx context.Context, parentID FolderDTOID) ([]*FolderDTO, error) {
	q := `SELECT id, owner_id, name, parent_id, created_at, updated_at FROM public.folder WHERE parent_id = $1`
	rows, err := r.client.Query(ctx, q, parentID)
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

	var folders []*FolderDTO
	for rows.Next() {
		var f FolderDTO
		if err := rows.Scan(&f.ID, &f.OwnerID, &f.Name, &f.ParentID, &f.CreatedAt, &f.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan folder: %w", err)
		}
		folders = append(folders, &f)
	}
	return folders, nil
}

// UpdateFolder updates a folder in the database.
func (r repository) UpdateFolder(ctx context.Context, folder *FolderDTO) error {
	q := `UPDATE public.folder SET name = $1, parent_id = $2 WHERE id = $3`
	if _, err := r.client.Exec(ctx, q, folder.Name, folder.ParentID, folder.ID); err != nil {
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

// DeleteFolder deletes a folder from the database.

func (r repository) DeleteFolder(ctx context.Context, id FolderDTOID) error {
	q := `DELETE FROM public.folder WHERE id = $1`
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

// New creates a new folder repository.
func New(client postgresql.Client, logger *logging.Logger) repository {
	return repository{
		client: client,
		logger: logger,
	}
}

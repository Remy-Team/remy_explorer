package postgresql

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kit/log"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	model "remy_explorer/internal/explorer/dto"
	modelerr "remy_explorer/internal/explorer/err"
	"strconv"
)

type folderRepository struct {
	client Client
	log    log.Logger
}

// CreateFolder creates a new folder in the database.
func (r folderRepository) CreateFolder(ctx context.Context, folder *model.FolderDTO) (*string, error) {
	q := `INSERT INTO public.folder (name, parent_id, owner_id) VALUES ($1, $2, $3) RETURNING id`
	if err := r.client.QueryRow(ctx, q, folder.Name, folder.ParentID, folder.OwnerID).Scan(&folder.ID); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			newErr := fmt.Errorf("SQL Error: %s, Details: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())
			return nil, newErr
		}
		return nil, err
	}
	res := strconv.Itoa(folder.ID)
	return &res, nil
}

// GetFolderByID retrieves a folder by its ID.
func (r folderRepository) GetFolderByID(ctx context.Context, id string) (*model.FolderDTO, error) {
	query := `SELECT id, owner_id, name, parent_id, created_at, updated_at FROM public.folder WHERE id = $1`
	var folder model.FolderDTO
	str := r.client.QueryRow(ctx, query, id)
	err := str.Scan(
		&folder.ID,
		&folder.OwnerID,
		&folder.Name,
		&folder.ParentID,
		&folder.CreatedAt,
		&folder.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, &modelerr.NotFound{ID: id}
		}

		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			newErr := fmt.Errorf("SQL Error: %s, Details: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())
			return nil, newErr
		}
		return nil, err
	}
	return &folder, nil
}

// GetFoldersByParentID retrieves all folders with a given parent ID.
func (r folderRepository) GetFoldersByParentID(ctx context.Context, FolderID string) ([]*model.FolderDTO, error) {
	q := `SELECT id, owner_id, name, parent_id, created_at, updated_at FROM public.folder WHERE parent_id = $1`
	rows, err := r.client.Query(ctx, q, FolderID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			newErr := fmt.Errorf("SQL Error: %s, Details: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())
			return nil, newErr
		}
		return nil, err
	}
	defer rows.Close()

	var folders []*model.FolderDTO
	for rows.Next() {
		var f model.FolderDTO
		if err := rows.Scan(&f.ID, &f.OwnerID, &f.Name, &f.ParentID, &f.CreatedAt, &f.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan folder: %w", err)
		}
		folders = append(folders, &f)
	}
	return folders, nil
}

// UpdateFolder updates a folder in the database.
func (r folderRepository) UpdateFolder(ctx context.Context, folder *model.FolderDTO) error {
	q := `UPDATE public.folder SET name = $1, parent_id = $2, updated_at = CURRENT_TIMESTAMP WHERE id = $3`
	if _, err := r.client.Exec(ctx, q, folder.Name, folder.ParentID, folder.ID); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			newErr := fmt.Errorf("SQL Error: %s, Details: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())
			return newErr
		}
		return err
	}
	return nil
}

// DeleteFolder deletes a folder from the database.

func (r folderRepository) DeleteFolder(ctx context.Context, id string) error {
	q := `DELETE FROM public.folder WHERE id = $1`
	if _, err := r.client.Exec(ctx, q, id); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			newErr := fmt.Errorf("SQL Error: %s, Details: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())
			return newErr
		}
		return err
	}
	return nil
}

// NewFolderRepo creates a new folder folderRepository.
func NewFolderRepo(client Client, logger log.Logger) model.FolderRepository {
	return folderRepository{
		client: client,
		log:    log.With(logger, "folderRepository", "folder"),
	}
}

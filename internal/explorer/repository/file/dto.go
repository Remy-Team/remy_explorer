package file

// DTO for File entity in the database.
import (
	"database/sql"
	"remy_explorer/pkg/domain/folder"
	"remy_explorer/pkg/domain/user"
	"time"
)

// FileDTOID is the type for the ID of the FileDTO.
type FileDTOID int64

// FileDTO is the data transfer object for the File entity in the database.
type FileDTO struct {
	ID         FileDTOID       `json:"id"`
	OwnerID    user.ID         `json:"owner_id"`
	Name       string          `json:"name"`
	FolderID   folder.FolderID `json:"folder_id"`
	ObjectPath sql.NullString  `json:"object_path"`
	Size       int             `json:"size"`
	Type       sql.NullString  `json:"type"`
	CreatedAt  time.Time       `json:"created_at"`
	UpdatedAt  time.Time       `json:"updated_at"`
	Tags       sql.NullString  `json:"tags"`
}

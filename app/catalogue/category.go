package catalogue

import (
	"time"

	"github.com/pborman/uuid"
)

type Category struct {
	ID          int
	Name        string
	ParentID    int
	CreatedDate time.Time
	UpdatedAT   *time.Time
	GuidID      uuid.UUID
}

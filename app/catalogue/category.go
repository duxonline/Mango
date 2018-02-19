package catalogue

import (
	"time"

	"github.com/pborman/uuid"
)

type Category struct {
	ID        int
	Name      string
	ParentID  int
	CreatedAt time.Time
	UpdatedAt time.Time
	GuidID    uuid.UUID
}

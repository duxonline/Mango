package catalogue

import (
	"time"

	"github.com/pborman/uuid"
)

type CurrentLease struct {
	ID           int
	CategoryName string
	ParentID     int
	CreatedDate  time.Time
	UpdatedTime  *time.Time
	GuidID       uuid.UUID
}

package domain

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID uuid.UUID
	Title string
	Description string
	Completed bool
    UserID uuid.UUID
    CreatedAt time.Time
    UpdatedAt time.Time
}

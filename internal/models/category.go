package models

import "time"

type Category struct {
	ID          uint64
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type BookCategory struct {
	BookID     uint64
	CategoryID uint64
}

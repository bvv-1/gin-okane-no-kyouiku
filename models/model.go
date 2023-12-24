package models

import "time"

type Model struct {
	ID        uint      `json:"-"`
	CreatedAt time.Time `json:"-"`
}

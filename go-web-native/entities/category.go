package entities

import "time"

type Category struct {
	Id        uint
	Nama      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

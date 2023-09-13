package entities

import "time"

type Product struct {
	Id        uint
	Nama      string
	Deskripsi string
	Kategori  Category
	Stock     int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

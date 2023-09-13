package productmodel

import (
	"go-web-native/config"
	"go-web-native/entities"
)

func GetAll() []entities.Product {
	rows, err := config.DB.Query(`
		SELECT p.id, p.nama, p.stock, k.nama as kategori 
		FROM produk p JOIN kategori k ON p.id_kategori = k.id
	`)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var products []entities.Product

	for rows.Next() {
		var produk entities.Product
		if err := rows.Scan(
			&produk.Id,
			&produk.Nama,
			&produk.Stock,
			&produk.Kategori.Nama,
		); err != nil {
			panic(err)
		}

		products = append(products, produk)
	}

	return products
}

func Create(product entities.Product) bool {

	var lastInsertId int64

	err := config.DB.QueryRow(`
		INSERT INTO produk(
			nama, stock, deskripsi, id_kategori, created_at, updated_at
		)VALUES(?, ?, ?, ?, ?, ?);
		SELECT ID = convert(bigint, SCOPE_IDENTITY())
		`,
		product.Nama,
		product.Stock,
		product.Deskripsi,
		product.Kategori.Id,
		product.CreatedAt,
		product.UpdatedAt,
	).Scan(&lastInsertId)

	if err != nil {
		panic(err)
	}

	return lastInsertId > 0

}

func Detail(id int) entities.Product {

	var product entities.Product

	err := config.DB.QueryRow(`
		SELECT p.id, p.nama, p.stock, k.nama as kategori, 
			p.deskripsi, p.created_at, p.updated_at 
		FROM produk p JOIN kategori k ON p.id_kategori = k.id
		WHERE p.id = ?
	`, id).Scan(
		&product.Id,
		&product.Nama,
		&product.Stock,
		&product.Kategori.Nama,
		&product.Deskripsi,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	if err != nil {
		panic(err)
	}

	return product

}

func Update(id int, product entities.Product) bool {
	query, err := config.DB.Exec(`UPDATE produk 
		SET 
		id_kategori = ?,
		nama = ?, 
		deskripsi = ?, 
		stock = ?, 
		updated_at = ? 
		WHERE id = ?`,
		product.Kategori.Id,
		product.Nama,
		product.Deskripsi,
		product.Stock,
		product.UpdatedAt,
		id,
	)

	if err != nil {
		panic(err)
	}

	results, err := query.RowsAffected()

	if err != nil {
		panic(err)
	}

	return results > 0

}

func Delete(id int) error {
	_, err := config.DB.Exec(`DELETE FROM produk WHERE id = ?`, id)
	return err
}

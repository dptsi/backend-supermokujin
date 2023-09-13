package categorymodel

import (
	"go-web-native/config"
	"go-web-native/entities"
)

func GetAll() []entities.Category {
	rows, err := config.DB.Query(`SELECT id, nama FROM dbo.kategori`)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var kategori []entities.Category

	for rows.Next() {
		var ikategori entities.Category
		if err := rows.Scan(&ikategori.Id, &ikategori.Nama); err != nil {
			panic(err)
		}

		kategori = append(kategori, ikategori)
	}

	return kategori

}

func Create(category entities.Category) bool {

	var lastInsertId int64

	err := config.DB.QueryRow(`
		INSERT INTO kategori(nama, created_at, updated_at)
		VALUES(?, ?, ?);
		SELECT ID = convert(bigint, SCOPE_IDENTITY())
		`,
		category.Nama, category.CreatedAt, category.UpdatedAt,
	).Scan(&lastInsertId)

	if err != nil {
		panic(err)
	}
	return lastInsertId > 0

}

func Detail(id int) entities.Category {
	row := config.DB.QueryRow(`SELECT id, nama FROM kategori where id = ?`, id)

	var category entities.Category
	if err := row.Scan(&category.Id, &category.Nama); err != nil {
		panic(err)
	}

	return category
}

func Update(id int, category entities.Category) bool {
	query, err := config.DB.Exec(`UPDATE kategori set nama = ?, updated_at = ? WHERE id = ?`, category.Nama, category.UpdatedAt, id)

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
	_, err := config.DB.Exec(`DELETE FROM kategori WHERE id = ?`, id)

	return err

}

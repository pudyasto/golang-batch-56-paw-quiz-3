package repository

import (
	"database/sql"
	"quiz-3/structs"
)

func GetAllCategory(db *sql.DB) (err error, results []structs.Category) {
	sql := "SELECT * FROM category"

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var category = structs.Category{}

		err = rows.Scan(&category.Id, &category.Name, &category.Created_at, &category.Updated_at)
		if err != nil {
			panic(err)
		}
		results = append(results, category)
	}

	return
}

func GetAllBookByCategories(db *sql.DB, id int) (err error, results []structs.Book) {
	sql := "SELECT * FROM book WHERE category_id = $1"

	rows, err := db.Query(sql, id)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var book = structs.Book{}

		err = rows.Scan(
			&book.Id,
			&book.Title,
			&book.Description,
			&book.Image_url,
			&book.Release_year,
			&book.Price,
			&book.Total_page,
			&book.Thickness,
			&book.Created_at,
			&book.Updated_at,
			&book.Category_id)
		if err != nil {
			panic(err)
		}
		results = append(results, book)
	}

	return
}

func InsertCategory(db *sql.DB, category structs.Category) (err error) {
	sql := "INSERT INTO category(name, created_at, updated_at) VALUES ($1, $2, $3)"
	errs := db.QueryRow(sql, category.Name, category.Created_at, category.Updated_at)
	return errs.Err()
}

func UpdateCategory(db *sql.DB, category structs.Category) (err error) {
	sql := "UPDATE public.category SET name=$1 , updated_at=$2 WHERE id=$3"
	errs := db.QueryRow(sql, category.Name, category.Updated_at, category.Id)
	return errs.Err()
}

func DeleteCategory(db *sql.DB, category structs.Category) (err error) {
	sql := "DELETE FROM public.category WHERE id=$1"
	errs := db.QueryRow(sql, category.Id)
	return errs.Err()
}

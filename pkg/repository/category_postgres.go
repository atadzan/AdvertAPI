package repository

import (
	"fmt"
	"github.com/atadzan/AdvertAPI"
	"github.com/jmoiron/sqlx"
)

type CategoryPostgres struct {
	db *sqlx.DB
}

func NewCategoryPostgres(db *sqlx.DB) *CategoryPostgres{
	return &CategoryPostgres{db: db}
}

func(r *CategoryPostgres) Add(category AdvertAPI.CategoryInput) (int, error){
	var id int
	query := fmt.Sprintf("INSERT INTO %s (title, main, parent_id) VALUES($1, $2, $3) RETURNING id", categoriesTable)
	row := r.db.QueryRow(query, category.Title, category.Main, category.ParentId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func(r *CategoryPostgres) GetMain()([]AdvertAPI.CategoryOutput, error){
	var mainCategories []AdvertAPI.CategoryOutput
	query := fmt.Sprintf("SELECT * FROM %s WHERE main=true", categoriesTable)
	row, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	for row.Next(){
		var category AdvertAPI.CategoryOutput
		err = row.Scan(&category.Id, &category.Title, &category.Main, &category.ParentId)
		mainCategories = append(mainCategories, category)
	}
	return mainCategories, nil
}

func(r *CategoryPostgres) GetNested(categoryId int)([]AdvertAPI.CategoryOutput, error){
	var nestedCategories []AdvertAPI.CategoryOutput
	query := fmt.Sprintf("SELECT * FROM %s WHERE parent_id=$1", categoriesTable)
	row, err := r.db.Query(query, categoryId)
	if err != nil {
		return nil, err
	}
	for row.Next(){
		var category AdvertAPI.CategoryOutput
		err = row.Scan(&category.Id, &category.Title, &category.Main, &category.ParentId)
		nestedCategories = append(nestedCategories, category)
	}
	return nestedCategories, nil
}

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

func(r *CategoryPostgres) GetCategoryAdverts(categoryId int)([]AdvertAPI.AdvertOutput, error){
	var categoryAdverts []AdvertAPI.AdvertOutput
	query := fmt.Sprintf("SELECT * FROM %s WHERE category_id=$1", advertsTable)
	row, err := r.db.Query(query, categoryId)
	if err != nil {
		return nil, err
	}
	if row != nil{
		for row.Next(){
			var advert AdvertAPI.AdvertOutput
			row.Scan(&advert.Id, &advert.Title, &advert.Description, &advert.Location, &advert.PhoneNumber,
				&advert.Price, &advert.PublishDate, &advert.Views, &advert.ImagesCount, &advert.UserId, &advert.CommentCount,
				&advert.Category)
			if advert.ImagesCount != 0 {
				imageQuery := fmt.Sprintf("SELECT * FROM %s WHERE advert_id = $1", advertImages)
				imageRow, err1 := r.db.Query(imageQuery, advert.Id)
				if err1 != nil {
					return nil, err
				}
				for imageRow.Next() {
					var image AdvertAPI.AdvertImage
					if imageRow != nil {
						if err = imageRow.Scan(&image.Id, &image.Fname, &image.Fsize, &image.Ftype, &image.Path, &image.AdvertId); err != nil {
							return nil, err
						}
					}
					var res AdvertAPI.ImageUrl
					url := fmt.Sprintf("http://192.168.1.181:8080/api/advert/image/%d", image.Id)
					res.URL = url
					advert.Images = append(advert.Images, res)
				}
			}
			if advert.CommentCount != 0 {
				commentQuery := fmt.Sprintf("SELECT * from %s WHERE advert_id=$1", commentsTable)
				rows, err2 := r.db.Query(commentQuery, advert.Id)
				if err2 != nil {
					return nil, err2
				}
				for rows.Next() {
					var comment AdvertAPI.Comment
					if err = rows.Scan(&comment.Id, &comment.AdvertId, &comment.Body, &comment.UserId, &comment.CreatedAt, &comment.UpdatedAt); err != nil {
						return nil, err
					}
					advert.Comments = append(advert.Comments, comment)
				}
			}
			categoryAdverts = append(categoryAdverts, advert)
		}
	}
	return categoryAdverts, nil
}


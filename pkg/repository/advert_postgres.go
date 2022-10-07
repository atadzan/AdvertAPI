package repository

import (
	"fmt"
	"github.com/atadzan/AdvertAPI"
	"github.com/jmoiron/sqlx"
	"time"
)

type AdvertPostgres struct {
	db         *sqlx.DB
}

func NewAdvertPostgres(db *sqlx.DB) *AdvertPostgres{
	return &AdvertPostgres{db: db}
}

func(r *AdvertPostgres) Add(advert AdvertAPI.AdvertInput)(int, error){
	tx, err := r.db.Begin()
	if err != nil {
		return 0, nil
	}
	var id int
	createAdvertQuery := fmt.Sprintf("INSERT INTO %s (title, description, category, location, phone_number," +
		" price, publish_date) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id", advertsTable)
	row := tx.QueryRow(createAdvertQuery, advert.Title, advert.Description, advert.Category, advert.Location,
		advert.PhoneNumber, advert.Price, time.Now())
	if err := row.Scan(&id); err != nil{
		err = tx.Rollback()
		return 0, err
	}
	count := 0
	if len(advert.Images) != 0{
		for _, path := range advert.Images{
			createAdvertImages := fmt.Sprintf("INSERT INTO %s (fname, fsize, ftype, path, advert_id) VALUES($1, $2, $3, $4, $5)", advertImages)
			_, err := tx.Exec(createAdvertImages, path.Fname, path.Fsize, path.Ftype, path.Path, id )
			if err != nil {
				err = tx.Rollback()
				return 0, err
			}
			count += 1
			addCount := fmt.Sprintf("UPDATE %s SET image_count = $1 WHERE id = $2", advertsTable)
			_, err = tx.Exec(addCount, count, id)
			if err != nil {
				tx.Rollback()
				return 0, err
			}
		}
	}
	return id, tx.Commit()
}

func(r *AdvertPostgres) GetAll(advertPerPage, offset int)([]AdvertAPI.AdvertInfo, error){
	var adverts []AdvertAPI.AdvertInfo
	query := fmt.Sprintf("SELECT * FROM %s ORDER BY publish_date DESC LIMIT $1 OFFSET $2", advertsTable)
	row, err := r.db.Query(query, advertPerPage, offset)
	if err != nil {
		return nil, err
	}
	for row.Next(){
		var advert AdvertAPI.AdvertInfo
		if err := row.Scan(&advert.Id, &advert.Title, &advert.Description, &advert.Category, &advert.Location,
			&advert.PhoneNumber, &advert.Price, &advert.PublishDate, &advert.Views, &advert.ImagesCount); err != nil{
			return adverts, err
		}
		if advert.ImagesCount != 0 {
			imageQuery := fmt.Sprintf("SELECT * FROM %s WHERE advert_id = $1", advertImages)
			imageRow, err := r.db.Query(imageQuery, advert.Id)
			if err != nil {
				return nil, err
			}

			for imageRow.Next(){
				var image AdvertAPI.AdvertImage
				if imageRow != nil {
					if err := imageRow.Scan(&image.Id, &image.Fname, &image.Fsize, &image.Ftype, &image.Path, &image.AdvertId); err != nil {
						return nil, err
					}
				}
				advert.Images = append(advert.Images, image)
			}
		}
		adverts = append(adverts, advert)
	}
	return adverts, err
}

func(r *AdvertPostgres) GetById(id int)(AdvertAPI.AdvertInfo, error){
	var advert AdvertAPI.AdvertInfo
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", advertsTable)
	row := r.db.QueryRow(query, id)
	err := row.Scan(&advert.Id, &advert.Title, &advert.Description, &advert.Category, &advert.Location,
		&advert.PhoneNumber, &advert.Price, &advert.PublishDate, &advert.Views, &advert.ImagesCount)
	if advert.ImagesCount != 0 {
		imageQuery := fmt.Sprintf("SELECT * FROM %s WHERE advert_id = $1", advertImages)
		imageRow, err := r.db.Query(imageQuery, advert.Id)
		if err != nil {
			return advert, err
		}
		for imageRow.Next(){
			var image AdvertAPI.AdvertImage
			if imageRow != nil {
				if err := imageRow.Scan(&image.Id, &image.Fname, &image.Fsize, &image.Ftype, &image.Path, &image.AdvertId); err != nil {
					return advert, err
				}
			}
			advert.Images = append(advert.Images, image)
		}
	}
	return advert, err
}

func(r *AdvertPostgres) CountAdverts()(int, error){
	var count int
	query := fmt.Sprintf("SELECT COUNT(id) AS count FROM %s", advertsTable)
	row := r.db.QueryRow(query)
	err := row.Scan(&count)
	return count, err
}

func(r *AdvertPostgres)AddDB(file AdvertAPI.AdvertImage)(string, error){
	query:= fmt.Sprintf("INSERT INTO %s (fname, fsize, ftype, path, advert_id) VALUES($1, $2, $3, $4, $5)", advertImages)
	_, err := r.db.Exec(query, file.Fname, file.Fsize, file.Ftype, file.Path, file.AdvertId)
	if err != nil {
		return err.Error(), nil
	}
	return "Success", nil
}

func(r *AdvertPostgres) GetImage(id int)([]AdvertAPI.AdvertImage, error){
	var image AdvertAPI.AdvertImage
	var images []AdvertAPI.AdvertImage
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", advertImages)
	row, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}

	for row.Next(){
		err = row.Scan(&image.Id, &image.Fname, &image.Fsize, &image.Ftype, &image.Path, &image.AdvertId)
		if err != nil {
			return nil, err
		}
		images = append(images, image)
		if len(images) >= 1 {
			return images, nil
		}else{
			return nil, err
		}
	}
	return images, nil
}

//func(r *AdvertPostgres) GetPhoto(id int)([]AdvertAPI.AdvertImage, error){
//	var image AdvertAPI.AdvertImage
//	var images []AdvertAPI.AdvertImage
//	query := fmt.Sprintf("SELECT * FROM %s WHERE advert_id = $1", advertImages)
//	row, err := r.db.Query(query, id)
//	if err != nil {
//		return nil, err
//	}
//
//	for row.Next(){
//		err = row.Scan(&image.Id, &image.Fname, &image.Fsize, &image.Ftype, &image.Path, &image.AdvertId)
//		if err != nil {
//			return nil, err
//		}
//		images = append(images, image)
//		if len(images) >= 1 {
//			return images, nil
//		}else{
//			return nil, err
//		}
//	}
//	return images, nil
//}


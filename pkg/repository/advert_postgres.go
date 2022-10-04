package repository

import (
	"fmt"
	"github.com/atadzan/AdvertAPI"
	"github.com/jmoiron/sqlx"
	"time"
)

type AdvertPostgres struct {
	db *sqlx.DB
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
	if len(advert.Images) != 0{
		for _, path := range advert.Images{
			createAdvertImages := fmt.Sprintf("INSERT INTO %s (path, advert_id) VALUES($1, $2)", advertImages)
			_, err := tx.Exec(createAdvertImages, path.ImagePath, id )
			if err != nil {
				err = tx.Rollback()
				return 0, err
			}
		}
	}
	return id, tx.Commit()
}

func(r *AdvertPostgres) GetAll()([]AdvertAPI.AdvertInfo, error){
	var adverts []AdvertAPI.AdvertInfo
	rows, err := r.db.Query("SELECT * FROM adverts")
	if err != nil {
		return nil, err
	}
	for rows.Next(){
		var advert AdvertAPI.AdvertInfo
		if err := rows.Scan(&advert.Id, &advert.Title, &advert.Description, &advert.Category, &advert.Location,
			&advert.PhoneNumber, &advert.Price, &advert.PublishDate, &advert.Views); err != nil{
			return adverts, err
		}
		//imageRow, err := r.db.Query("SELECT path FROM images WHERE advert_id = ?", advert.Id)
		//if err != nil {
		//	return nil, err
		//}
		//for imageRow.Next(){
		//	var advertImage AdvertAPI.Image
		//	if err := imageRow.Scan(&advertImage); err != nil{
		//		return adverts, err
		//	}
		//	advert = append(advert.Images, advertImage)
		//}
		adverts = append(adverts, advert)
	}
	if err = rows.Err(); err != nil {
		return adverts, err
	}

	//fmt.Println(rows)
	//defer rows.Close()
	//for rows.Next()


	return adverts, err
}
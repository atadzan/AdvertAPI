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
			&advert.PhoneNumber, &advert.Price, &advert.PublishDate, &advert.Views); err != nil{
			return adverts, err
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
		&advert.PhoneNumber, &advert.Price, &advert.PublishDate, &advert.Views)
	return advert, err
}

func(r *AdvertPostgres) CountAdverts()(int, error){
	var count int
	query := fmt.Sprintf("SELECT COUNT(id) AS count FROM %s", advertsTable)
	row := r.db.QueryRow(query)
	err := row.Scan(&count)
	return count, err
}

func(r *AdvertPostgres)AddDB(fname, ftype, filepath string, fsize int64)(string, error){
	query:= fmt.Sprintf("INSERT INTO %s (fname, fsize, ftype, path) VALUES($1, $2, $3, $4)", advertImages)
	_, err := r.db.Exec(query, fname, fsize, ftype, filepath)
	if err != nil {
		return err.Error(), nil
	}
	return "Success", nil
}

func(r *AdvertPostgres) GetImage(id int)([]AdvertAPI.AdvertImage, error){
	query := fmt.Sprintf(" SELECT * FROM %s WHERE id=$1", advertImages)
	row, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	upld := AdvertAPI.AdvertImage{}
	var res []AdvertAPI.AdvertImage
	for row.Next(){
		err = row.Scan(&upld.Id, &upld.Fname, &upld.Fsize, &upld.Ftype, &upld.Path, &upld.AdvertId)
		if err != nil {
			return nil, err
		}
		res = append(res, upld)
		if len(res) >= 1 {
			return res, nil
		}else{
			return nil, err
		}
	}
	return res, nil
}
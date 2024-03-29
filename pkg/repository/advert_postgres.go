package repository

import (
	"database/sql"
	"fmt"
	"github.com/atadzan/AdvertAPI"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"strconv"
	"strings"
	"time"
)

type AdvertPostgres struct {
	db *sqlx.DB
}

func NewAdvertPostgres(db *sqlx.DB) *AdvertPostgres {
	return &AdvertPostgres{db: db}
}

func (r *AdvertPostgres) Add(advert AdvertAPI.AdvertInput) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, nil
	}
	var id int
	createAdvertQuery := fmt.Sprintf("INSERT INTO %s (title, description, category_id, location, phone_number,"+
		" price, publish_date, user_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id", advertsTable)
	row := tx.QueryRow(createAdvertQuery, advert.Title, advert.Description, advert.Category, advert.Location,
		advert.PhoneNumber, advert.Price, time.Now(), advert.UserId)
	if ok := row.Scan(&id); ok != nil {
		ok = tx.Rollback()
		return 0, err
	}
	count := 0
	if len(advert.Images) != 0 {
		for _, path := range advert.Images {
			createAdvertImages := fmt.Sprintf("INSERT INTO %s (fname, fsize, ftype, path, advert_id) VALUES($1, $2, $3, $4, $5)", advertImages)
			_, ok := tx.Exec(createAdvertImages, path.Fname, path.Fsize, path.Ftype, path.Path, id)
			if ok != nil {
				ok = tx.Rollback()
				return 0, ok
			}
			count += 1
			addCount := fmt.Sprintf("UPDATE %s SET image_count = $1 WHERE id = $2", advertsTable)
			_, ok = tx.Exec(addCount, count, id)
			if ok != nil {
				ok = tx.Rollback()
				return 0, ok
			}
		}
	}
	return id, tx.Commit()
}

func (r *AdvertPostgres) GetAll(advertPerPage, offset int) ([]AdvertAPI.AdvertOutput, error) {
	var adverts []AdvertAPI.AdvertOutput
	query := fmt.Sprintf("SELECT * FROM %s ORDER BY publish_date DESC LIMIT $1 OFFSET $2", advertsTable)
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	row, ok := r.db.Query(query, advertPerPage, offset)
	if ok != nil {
		return nil, ok
	}
	for row.Next() {
		advert, err1  := r.getAdvert(row, tx)
		if err1 != nil{
			return nil, err1
		}
		adverts = append(adverts, advert)
	}
	return adverts, tx.Commit()
}

func (r *AdvertPostgres) GetById(id int) (AdvertAPI.AdvertOutput, error) {
	var advert AdvertAPI.AdvertOutput
	tx, err := r.db.Begin()
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", advertsTable)
	row := tx.QueryRow(query, id)
	err = row.Scan(&advert.Id, &advert.Title, &advert.Description, &advert.Location,&advert.PhoneNumber,&advert.Price,
		&advert.PublishDate, &advert.Views, &advert.ImagesCount, &advert.UserId,&advert.CommentCount, &advert.Category)
	if err != nil {
		return advert, err
	}
	advert.Views += 1
	updateQuery := fmt.Sprintf("UPDATE %s SET views = $1 WHERE id = $2", advertsTable)
	_, err = tx.Exec(updateQuery, advert.Views, id)
	if err != nil {
		err = tx.Rollback()
		return advert, err
	}
	if advert.ImagesCount != 0 {
		advert, err = r.getAdvertImages(advert, tx)
	}
	if advert.CommentCount != 0 {
		advert, err = r.getAdvertComments(advert, tx)
	}
	return advert, tx.Commit()
}

func (r *AdvertPostgres) CountAdverts() (int, error) {
	var count int
	query := fmt.Sprintf("SELECT COUNT(id) AS count FROM %s", advertsTable)
	row := r.db.QueryRow(query)
	err := row.Scan(&count)
	return count, err
}

func (r *AdvertPostgres) GetImage(id int) ([]AdvertAPI.AdvertImage, error) {
	var image AdvertAPI.AdvertImage
	var images []AdvertAPI.AdvertImage
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", advertImages)
	row, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		err = row.Scan(&image.Id, &image.Fname, &image.Fsize, &image.Ftype, &image.Path, &image.AdvertId)
		if err != nil {
			return nil, err
		}
		images = append(images, image)
		if len(images) >= 1 {
			return images, nil
		} else {
			return nil, err
		}
	}
	return images, nil
}

func (r *AdvertPostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id= $1", advertsTable)
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return err
}

func (r *AdvertPostgres) Update(id int, advert AdvertAPI.AdvertInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if advert.Title != "" {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, advert.Title)
		argId++
	}
	if advert.Description != "" {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, advert.Description)
		argId++
	}
	if advert.Category != 0 {
		setValues = append(setValues, fmt.Sprintf("category=$%d", argId))
		args = append(args, advert.Category)
		argId++
	}
	if advert.Location != "" {
		setValues = append(setValues, fmt.Sprintf("location=$%d", argId))
		args = append(args, advert.Location)
		argId++
	}
	if advert.PhoneNumber != "" {
		setValues = append(setValues, fmt.Sprintf("phone_number=$%d", argId))
		args = append(args, advert.PhoneNumber)
		argId++
	}
	if advert.Price != 0 {
		setValues = append(setValues, fmt.Sprintf("price=$%d", argId))
		args = append(args, advert.Price)
		argId++
	}
	if len(advert.Images) != 0 {
		for _, path := range advert.Images {
			updateAdvertImages := fmt.Sprintf("UPDATE %s  SET fname = $1,  fsize = $2, ftype = $3, path = $4, advert_id = $5 WHERE advert_id = $6", advertImages)
			_, err := r.db.Exec(updateAdvertImages, path.Fname, path.Fsize, path.Ftype, path.Path, id, id)
			if err != nil {
				return err
			}
		}
	}
	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s,  publish_date = $%d, user_id=$%d WHERE id = $%d", advertsTable, setQuery, argId, argId+1, argId+2)
	args = append(args, time.Now(), advert.UserId, id)
	_, err := r.db.Exec(query, args...)
	return err
}

func (r *AdvertPostgres) AddFav(userId, advertId int) error {
	id := strconv.Itoa(advertId)
	query := fmt.Sprintf("UPDATE %s SET fav_list = array_append(fav_list, $1) WHERE id = $2", usersTable)
	_, err := r.db.Exec(query, id, userId)
	if err != nil {
		return err
	}
	return nil
}

func (r *AdvertPostgres) GetFav(userId int) ([]AdvertAPI.AdvertOutput, error) {
	var adverts []AdvertAPI.AdvertOutput
	var result pq.Int64Array
	favQuery := fmt.Sprintf("SELECT fav_list[1:] FROM %s WHERE id=$1", usersTable)
	favlist := r.db.QueryRow(favQuery, userId)
	if err := favlist.Scan(&result); err != nil {
		fmt.Println("Error in favlist")
		return nil, err
	}
	favIds := []int64(result)
	tx, err := r.db.Begin()
	for _, id := range favIds {
		var advert AdvertAPI.AdvertOutput
		query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", advertsTable)

		if err != nil {
			return nil, err
		}
		row := tx.QueryRow(query, id)
		if err = row.Scan(&advert.Id, &advert.Title, &advert.Description, &advert.Location, &advert.PhoneNumber,
			&advert.Price, &advert.PublishDate, &advert.Views, &advert.ImagesCount, &advert.UserId, &advert.CommentCount,
			&advert.Category); err != nil {
			return adverts, err
		}
		if advert.ImagesCount != 0 {
			advert, err = r.getAdvertImages(advert, tx)
			if err != nil {
				return nil, err
			}
		}
		if advert.CommentCount != 0 {
			advert, err = r.getAdvertComments(advert, tx)
			if err != nil {
				return nil, err
			}
		}
		adverts = append(adverts, advert)
	}
	return adverts, tx.Commit()
}

func (r *AdvertPostgres) DeleteFav(userId, advertId int) error {
	var result1 pq.Int64Array
	id := int64(advertId)
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	favQuery := fmt.Sprintf("SELECT fav_list FROM %s WHERE id=$1", usersTable)
	favlist := tx.QueryRow(favQuery, userId)
	if err = favlist.Scan(&result1); err != nil {
		 err = tx.Rollback()
		return err
	}
	intRes := []int64(result1)
	for i, k := range intRes {
		if id == k {
			intRes[i] = intRes[len(intRes)-1]
			intRes[len(intRes)-1] = 0
			intRes = intRes[:len(intRes)-1]
		}
	}
	inputRes := fmt.Sprintf("UPDATE %s SET fav_list = $1 WHERE id=$2", usersTable)
	_, err = tx.Exec(inputRes, pq.Array(intRes), userId)
	if err != nil {
		err = tx.Rollback()
		return err
	}
	return tx.Commit()
}

func(r *AdvertPostgres) CheckFavList(userId, advertId int)(bool, error){
	var response bool
	var result pq.Int64Array
	favQuery := fmt.Sprintf("SELECT fav_list FROM %s WHERE id=$1", usersTable)
	favlist := r.db.QueryRow(favQuery, userId)
	if err := favlist.Scan(&result); err != nil {
		fmt.Println("Error in favlist")
		return response, err
	}
	favIds := []int64(result)
	id := int64(advertId)
	for _, v := range favIds {
		if  id == v  {
			response = true
		}else{
			response = false
		}
	}
	return response, nil
}

func (r *AdvertPostgres) Search(search string) ([]AdvertAPI.AdvertOutput, error) {
	var adverts []AdvertAPI.AdvertOutput
	query := fmt.Sprintf("SELECT * FROM %s AS a WHERE a.title LIKE $1 ORDER BY a.publish_date DESC", advertsTable)
	row, err := r.db.Query(query, "%"+search+"%")
	if err != nil {
		return adverts, err
	}
	tx, err := r.db.Begin()
	for row.Next() {
		advert, ok := r.getAdvert(row, tx)
		if ok != nil {
			return nil, ok
		}
		adverts = append(adverts, advert)
	}
	if adverts == nil{
		emptyAdverts := []AdvertAPI.AdvertOutput{}
		adverts = emptyAdverts
	}
	return adverts, tx.Commit()
}

func (r *AdvertPostgres) getAdvert(row *sql.Rows, tx *sql.Tx)(AdvertAPI.AdvertOutput, error){
	var advert AdvertAPI.AdvertOutput
	var err error
	if err = row.Scan(&advert.Id, &advert.Title, &advert.Description, &advert.Location,
		&advert.PhoneNumber, &advert.Price, &advert.PublishDate, &advert.Views, &advert.ImagesCount, &advert.UserId, &advert.CommentCount, &advert.Category); err != nil {
		return advert, err
	}
	if advert.ImagesCount != 0 {
		advert, err = r.getAdvertImages(advert, tx)
		if err !=nil {
			fmt.Println("Internal server error")
		}
	}
	if advert.CommentCount != 0 {
		advert, err = r.getAdvertComments(advert, tx)
	}
	return advert, nil
}

func(r *AdvertPostgres) getAdvertImages(advert AdvertAPI.AdvertOutput, tx *sql.Tx) (AdvertAPI.AdvertOutput, error) {
	imageQuery := fmt.Sprintf("SELECT * FROM %s WHERE advert_id = $1", advertImages)
	imageRow, ok := tx.Query(imageQuery, advert.Id)
	if ok != nil {
		ok = tx.Rollback()
		return advert, ok
	}
	for imageRow.Next() {
		var image AdvertAPI.AdvertImage
		if imageRow != nil {
			if err := imageRow.Scan(&image.Id, &image.Fname, &image.Fsize, &image.Ftype, &image.Path, &image.AdvertId); err != nil {
				return advert, err
			}
		}
		var res AdvertAPI.ImageUrl
		url := fmt.Sprintf("http://192.168.1.181:8080/api/advert/image/%d", image.Id)
		res.URL = url
		advert.Images = append(advert.Images, res)
	}
	return advert, nil
}

func(r *AdvertPostgres) getAdvertComments(advert AdvertAPI.AdvertOutput, tx *sql.Tx)(AdvertAPI.AdvertOutput, error){
	commentQuery := fmt.Sprintf("SELECT * from %s WHERE advert_id=$1", commentsTable)
	rows, er := tx.Query(commentQuery, advert.Id)
	if er != nil {
		er = tx.Rollback()
		fmt.Println("No comments")
	}
	for rows.Next() {
		var comment AdvertAPI.Comment
		if err := rows.Scan(&comment.Id, &comment.AdvertId, &comment.Body, &comment.UserId, &comment.CreatedAt, &comment.UpdatedAt); err != nil {
			return advert, err
		}
		advert.Comments = append(advert.Comments, comment)
	}
	return advert, nil
}
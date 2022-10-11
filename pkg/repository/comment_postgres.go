package repository

import (
	"fmt"
	"github.com/atadzan/AdvertAPI"
	"github.com/jmoiron/sqlx"
	"time"
)

type CommentPostgres struct {
	db *sqlx.DB
}

func NewCommentPostgres(db *sqlx.DB) *CommentPostgres{
	return &CommentPostgres{db: db}
}

func(r *CommentPostgres)AddCom(comment AdvertAPI.InputComm, userId, advertId int) error{
	query := fmt.Sprintf("INSERT INTO %s(advert_id ,body, user_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)", commentsTable)
	_, err := r.db.Exec(query, advertId, comment.Body, userId, time.Now(), time.Now())
	if err != nil {
		return err
	}
	return nil
}

func(r *CommentPostgres)GetCom(advertId, userId int)([]AdvertAPI.Comment, error){
	var comments []AdvertAPI.Comment
	query := fmt.Sprintf("SELECT * FROM %s WHERE advert_id=$1 AND user_id=$2", commentsTable)
	rows, err := r.db.Query(query, advertId, userId)
	if err != nil {
		return nil, err
	}
	for rows.Next(){
		var comment AdvertAPI.Comment
		if err = rows.Scan(&comment.Id, &comment.AdvertId, &comment.Body, &comment.UserId, &comment.CreatedAt,
			&comment.UpdatedAt); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, err
}

func(r *CommentPostgres) DelCom(advertId, userId, commentId int) error{
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1 AND advert_id=$2 AND user_id=$3", commentsTable)
	_, err := r.db.Exec(query, commentId, advertId, userId)
	if err != nil {
		return err
	}
	return nil
}

func(r *CommentPostgres) UpdCom(comment AdvertAPI.InputComm, userId, advertId, commentId int) error{
	query := fmt.Sprintf("UPDATE %s SET body=$1, updated_at=$2", commentsTable)
	_, err := r.db.Exec(query, comment.Body, time.Now())
	if err != nil {
		return err
	}
	return nil
}
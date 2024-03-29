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
	tx, err := r.db.Begin()
	_, err = tx.Exec(query, advertId, comment.Body, userId, time.Now(), time.Now())
	if err != nil {
		return err
	}
	addCount := fmt.Sprintf("UPDATE %s SET comment_count=comment_count+1 WHERE id=$1", advertsTable)
	_, err = tx.Exec(addCount, advertId)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func(r *CommentPostgres) DelCom(advertId, userId, commentId int) error{
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1 AND advert_id=$2 AND user_id=$3", commentsTable)
	tx, err := r.db.Begin()
	_, err = tx.Exec(query, commentId, advertId, userId)
	if err != nil {
		return err
	}
	reduceCount := fmt.Sprintf("UPDATE %s SET comment_count=comment_count-1 WHERE id=$1",advertsTable)
	_, err = tx.Exec(reduceCount, advertId)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func(r *CommentPostgres) UpdCom(comment AdvertAPI.InputComm, userId, advertId, commentId int) error{
	query := fmt.Sprintf("UPDATE %s SET body=$1, updated_at=$2 WHERE id=$3 AND advert_id=$4 AND user_id=$5", commentsTable)
	_, err := r.db.Exec(query, comment.Body, time.Now(), commentId, advertId, userId)
	if err != nil {
		return err
	}
	return nil
}
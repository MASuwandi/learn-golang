package repository

import (
	"context"
	"database/sql"
	"errors"
	"go-db-mysequel/server/entity"
	"strconv"
)

type commentRepositoryImplement struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImplement{DB: db}
}

func (repository *commentRepositoryImplement) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	query := `INSERT INTO comments(email, comment)
			  VALUES (?, ?)`
	result, err := repository.DB.ExecContext(ctx, query, comment.Email, comment.Comment)
	if err != nil {
		return comment, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}
	comment.Id = int32(id)
	return comment, err
}

func (repository *commentRepositoryImplement) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	query := `SELECT id, email, comment
			  FROM comments 
			  WHERE id = ?
			  LIMIT 1`
	rows, err := repository.DB.QueryContext(ctx, query, id)
	comment := entity.Comment{}
	if err != nil {
		return comment, err
	}
	defer rows.Close()

	if rows.Next() { // data exist
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil
	}

	// data not exist
	return comment, errors.New("Id: " + strconv.Itoa(int(id)) + " Not Found")
}

func (repository *commentRepositoryImplement) FindAll(ctx context.Context) ([]entity.Comment, error) {
	query := `SELECT id, email, comment
			  FROM comments`
	rows, err := repository.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []entity.Comment
	for rows.Next() { // loop existing data
		comment := entity.Comment{}
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		comments = append(comments, comment)
	}

	return comments, nil
}

package repository

import (
	"context"
	"fmt"
	go_db_mysequel "go-db-mysequel"
	"go-db-mysequel/server/entity"
	"testing"

	_ "github.com/go-sql-driver/mysql" // init
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(go_db_mysequel.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "repotest4@mail.com",
		Comment: "Test Repository 4",
	}

	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println("result: ", result)
}

func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepository(go_db_mysequel.GetConnection())

	ctx := context.Background()

	comment, err := commentRepository.FindById(ctx, 34)
	if err != nil {
		panic(err)
	}

	fmt.Println("comment: ", comment)
}

func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(go_db_mysequel.GetConnection())

	ctx := context.Background()

	comments, err := commentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	for _, comment := range comments {
		fmt.Println("comment: ", comment)
	}
}

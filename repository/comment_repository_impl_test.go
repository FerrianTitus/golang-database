package repository

import (
	"context"
	"fmt"
	golang_database "github.com/ferriantitus/golang-database"
	"github.com/ferriantitus/golang-database/entity"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(golang_database.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "repository@test.com",
		Comment: "Test Repository",
	}

	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestById(t *testing.T) {
	commentRepository := NewCommentRepository(golang_database.GetConnection())

	comment, err := commentRepository.FindById(context.Background(), 35)
	if err != nil {
		panic(err)
	}
	fmt.Println(comment)
}

func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(golang_database.GetConnection())

	comments, err := commentRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}
	for _, comment := range comments {
		fmt.Println(comment)
	}
}

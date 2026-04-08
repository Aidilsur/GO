package repository

import (
	"context"
	"fmt"
	"go_database"
	"go_database/entity"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	commenRepository := NewCommentRepository(go_database.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email: "repository@test.com",
		Comment: "Test Repository",
	}

	result, err := commenRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindByid(t *testing.T) {
	commenRepository := NewCommentRepository(go_database.GetConnection())

	comment, err := commenRepository.FindById(context.Background(), 37)
	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
}

func TestFindByid(t *testing.T) {
	commenRepository := NewCommentRepository(go_database.GetConnection())

	comment, err := commenRepository.FindById(context.Background(), 37)
	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
}
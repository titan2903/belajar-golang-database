package repository

import (
	belajargolangdatabase "belajar-golang-database"
	"belajar-golang-database/entity"
	"context"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)


func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(belajargolangdatabase.GetConnection())
	ctx := context.Background()

	comment := entity.Comment {
		Email: "repository@test.com",
		Comment: "test2 repository",
	}

	result , err := commentRepository.Insert(ctx, comment)

	if err != nil {
		panic(err)
	}

	fmt.Println("Insert result: ", result)
}

func TestCommentFindById(t *testing.T) {
	commentRepository := NewCommentRepository(belajargolangdatabase.GetConnection())

	comment, err := commentRepository.FindById(context.Background(), 90)

	if err != nil {
		panic(err)
	}

	fmt.Println("Find by Id Comment: ", comment)
}

func TestCommentFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(belajargolangdatabase.GetConnection())

	comment, err := commentRepository.FIndAll(context.Background())

	if err != nil {
		panic(err)
	}

	for _, comment := range comment {
		fmt.Println("Find All Comment: ", comment)
	}

}
package repository

import (
	"belajar-golang-database/entity"
	"context"
)


type CommentRepository interface {
	Insert(ctx context.Context, comment entity.Comment)(entity.Comment, error)
	FindById(ctx context.Context, id int32) (entity.Comment, error)
	FIndAll(ctx context.Context) ([]entity.Comment, error)
}

/*
! biasnaya nama repository seperti nama dari entity nya
*/
package service

import (
	"crypto/rand"
	"errors"

	"github.com/xiemenger/GolangMock/entity"
	"../repository"
)

type PostService interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type service struct{}

var (
	repo repository.PostRepository = repository.NewPostRepository()
)

func NewPostService() PostService {
	return &service
}

func (*service) Validate(post *entity.Post) error {
	if post == nil {
		err := errors.New("The arg post is empty.")
		return err
	}

	if post == nil {
		err := errors.New("The arg post is empty.")
		return err
	}
	return nil
}

func (*service) Create(post *entity.Post) (*entity.Post, error) {
	postID := rand.Int()
	return repo.()
}

func (*service) FindAll() ([]entity.Post, error) {

}

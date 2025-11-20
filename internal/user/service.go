package user

import (
	"context"
)

type Service interface {
	CreateUser(ctx context.Context, u *User) error
	GetUser(ctx context.Context, id uint) (*User, error)
}

type service struct {
	repo Repository
}

func NewUserService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateUser(ctx context.Context, u *User) error {
	return s.repo.Create(ctx, u)
}

func (s *service) GetUser(ctx context.Context, id uint) (*User, error) {
	return s.repo.FindByID(ctx, id)
}

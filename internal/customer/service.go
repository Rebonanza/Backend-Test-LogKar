package customer

import (
    "context"
)

type Service interface {
    Create(ctx context.Context, c *Customer) error
    Get(ctx context.Context, id uint) (*Customer, error)
}

type service struct{ repo Repository }

func NewCustomerService(repo Repository) Service { return &service{repo: repo} }

func (s *service) Create(ctx context.Context, c *Customer) error { return s.repo.Create(ctx, c) }

func (s *service) Get(ctx context.Context, id uint) (*Customer, error) { return s.repo.FindByID(ctx, id) }

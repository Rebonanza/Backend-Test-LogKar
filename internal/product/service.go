package product

import (
    "context"
)

type Service interface {
    Create(ctx context.Context, p *Product) error
    Get(ctx context.Context, id uint) (*Product, error)
    List(ctx context.Context) ([]Product, error)
}

type service struct {
    repo Repository
}

func NewProductService(repo Repository) Service {
    return &service{repo: repo}
}

func (s *service) Create(ctx context.Context, p *Product) error {
    return s.repo.Create(ctx, p)
}

func (s *service) Get(ctx context.Context, id uint) (*Product, error) {
    return s.repo.FindByID(ctx, id)
}

func (s *service) List(ctx context.Context) ([]Product, error) {
    return s.repo.List(ctx)
}

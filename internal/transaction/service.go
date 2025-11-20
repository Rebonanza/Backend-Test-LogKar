package transaction

import (
    "context"
    "errors"
    "time"

    "github.com/google/uuid"
    "github.com/local/be-test-logkar/internal/customer"
    "github.com/local/be-test-logkar/internal/product"
)

var (
    ErrInsufficientPoints = errors.New("insufficient points")
    ErrOutOfStock         = errors.New("product out of stock")
)

var RedeemCost = map[string]int{
    "Small":  200,
    "Medium": 300,
    "Large":  500,
}

type Service interface {
    CreatePurchase(ctx context.Context, customerID uint, productID uint, size, flavor string, qty int, unitPrice int) (*Transaction, error)
    RedeemBySize(ctx context.Context, customerID uint, size string) (*Transaction, error)
}

type service struct {
    repo Repository
    prodRepo product.Repository
    custRepo customer.Repository
}

func NewTransactionService(repo Repository, pr product.Repository, cr customer.Repository) Service {
    return &service{repo: repo, prodRepo: pr, custRepo: cr}
}

func (s *service) CreatePurchase(ctx context.Context, customerID uint, productID uint, size, flavor string, qty int, unitPrice int) (*Transaction, error) {
    if err := s.prodRepo.DecreaseQuantity(ctx, productID, qty); err != nil {
        return nil, ErrOutOfStock
    }

    total := unitPrice * qty
    points := total / 1000

    if points > 0 {
        if err := s.custRepo.UpdatePoints(ctx, customerID, points); err != nil {
            return nil, err
        }
    }

    t := &Transaction{
        ID:         uuid.NewString(),
        CustomerID: customerID,
        ProductID:  productID,
        Size:       size,
        Flavor:     flavor,
        Quantity:   qty,
        CreatedAt:  time.Now().UTC(),
    }
    if err := s.repo.Create(ctx, t); err != nil {
        return nil, err
    }
    return t, nil
}

func (s *service) RedeemBySize(ctx context.Context, customerID uint, size string) (*Transaction, error) {
    cost, ok := RedeemCost[size]
    if !ok {
        return nil, errors.New("invalid size")
    }

    cu, err := s.custRepo.FindByID(ctx, customerID)
    if err != nil {
        return nil, err
    }
    if cu.Points < cost {
        return nil, ErrInsufficientPoints
    }

    p, err := s.prodRepo.FindAvailableBySize(ctx, size)
    if err != nil {
        return nil, ErrOutOfStock
    }

    if err := s.prodRepo.DecreaseQuantity(ctx, p.ID, 1); err != nil {
        return nil, ErrOutOfStock
    }

    if err := s.custRepo.UpdatePoints(ctx, customerID, -cost); err != nil {
        return nil, err
    }

    t := &Transaction{
        ID:         uuid.NewString(),
        CustomerID: customerID,
        ProductID:  p.ID,
        Size:       size,
        Flavor:     p.Flavor,
        Quantity:   1,
        CreatedAt:  time.Now().UTC(),
    }
    if err := s.repo.Create(ctx, t); err != nil {
        return nil, err
    }
    return t, nil
}

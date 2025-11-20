package customer

import (
    "context"

    "gorm.io/gorm"
)

type Repository interface {
    Create(ctx context.Context, c *Customer) error
    FindByID(ctx context.Context, id uint) (*Customer, error)
    UpdatePoints(ctx context.Context, id uint, delta int) error
}

type repository struct{ db *gorm.DB }

func NewCustomerRepository(db *gorm.DB) Repository { return &repository{db: db} }

func (r *repository) Create(ctx context.Context, c *Customer) error {
    return r.db.WithContext(ctx).Create(c).Error
}

func (r *repository) FindByID(ctx context.Context, id uint) (*Customer, error) {
    var c Customer
    if err := r.db.WithContext(ctx).First(&c, id).Error; err != nil {
        return nil, err
    }
    return &c, nil
}

func (r *repository) UpdatePoints(ctx context.Context, id uint, delta int) error {
    return r.db.WithContext(ctx).Model(&Customer{}).Where("id = ?", id).UpdateColumn("points", gorm.Expr("points + ?", delta)).Error
}

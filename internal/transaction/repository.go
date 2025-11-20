package transaction

import (
    "context"

    "gorm.io/gorm"
)

type Repository interface {
    Create(ctx context.Context, t *Transaction) error
}

type repository struct{ db *gorm.DB }

func NewTransactionRepository(db *gorm.DB) Repository { return &repository{db: db} }

func (r *repository) Create(ctx context.Context, t *Transaction) error {
    return r.db.WithContext(ctx).Create(t).Error
}

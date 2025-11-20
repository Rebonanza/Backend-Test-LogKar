package product

import (
    "context"

    "gorm.io/gorm"
)

type Repository interface {
    Create(ctx context.Context, p *Product) error
    FindByID(ctx context.Context, id uint) (*Product, error)
    List(ctx context.Context) ([]Product, error)
    DecreaseQuantity(ctx context.Context, id uint, amount int) error
    FindAvailableBySize(ctx context.Context, size string) (*Product, error)
}

type repository struct {
    db *gorm.DB
}

func NewProductRepository(db *gorm.DB) Repository {
    return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, p *Product) error {
    return r.db.WithContext(ctx).Create(p).Error
}

func (r *repository) FindByID(ctx context.Context, id uint) (*Product, error) {
    var p Product
    if err := r.db.WithContext(ctx).First(&p, id).Error; err != nil {
        return nil, err
    }
    return &p, nil
}

func (r *repository) List(ctx context.Context) ([]Product, error) {
    var out []Product
    if err := r.db.WithContext(ctx).Find(&out).Error; err != nil {
        return nil, err
    }
    return out, nil
}

func (r *repository) DecreaseQuantity(ctx context.Context, id uint, amount int) error {
    return r.db.WithContext(ctx).Model(&Product{}).Where("id = ? AND quantity >= ?", id, amount).UpdateColumn("quantity", gorm.Expr("quantity - ?", amount)).Error
}

func (r *repository) FindAvailableBySize(ctx context.Context, size string) (*Product, error) {
    var p Product
    if err := r.db.WithContext(ctx).Where("size = ? AND quantity > 0", size).First(&p).Error; err != nil {
        return nil, err
    }
    return &p, nil
}

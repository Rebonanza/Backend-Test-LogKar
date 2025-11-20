package db

import (
	"fmt"
	"log"

	"github.com/local/be-test-logkar/internal/config"
	"github.com/local/be-test-logkar/internal/user"
	"github.com/local/be-test-logkar/internal/product"
	"github.com/local/be-test-logkar/internal/customer"
	"github.com/local/be-test-logkar/internal/transaction"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGormDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDB)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto-migrate example model(s) and domain tables
	if err := db.AutoMigrate(&user.User{}, &product.Product{}, &customer.Customer{}, &transaction.Transaction{}); err != nil {
		log.Printf("warning: auto-migrate failed: %v", err)
	}

	return db, nil
}

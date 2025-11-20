package transaction

import (
	"time"
)

type Transaction struct {
	ID         string    `gorm:"primaryKey;size:36" json:"id"`
	CustomerID uint      `json:"customer_id"`
	ProductID  uint      `json:"product_id"`
	Size       string    `gorm:"size:50" json:"size"`
	Flavor     string    `gorm:"size:100" json:"flavor"`
	Quantity   int       `json:"quantity"`
	CreatedAt  time.Time `json:"created_at"`
}

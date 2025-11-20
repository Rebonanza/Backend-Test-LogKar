package user

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Email     string    `gorm:"uniqueIndex;size:255" json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

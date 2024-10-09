package domain

import (
	"time"

	"gorm.io/gorm"
)

// Auth model
type Auth struct {
	ID        int       `gorm:"column:id;primary_key" json:"id"`
	Field     string    `gorm:"column:field;type:varchar(255)" json:"field"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName return table name of Auth model
func (Auth) TableName() string {
	return "auths"
}

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	// Store     Store          `json:"store" gorm:"foreignKey:ID;references:UserID"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

func (User) TableName() string {
	return "users"
}

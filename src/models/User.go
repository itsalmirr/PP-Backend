package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Avatar    string    `gorm:"type:varchar(255)"`
	Email     string    `gorm:"type:varchar(50);uniqueIndex;not null"`
	Username  string    `gorm:"type:varchar(15);uniqueIndex;not null"`
	FullName  string    `gorm:"type:varchar(100);not null"`
	StartDate time.Time `gorm:"type:timestamp with time zone;default:CURRENT_TIMESTAMP"`
	IsStaff   bool      `gorm:"default:false"`
	IsActive  bool      `gorm:"default:true"`
	Password  string    `gorm:"type:varchar(128);not null"`

	// Standard GORM timestamps
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (User) TableName() string {
	return "users"
}

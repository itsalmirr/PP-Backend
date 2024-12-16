package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Realtor struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"` // Primary key with UUID
	FullName    string    `gorm:"type:varchar(100);not null"`
	Photo       string    `gorm:"type:varchar(255)"`                                       // Optional photo path
	Description string    `gorm:"type:text"`                                               // Optional description
	Phone       string    `gorm:"type:varchar(15);not null"`                               // Phone number
	Email       string    `gorm:"type:varchar(50);not null"`                               // Email address
	IsMVP       bool      `gorm:"default:false"`                                           // Boolean field
	HireDate    time.Time `gorm:"type:timestamp with time zone;default:CURRENT_TIMESTAMP"` // Hire date

	// Standard GORM timestamps
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// TableName overrides the default table name for the Realtor model
func (Realtor) TableName() string {
	return "realtors"
}

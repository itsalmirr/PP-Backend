package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Avatar     string    `gorm:"type:varchar(255)" json:"avatar,omitempty"`
	Email      string    `gorm:"type:varchar(120);uniqueIndex;not null" json:"email" validate:"required,email"`
	Username   string    `gorm:"type:varchar(120);uniqueIndex;not null" json:"username" validate:"required,min=3"`
	FullName   string    `gorm:"type:varchar(100);not null" json:"full_name"`
	StartDate  time.Time `gorm:"autoCreateTime" json:"start_date"`
	IsStaff    bool      `gorm:"default:false" json:"is_staff"`
	IsActive   bool      `gorm:"default:true" json:"is_active"`
	Password   string    `gorm:"type:varchar(128);not null" json:"password,omitempty"`
	Provider   string    `gorm:"default:email" json:"provider"`
	ProviderID string    `gorm:"default:null" json:"provider_id,omitempty"`

	// Standard GORM timestamps
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

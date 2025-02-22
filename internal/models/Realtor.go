package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Realtor struct {
	ID          uuid.UUID      `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	FullName    string         `gorm:"type:varchar(100);not null;index" json:"full_name" validate:"required,min=5,max=100"`
	Photo       datatypes.JSON `gorm:"type:jsonb" json:"photo,omitempty" validate:"omitempty,json"`
	Description string         `gorm:"type:text" json:"description,omitempty" validate:"max=500"`
	Phone       string         `gorm:"type:varchar(20);not null" json:"phone" validate:"required,e164"`
	Email       string         `gorm:"type:varchar(255);not null" json:"email" validate:"required,email"`
	IsMVP       bool           `gorm:"default:false" json:"is_mvp"`
	HireDate    time.Time      `gorm:"autoCreateTime" json:"hire_date"`
	Listings    []Listing      `gorm:"foreignKey:RealtorID;constraint:OnDelete:SET NULL;" json:"listings,omitempty"`

	// Standard GORM timestamps
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

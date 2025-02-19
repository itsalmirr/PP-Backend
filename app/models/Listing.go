package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/datatypes"
)

type Listing struct {
	ID             uuid.UUID       `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Title          string          `gorm:"type:varchar(120);index" json:"title" validate:"required,min=10"`
	Address        string          `gorm:"type:varchar(255);uniqueIndex" json:"address" validate:"required"`
	City           string          `gorm:"type:varchar(255)" json:"city" validate:"required"`
	State          string          `gorm:"type:varchar(3)" json:"state" validate:"required,uppercase,len=2"`
	ZipCode        string          `gorm:"type:varchar(6)" json:"zip_code" validate:"required,numeric,len=5"`
	Description    string          `gorm:"type:text" json:"description"`
	Price          decimal.Decimal `gorm:"type:decimal(12,2)" json:"price" validate:"required,gt=0"`
	Bedroom        int             `gorm:"type:int" json:"bedroom" validate:"gte=0"`
	Bathroom       float64         `gorm:"type:decimal(4,1)" json:"bathroom" validate:"gte=0"`
	Garage         int             `gorm:"type:int" json:"garage,omitzero"`
	Sqft           int64           `gorm:"type:int" json:"sqft"`
	TypeOfProperty string          `gorm:"type:varchar(255);index" json:"type_of_property" validate:"required"`
	LotSize        int64           `gorm:"type:int" json:"lot_size,omitzero" validate:"gte=0"`
	Pool           bool            `gorm:"type:bool" json:"pool,omitempty"`
	YearBuilt      int             `gorm:"type:int" json:"year_built" validate:"gte=1800,lte=2025"`
	// Media fields with cloudinary
	Media  datatypes.JSON `gorm:"type:jsonb" json:"media,omitempty" validate:"mediaarray"`
	Status string         `gorm:"type:varchar(20);default:'DRAFT'" json:"status" validate:"oneof=DRAFT PUBLISHED ARCHIVED"`

	// ForeignKey to Realtor model
	RealtorID *uuid.UUID `gorm:"type:uuid;index" json:"realtor_id"`
	Realtor   Realtor    `gorm:"foreignKey:RealtorID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"realtor"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

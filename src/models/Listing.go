package models

import (
	"time"

	"github.com/google/uuid"
)

type Listing struct {
	ID             uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"` // Primary key with UUID
	Title          string    `gorm:"type:varchar(120) unique"`
	Address        string    `gorm:"type:varchar(255) unique"`
	City           string    `gorm:"type:varchar(255) unique"`
	State          string    `gorm:"type:varchar(3) unique"`
	ZipCode        string    `gorm:"type:varchar(6) unique"`
	Description    string    `gorm:"type:text"`
	Price          string    `gorm:"type:varchar(255)"`
	Bedroom        int       `gorm:"type:int"`
	Bathroom       float32   `gorm:"type:float"`
	Garage         int       `gorm:"type:int"`
	Sqft           int64     `gorm:"type:int"`
	TypeOfProperty string    `gorm:"type:varchar(255) unique"`
	LotSize        int64     `gorm:"type:int"`
	Pool           bool      `gorm:"type:bool"`
	YearBuilt      string    `gorm:"type:varchar(255)"`
	// Media fields with cloudinary
	PhotoMain string `gorm:"type:text"`
	Photo1    string `gorm:"type:text"`
	Photo2    string `gorm:"type:text"`
	Photo3    string `gorm:"type:text"`
	Photo4    string `gorm:"type:text"`
	Photo5    string `gorm:"type:text"`

	IsPublished bool      `gorm:"type:bool"`
	PublishDate time.Time `gorm:"type:date"`
	// ForeignKey to Realtor model
	RealtorID uuid.UUID `gorm:"type:uuid;not null"`
	Realtor   Realtor   `gorm:"foreignKey:RealtorID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

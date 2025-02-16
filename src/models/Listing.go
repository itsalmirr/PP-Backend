package models

import (
	"time"

	"github.com/google/uuid"
)

type Listing struct {
	ID             uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"` // Primary key with UUID
	Title          string    `gorm:"type:varchar(120) unique" json:"title"`
	Address        string    `gorm:"type:varchar(255) unique" json:"address"`
	City           string    `gorm:"type:varchar(255) unique" json:"city"`
	State          string    `gorm:"type:varchar(3) unique" json:"state"`
	ZipCode        string    `gorm:"type:varchar(6) unique" json:"zip_code"`
	Description    string    `gorm:"type:text" json:"description"`
	Price          string    `gorm:"type:varchar(255)" json:"price"`
	Bedroom        int       `gorm:"type:int" json:"bedroom"`
	Bathroom       float32   `gorm:"type:float" json:"bathroom"`
	Garage         int       `gorm:"type:int" json:"garage,omitzero"`
	Sqft           int64     `gorm:"type:int" json:"sqft"`
	TypeOfProperty string    `gorm:"type:varchar(255) unique" json:"type_of_property"`
	LotSize        int64     `gorm:"type:int" json:"lot_size,omitzero"`
	Pool           bool      `gorm:"type:bool" json:"pool"`
	YearBuilt      string    `gorm:"type:varchar(255)" json:"year_built"`
	// Media fields with cloudinary
	PhotoMain string `gorm:"type:text" json:"photo_main"`
	Photo1    string `gorm:"type:text" json:"photo_1,omitempty"`
	Photo2    string `gorm:"type:text" json:"photo_2,omitempty"`
	Photo3    string `gorm:"type:text" json:"photo_3,omitempty"`
	Photo4    string `gorm:"type:text" json:"photo_4,omitempty"`
	Photo5    string `gorm:"type:text" json:"photo_5,omitempty"`

	IsPublished bool      `gorm:"type:bool" json:"is_published"`
	PublishDate time.Time `gorm:"type:date" json:"publish_date"`
	// ForeignKey to Realtor model
	RealtorID uuid.UUID `gorm:"type:uuid;not null" json:"realtor_id"`
	Realtor   Realtor   `gorm:"foreignKey:RealtorID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"realtor"`
}

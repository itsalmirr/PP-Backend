package repositories

type CreateListingInput struct {
	Title          string  `json:"title"`
	Address        string  `json:"address"`
	City           string  `json:"city"`
	State          string  `json:"state"`
	ZipCode        string  `json:"zip_code"`
	Description    string  `json:"description"`
	Price          string  `json:"price"`
	Bedroom        int     `json:"bedroom"`
	Bathroom       float32 `json:"bathroom"`
	Garage         int     `json:"garage"`
	Sqft           int64   `json:"sqft"`
	TypeOfProperty string  `json:"type_of_property"`
	LotSize        int64   `json:"lot_size"`
	Pool           bool    `json:"pool"`
	YearBuilt      string  `json:"year_built"`

	PhotoMain string `json:"photo_main"`
	Photo1    string `json:"photo_1"`
	Photo2    string `json:"photo_2"`
	Photo3    string `json:"photo_3"`
	Photo4    string `json:"photo_4"`
	Photo5    string `json:"photo_5"`

	IsPublished bool   `json:"is_published"`
	RealtorID   string `json:"realtor_id"`
}

func CreateListingRepository(data CreateRealtorInput) {}

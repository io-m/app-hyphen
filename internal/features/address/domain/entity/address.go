package address

type Address struct {
	Id          string  `json:"id" db:"id"`
	StreetName  string  `json:"street_name" db:"street_name"`
	HouseNumber string  `json:"house_number" db:"house_number"`
	City        string  `json:"city" db:"city"`
	ZipCode     string  `json:"zip_code" db:"zip_code"`
	Country     string  `json:"country" db:"country"`
	State       *string `json:"state,omitempty" db:"state"`
	Region      *string `json:"region,omitempty" db:"region"`
	ExtraInfo   *string `json:"extra_info,omitempty" db:"extra_info"`
	CreatedAt   *string `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt   *string `json:"updated_at,omitempty" db:"updated_at"`
}

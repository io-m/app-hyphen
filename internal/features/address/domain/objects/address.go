package address_objects

type AddressRequest struct {
	StreetName  string  `json:"street_name"`
	HouseNumber string  `json:"house_number"`
	City        string  `json:"city"`
	ZipCode     string  `json:"zip_code"`
	Country     string  `json:"country"`
	State       *string `json:"state,omitempty"`
	Region      *string `json:"region,omitempty"`
	ExtraInfo   *string `json:"extra_info,omitempty"`
	CreatedAt   *string `json:"created_at"`
	UpdatedAt   *string `json:"updated_at,omitempty"`
}

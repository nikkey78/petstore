package models

type Location struct {
	ID           int    `json:"id"`
	LocationName string `json:"location_name"`
}

type LocationService interface {
	CreateLocation() (int, error)
	GetAllLocations() ([]*Location, error)
	// GetLocation(id int)(*Location, error)
}

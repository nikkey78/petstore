package models

type Location struct {
	ID int `json:"id"`
	LocationName string `json:"locationName"`
}

type LocationService interface{
	CreateLocation() error
	GetAllLocations()([]*Location, error)
	GetLocation(id int)(*Location, error)
}

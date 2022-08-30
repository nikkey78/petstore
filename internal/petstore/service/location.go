package service

import (
	db "petstore/internal/petstore/repo/sqlite"
	"petstore/pkg/models"
)

var ls db.LocationService
var loc db.Location

func CreateLocation(l db.Location) (int, error) {
	ls = l
	id, err := ls.CreateLocation()
	if err != nil {
		return id, err
	}
	return id, nil
}

func GetAllLocations() ([]*models.Location, error) {
	ls = &loc
	locs, err := ls.GetAllLocations()
	if err != nil {
		return nil, err
	}
	return locs, nil
}

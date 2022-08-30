package sqlite

import (
	"fmt"
	client "petstore/internal/database/sqlite"
	"petstore/pkg/models"
)

type Location models.Location
type LocationService models.LocationService

func (l Location) CreateLocation() (int, error) {
	result, err := client.DbClient.Exec("INSERT INTO location (location_name) VALUES ($1);", l.LocationName)
	if err != nil {
		return l.ID, fmt.Errorf("Inserted failed. error: %v", err)
	}
	insertedId, err := result.LastInsertId()
	if err != nil {
		return l.ID, fmt.Errorf("Inserted ID failed. error: %v", err)
	}
	l.ID = int(insertedId)
	return l.ID, nil
}

func (l Location) GetAllLocations() ([]*models.Location, error) {
	var locations []*models.Location
	rows, err := client.DbClient.Query("SELECT * FROM location;")
	if err != nil {
		return nil, fmt.Errorf("Query to database failed. error: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&l.ID, &l.LocationName)
		newLocation := models.Location(l)
		locations = append(locations, &newLocation)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("database error: %v", err)
	}
	return locations, nil
}

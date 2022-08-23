package models

type Pet struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Age         float64  `json:"age"`
	Description string   `json:"description"`
	ImageURL    string   `json:"imageUrl"`
	PetLocation Location `json:"location"`
	PetCategory Category `json:"category"`
	PetBreed    Breed    `json:"breed"`
}

type PetService interface {
	CreatePet() error
	GetPet(id int) (*Pet, error)
	GetPetsByCategory(catogoryID int) ([]*Pet, error)
	DeletePet(id int) error
}

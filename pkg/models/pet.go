package models

type Pet struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Age         float64  `json:"age"`
	ImageURL    string   `json:"imageUrl"`
	Description string   `json:"description"`
	PetCategory Category `json:"category"`
	PetBreed    Breed    `json:"breed"`
	PetLocation Location `json:"location"`
}

type PetService interface {
	CreatePet() (int, error)
	GetPetsByCategory(catogoryID int) ([]*Pet, error)
	GetPet(id int) (*Pet, error)
	DeletePet(id int) (int, error)
}


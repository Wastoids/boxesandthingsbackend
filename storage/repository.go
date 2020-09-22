package storage

import (
	"github.com/Wastoids/boxesandthingsbackend/models"
)

// Repository is a concrete implementation of the repository
type Repository struct {
	db dynamo
}

// NewRepository is a constructor for the repository implementation
func NewRepository() Repository {
	return Repository{db: newDynamo()}
}

// GetBoxesByEmail gets the boxes of a user by email
func (r Repository) GetBoxesByEmail(email string) ([]*models.Box, error) {
	return r.db.getBoxesByEmail(email)
}

package storage

import (
	"github.com/Wastoids/boxesandthingsbackend/domain"
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
func (r Repository) GetBoxesByEmail(email string) ([]*domain.Box, error) {
	return r.db.getBoxesByEmail(email)
}

// GetThingsByBox gets the things by box id
func (r Repository) GetThingsByBox(boxID string) ([]*domain.Thing, error) {
	return r.db.getThingsByBox(boxID)
}

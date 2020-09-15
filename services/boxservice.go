package services

import (
	"errors"

	"github.com/Wastoids/boxesandthingsbackend/models"
)

// Repository represents the contract between the application layer and the data layer
type Repository interface {
	GetBoxesByEmail(email string) ([]*models.Box, error)
}

// BoxServicer represents an interface to the services exposed for boxes
type BoxServicer interface {
	GetBoxesByEmail(email string) ([]*models.Box, error)
}

// BoxService provides an implementation for BoxServicer interface
type BoxService struct {
	repository Repository
}

// GetBoxesByEmail is an implementation of the method exposed by BoxServicer
func (b BoxService) GetBoxesByEmail(email string) ([]*models.Box, error) {
	if len(email) == 0 {
		return nil, errors.New("invalid email")
	}

	boxes, err := b.repository.GetBoxesByEmail(email)
	if err != nil {
		return nil, errors.New("something went wrong")
	}

	return boxes, nil
}

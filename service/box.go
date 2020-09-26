package service

import (
	"errors"

	"github.com/Wastoids/boxesandthingsbackend/domain"
)

// BoxServicer represents an interface to the services exposed for boxes
type BoxServicer interface {
	GetTopBoxesByEmail(email string) ([]*domain.Box, error)
}

// BoxService provides an implementation for BoxServicer interface
type BoxService struct {
	repository domain.BoxRepository
}

// NewBoxService is a function which returns an instance of the BoxService
func NewBoxService(repo domain.BoxRepository) BoxService {
	return BoxService{repository: repo}
}

// GetTopBoxesByEmail is an implementation of the method exposed by BoxServicer
func (b BoxService) GetTopBoxesByEmail(email string) ([]*domain.Box, error) {
	if len(email) == 0 {
		return nil, errors.New("invalid email")
	}

	boxes, err := b.repository.GetBoxesByEmail(email)
	if err != nil {
		return nil, errors.New("something went wrong")
	}

	return boxes, nil
}

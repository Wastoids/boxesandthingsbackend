package service

import (
	"github.com/Wastoids/boxesandthingsbackend/domain"
	"github.com/sirupsen/logrus"
)

// ThingServicer represents the services offered
type ThingServicer interface {
	GetThingsByBox(boxID string) ([]*domain.Thing, error)
}

// ThingService implements the methods exposed by the service interface
type ThingService struct {
	repo domain.ThingRepository
}

func NewThingService(repo domain.ThingRepository) *ThingService {
	return &ThingService{repo: repo}
}

// GetThingsByBox is an implementaion of the service method
func (t ThingService) GetThingsByBox(boxID string) (things []*domain.Thing, err error) {
	if len(boxID) == 0 {
		logrus.Info("invalid boxID")
		return nil, nil
	}

	if things, err = t.repo.GetThingsByBox(boxID); err != nil {
		logrus.Errorf("something went wrong while getting the things by box: %v", err)
		return nil, err
	}

	return
}

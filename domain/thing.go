package domain

import "time"

// Thing represents any item or artifact that the user wants to store in a box
type Thing struct {
	ID          string
	BoxID       string
	Name        string
	Description string
	ExpiresOn   time.Time
}

// ThingRepository represents the repository methods it depends on
type ThingRepository interface {
	GetThingsByBox(boxID string) ([]*Thing, error)
}

package models

// Box represents a basic box in the system
type Box struct {
	ID          string
	Name        string
	Description string
	ParentBoxID string
}

// Repository represents the methods that have to be exposed by the data layer
type Repository interface {
	GetBoxesByEmail(email string) ([]*Box, error)
}

package domain

// Box represents a basic box in the system
type Box struct {
	ID          string
	Name        string
	Description string
	ParentBoxID string
}

// BoxRepository represents the methods that have to be exposed by the data layer
type BoxRepository interface {
	GetBoxesByEmail(email string) ([]*Box, error)
}

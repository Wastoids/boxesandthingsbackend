package models

// Box represents a basic box in the system
type Box struct {
	ID             string
	Name           string
	Description    string
	ContainedBoxes []*Box
}

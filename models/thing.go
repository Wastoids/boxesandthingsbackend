package models

import "time"

// Thing represents any item or artifact that the user wants to store in a box
type Thing struct {
	ID          string
	Box         *Box
	Name        string
	Description string
	ExpiresOn   time.Time
}

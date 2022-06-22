package model

// User represents the user for this application
//
// swagger:model User
type User struct {
	// the name for this user
	// required: true
	// min length: 5
	Name string `json:"name"`

	// The birth year for the user
	// min: 1900
	// max: 2022
	BirthYear int `json:"birth_year"`
}

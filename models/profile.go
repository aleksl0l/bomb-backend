package models

import "fmt"

type Profile struct {
	Id     int
	Email  string
	UserId int
}

type ProfileCreationError struct {
	err string
}

func (e *ProfileCreationError) Error() string {
	return fmt.Sprintf("profile creation failed: %s", e.err)
}

package models

type BaseModel interface {
	LoadFromDatabase() bool
	SaveToDatabase() bool
}

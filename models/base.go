package models

type BaseModel interface {
	SaveToDatabaseQuery() string
	LoadFromDatabaseQuery() string
}


func (m *BaseModel)SaveToDatabase()
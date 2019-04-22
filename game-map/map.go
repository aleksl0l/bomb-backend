package game_map

import (
	"github.com/aleksl0l/bomb-backend/message"
)

type GameMap interface {
	SetCell(obj *message.GameObject)
	RemoveObject(id int32) bool
	Field() [][]int
}

type NaiveMap struct {
	Objects []*message.GameObject
	size    int
}

func NewNaiveMap(size int) *NaiveMap {
	return &NaiveMap{
		make([]*message.GameObject, 0),
		size,
	}
}

func (m *NaiveMap) SetCell(obj *message.GameObject) {
	m.Objects = append(m.Objects, obj)
}

func (m *NaiveMap) RemoveObject(id int32) bool {
	for i, o := range m.Objects {
		if o.ID == id {
			m.Objects = append(m.Objects[:i], m.Objects[i+1:]...)
			return true
		}
	}
	return false
}

func (m *NaiveMap) Field() [][]int {
	newField := make([][]int, m.size)

	for i := range newField {
		newField[i] = make([]int, m.size)
	}

	// Fill field with default bomberman pattern
	// https://st.emulroom.com/images/dendy/bomberman/bomberman-(4).png
	for i := 1; i < m.size-1; i += 2 {
		for j := 1; j < m.size-1; j += 2 {
			newField[i][j] = -1
		}
	}

	for _, o := range m.Objects {
		newField[o.Position.X][o.Position.Y] = int(o.ID)
	}

	return newField
}

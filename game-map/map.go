package game_map

import (
	"github.com/aleksl0l/bomb-backend/game-object"
)

type GameMap interface {
	SetCell(obj game_object.GameObject)
	RemoveObject(id int) bool
	Field() [][]int
}

type NaiveMap struct {
	objects []game_object.GameObject
	size    int
}

func NewNaiveMap(size int) NaiveMap {
	return NaiveMap{make([]game_object.GameObject, 0), size}
}

func (m *NaiveMap) SetCell(obj game_object.GameObject) {
	m.objects = append(m.objects, obj)
}

func (m *NaiveMap) RemoveObject(id int) bool {
	for i, o := range m.objects {
		if o.Id == id {
			m.objects = append(m.objects[:i], m.objects[i+1:]...)
			return true
		}
	}
	return false
}

func (m NaiveMap) Field() [][]int {
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

	for _, o := range m.objects {
		newField[o.Position.X][o.Position.Y] = o.Id
	}

	return newField
}

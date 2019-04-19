package game_map

import (
	"github.com/aleksl0l/bomb-backend/game-object"
	"github.com/aleksl0l/bomb-backend/player"
)

type GameMap interface {
	setCell(target *GameMap, obj game_object.GameObject)
	setPlayer(target *GameMap, obj player.Player)
}

type NaiveMap struct {
	objects []game_object.GameObject
	field   [][]int
}

func NewNaiveMap(size int) NaiveMap {
	newField := make([][]int, size)

	for i := range newField {
		newField[i] = make([]int, size)
	}

	// Fill field with default bomberman pattern
	// https://st.emulroom.com/images/dendy/bomberman/bomberman-(4).png
	for i := 1; i < size-1; i += 2 {
		for j := 1; j < size-1; j += 2 {
			newField[i][j] = -1
		}
	}
	return NaiveMap{make([]game_object.GameObject, 0), newField}
}

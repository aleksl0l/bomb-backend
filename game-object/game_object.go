package game_object

type Point struct {
	x uint16
	y uint16
}

type GameObject struct {
	Id       int
	Type     string
	Position Point
}

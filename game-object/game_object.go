package game_object

type Point struct {
	X uint16
	Y uint16
}

type GameObject struct {
	Id       int
	Type     string
	Position Point
}

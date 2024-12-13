package shared

import "fmt"

type Vector2 struct {
	X int
	Y int
}

func (v Vector2) String() string {
	return "(" + fmt.Sprint(v.X) + ", " + fmt.Sprint(v.Y) + ")"
}

type Direction int

const (
	UP    Direction = 1
	RIGHT Direction = 2
	DOWN  Direction = 3
	LEFT  Direction = 4
)

func (d Direction) String() string {
	switch d {
	case UP:
		return "UP"
	case RIGHT:
		return "RIGHT"
	case DOWN:
		return "DOWN"
	case LEFT:
		return "LEFT"
	}
	return ""
}

type Tile int

const (
	OBSTACLE Tile = 0
	EMPTY    Tile = 1
	OUTSIDE  Tile = 2
)

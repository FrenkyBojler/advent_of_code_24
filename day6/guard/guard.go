package guard

import (
	. "frenkybojler/adventofcode24/day6/level"
	. "frenkybojler/adventofcode24/day6/shared"
	"slices"
)

type Guard struct {
	Position             Vector2
	Direction            Direction
	ShouldCheckForLoop   bool
	checkedLoopPositions []string
}

func (g *Guard) GetCurrentRotationImage() string {
	switch g.Direction {
	case UP:
		return "^"
	case RIGHT:
		return ">"
	case DOWN:
		return "v"
	case LEFT:
		return "<"
	}
	return ""
}

// Returns true if guard can move forward, false otherwise
// Returns true if guard really moved forward, false if only rotated
func (g *Guard) MoveForward(l *Level) (bool, bool) {
	nextPosition := Vector2{X: g.Position.X, Y: g.Position.Y}

	switch g.Direction {
	case UP:
		nextPosition.Y--
	case RIGHT:
		nextPosition.X++
	case DOWN:
		nextPosition.Y++
	case LEFT:
		nextPosition.X--
	}

	if g.ShouldCheckForLoop && l.IsEmptyTile(nextPosition) {
		isLoopPosition := g.CheckForLoopInStartingFromPosition(nextPosition, *l)

		if isLoopPosition {
			l.AddLoopPosition(nextPosition)
		}
	}

	if l.GetTileAtPosition(nextPosition) == EMPTY {
		g.MoveGuardToPosition(nextPosition, l)
		return true, true
	}

	if l.GetTileAtPosition(nextPosition) == OBSTACLE {
		g.Rotate(l)
		return true, false
	}

	return false, false
}

func (g *Guard) CheckForLoopInStartingFromPosition(target_position Vector2, level Level) bool {
	level_copy := level.Copy()

	//fmt.Println(level_copy.Print(target_position, false))

	guard := GetGuardFromLevel(level_copy, false)
	level_copy.LevelMap[target_position.Y][target_position.X] = "#"

	visited_positions := []string{}

	for {
		visited_positions = append(visited_positions, getPositionDirectionID(guard.Position, guard.Direction))
		can_move, _ := guard.MoveForward(&level_copy)

		if !can_move {
			return false
		}

		if slices.Contains(visited_positions, getPositionDirectionID(guard.Position, guard.Direction)) {
			return true
		}
	}
}

func (g *Guard) Rotate(l *Level) {
	g.Direction++
	if g.Direction > LEFT {
		g.Direction = UP
	}

	l.LevelMap[g.Position.Y][g.Position.X] = g.GetCurrentRotationImage()
}

func GetGuardFromLevel(l Level, shouldCheckForLoops bool) Guard {
	for y, row := range l.LevelMap {
		for x, cell := range row {
			switch cell {
			case "^":
				return Guard{Position: Vector2{X: x, Y: y}, Direction: UP, ShouldCheckForLoop: shouldCheckForLoops}
			case ">":
				return Guard{Position: Vector2{X: x, Y: y}, Direction: RIGHT, ShouldCheckForLoop: shouldCheckForLoops}
			case "v":
				return Guard{Position: Vector2{X: x, Y: y}, Direction: DOWN, ShouldCheckForLoop: shouldCheckForLoops}
			case "<":
				return Guard{Position: Vector2{X: x, Y: y}, Direction: LEFT, ShouldCheckForLoop: shouldCheckForLoops}
			}
		}
	}

	return Guard{}
}

func (g *Guard) MoveGuardToPosition(position Vector2, l *Level) {
	l.LevelMap[g.Position.Y][g.Position.X] = "."
	l.LevelMap[position.Y][position.X] = g.GetCurrentRotationImage()
	g.Position = position

	l.VisitTile(position)
}

func (g *Guard) IsPositionAndDirectionChecked(position Vector2, direction Direction) bool {
	id := getPositionDirectionID(position, direction)

	return slices.Contains(g.checkedLoopPositions, id)
}

func (g *Guard) AddPositionAndDirectionToCheckedLoopPositions(position Vector2, direction Direction) {
	id := getPositionDirectionID(position, direction)

	if !slices.Contains(g.checkedLoopPositions, id) {
		g.checkedLoopPositions = append(g.checkedLoopPositions, id)
	}
}

func getPositionDirectionID(position Vector2, direction Direction) string {
	return position.String() + direction.String()
}

package level

import (
	. "frenkybojler/adventofcode24/day6/shared"
	"slices"

	"strings"

	"github.com/samber/lo"
)

type Level struct {
	LevelMap                    [][]string
	uniqueTilesVisited          []Vector2
	UniqueObstaclesToCreateLoop []Vector2
}

func (l *Level) Print(origin Vector2, loopObstacles bool) string {
	size := 10

	result := ""

	level_map_copy := lo.Map(l.LevelMap, func(row []string, _ int) []string {
		return append([]string{}, row...)
	})

	if loopObstacles {
		for _, obstacle := range l.UniqueObstaclesToCreateLoop {
			level_map_copy[obstacle.Y][obstacle.X] = "O"
		}
	}

	for y := origin.Y - size; y <= origin.Y+size; y++ {
		if y < 0 || y >= len(level_map_copy) {
			continue
		}
		for x := origin.X - size; x <= origin.X+size; x++ {
			if x < 0 || x >= len(level_map_copy[y]) {
				continue
			}
			result += level_map_copy[y][x] + " "
		}
		result += "\n"
	}

	return result
}

func LoadLevel(fileContent string) Level {
	level := Level{}

	level.LevelMap = lo.Map(strings.Split(fileContent, "\n"), func(line string, _ int) []string {
		return strings.Split(line, "")
	})

	return level
}

func (l *Level) VisitTile(position Vector2) {
	if !slices.Contains(l.uniqueTilesVisited, position) {
		l.uniqueTilesVisited = append(l.uniqueTilesVisited, position)
	}
}

func (l *Level) GetNumberOfUniqueTilesVisited() int {
	return len(l.uniqueTilesVisited)
}

func (l *Level) GetNumberOfUniqueObstaclesToCreateLoop() int {
	return len(l.UniqueObstaclesToCreateLoop)
}

func (l *Level) GetTileAtPosition(position Vector2) Tile {
	if position.Y < 0 || position.Y >= len(l.LevelMap) {
		return OUTSIDE
	}
	if position.X < 0 || position.X >= len(l.LevelMap[position.Y]) {
		return OUTSIDE
	}
	if l.LevelMap[position.Y][position.X] == "#" {
		return OBSTACLE
	}
	return EMPTY
}

func (l *Level) AddLoopPosition(position Vector2) {
	if !slices.Contains(l.UniqueObstaclesToCreateLoop, position) {
		l.UniqueObstaclesToCreateLoop = append(l.UniqueObstaclesToCreateLoop, position)
	}
}

func (l *Level) IsEmptyTile(point Vector2) bool {
	return l.GetTileAtPosition(point) == EMPTY
}

func (l *Level) Copy() Level {
	level_copy := Level{}
	level_copy.LevelMap = [][]string{}

	for _, row := range l.LevelMap {
		level_copy.LevelMap = append(level_copy.LevelMap, append([]string{}, row...))
	}

	return level_copy
}

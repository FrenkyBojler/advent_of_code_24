package main

import (
	"fmt"
	. "frenkybojler/adventofcode24/day6/guard"
	. "frenkybojler/adventofcode24/day6/level"
	. "frenkybojler/adventofcode24/shared"
)

func main() {
	file_content, _ := ReadFileContent("input.txt")
	levelMap := LoadLevel(file_content)
	guard := GetGuardFromLevel(levelMap, true)

	//pterm.Print("\n\n")
	//area, _ := pterm.DefaultArea.WithCenter().Start()

	//area.Update(levelMap.Print(guard.Position, false))
	//time.Sleep(1 * time.Second)

	for {
		can_move, _ := guard.MoveForward(&levelMap)

		if !can_move {
			break
		}
		//area.Update(levelMap.Print(guard.Position, false))
		//time.Sleep(100 * time.Millisecond)
	}

	//area.Stop()

	//pterm.Print(levelMap.Print(guard.Position, true))

	fmt.Println("Unique tiles vissited: ", levelMap.GetNumberOfUniqueTilesVisited())
	fmt.Println("Unique obstacles to create loop: ", levelMap.GetNumberOfUniqueObstaclesToCreateLoop())
	//fmt.Println("Loops created: ", levelMap.UniqueObstaclesToCreateLoop)
}

// 2344

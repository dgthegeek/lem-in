package helper

import (
	"fmt"
)

type AntPaths map[int][]string

func ValidateStartingConnections(startConnections []RelationAndDistance, nonCrossing [][]string) int {

	connections := 0
	starts := []string{}

	nonCrossingStarter := []string{}

	for _, v := range startConnections {
		starts = append(starts, v.Name)
	}

	for _, v := range nonCrossing {
		nonCrossingStarter = append(nonCrossingStarter, v[0])
	}

	for _, v := range starts {
		if contains(v, nonCrossingStarter) {
			connections++
		}
	}

	return connections
}

func BigTraversal(connections int, paths [][]string, shortest []string, ants int) {
	if len(paths) > 0 {
		canMove := []int{}
		antPaths := giveEachAntHisPath(ants, paths, shortest)
		maxMove := 0
		minMove := connections

		for _, v := range paths {
			maxMove += len(v)
		}

		for i := 1; i <= connections; i++ {
			canMove = append(canMove, i)
		}

		step := 1
		lastIndex := len(canMove) - 1
		last := canMove[lastIndex]

		for len(canMove) > 0 {
			last = makeAStep(&canMove, ants, &antPaths, last, step, maxMove, minMove)
			step++
		}
	}
}

func updateCanMove(lastValue, step, maxMove, minMove, ants int, canMove *[]int) int {
	goal := lastValue + minMove
	for i := lastValue + 1; i <= goal; i++ {
		if !containsAnt(i, canMove) && i <= ants && len(*canMove) <= maxMove {
			*canMove = append(*canMove, i)
			lastValue = i
		}
	}
	return lastValue
}

func makeAStep(canMove *[]int, ants int, antPaths *AntPaths, lastValue, step, maxMove, minMove int) int {
	for _, v := range *canMove {
		pat := (*antPaths)[v]
		if len(pat) > 0 {
			move := pat[0]
			makeAmove(move, v)
			pat = pat[1:]
			(*antPaths)[v] = pat
		}
	}

	for i, v := range *canMove {
		pat := (*antPaths)[v]
		if len(pat) == 0 {
			if i < len(*canMove) {
				*canMove = append((*canMove)[:i], (*canMove)[i+1:]...)
			}
		}
	}

	if len(*canMove) > 0 {
		fmt.Println()
		lastValue = updateCanMove(lastValue, step, maxMove, minMove, ants, canMove)
	}
	if len(*canMove) == 0 {
		fmt.Println()
	}

	return lastValue
}

func giveEachAntHisPath(ants int, paths [][]string, shortest []string) AntPaths {

	Sort2DArrayByLength(paths)
	paths = RemoveDuplicateFrom2DArray(paths)

	if len(paths) > 0 {

		antPaths := AntPaths{}
		pathIndex := 0

		for i := 1; i <= ants; i++ {

			antPaths[i] = paths[pathIndex]

			if i == ants {
				antPaths[i] = paths[0]
			}

			pathIndex = pathIndex + 1
			if pathIndex == len(paths) {
				pathIndex = 0
			}
		}
		return antPaths
	}
	return AntPaths{}
}

func containsAnt(ant int, ants *[]int) bool {
	for _, v := range *ants {
		if v == ant {
			return true
		}
	}
	return false
}

func contains(ant string, ants []string) bool {
	for _, v := range ants {
		if v == ant {
			return true
		}
	}
	return false
}

func makeAmove(move string, antNumb int) {
	fmt.Printf("L%d-%s ", antNumb, move)
}

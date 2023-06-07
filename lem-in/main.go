package main

import (
	"fmt"
	"lem-in/helper"
	"os"
)

func main() {

	args := os.Args[1:]

	if len(args) > 0 {
		filename := args[0]
		relations, rooms, ants := helper.ParseInputFile(filename)
		table := helper.DFS(relations, helper.PeekStartRoom(rooms).Name)

		// neeed to add another check to see if at least start and end are connected
		// if all the nodes are not connected ps: not mandatory we can do fine withoout it
		if len(table) != len(rooms) {
			fmt.Println("ERROR: invalid data format,a not all rooms are connected")
			os.Exit(0)
		}

		start := helper.PeekStartRoom(rooms).Name
		end := helper.PeekEndRoom(rooms).Name

		allPaths := helper.FindAllPaths(relations, start, end)
		trimmed := helper.RemoveStart(allPaths)
		groups := helper.MakeGroups(trimmed)

		_, thePromised := helper.GetSmallestPathOfEachGroup(groups, relations)
		elected := [][]string{thePromised}
		eligeables := helper.GetEligeables(thePromised, groups)
		elected = helper.Elector(eligeables, thePromised, elected, end, groups)

		connextions := helper.ValidateStartingConnections(relations[start], elected)
		helper.BigTraversal(connextions, elected, thePromised, ants)
	}
}

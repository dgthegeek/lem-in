package helper

import (
	"math"
)

type Group map[string][][]string

func getSourcesAndTargets(groups Group, key string)([][]string,[][]string){
	sources := groups[key]
	targets :=[][]string{}
	
	for k,v:=range groups{
		if k!=key{
			targets = append(targets, v...)
		}
	}
	return sources,targets
}

func GenerateAllPossibilities(thePromised []string, eligeables [][]string)[][][]string{
	groups:= MakeGroups(eligeables)
	allPossibilities :=[][][]string{}

	for key:=range groups{
		sources,targets:= getSourcesAndTargets(groups,key)
		allPossibilities= append(allPossibilities,GenerateCombinations(thePromised,targets,sources)...)
	}

	return allPossibilities
}


func GiveTheOneWithMostCandidates(paths[][]string,groups Group)[]string{
	thePromised:=[]string{}
	max:= math.MinInt32

	for _,path := range paths{
		eligeables:=GetEligeables(path,groups)
		eligeables = FindNonCrossingPaths(eligeables)
		_,flat:=Flat2DArray(eligeables)
		if len(flat)>max{
			max = len(flat)
			thePromised = path
		}
	}
	return thePromised
}

func GetEligeables(shortestOfAll []string,groups Group) [][]string{

	eligeables := [][]string{}

	sortedGroups:=TurnGroupToOrdered3DArray(groups)

	for _,paths:= range sortedGroups{
		for _,path:=range paths{
			if !HasCommonElements(shortestOfAll,path){
				eligeables = append(eligeables, path)
			}
		}	
	}

	return eligeables
}

func GetSmallestPathOfEachGroup(groups Group,relations Relation) ([][]string, []string) {
	smallestPaths := [][]string{}
	var theSmallestOfAll []string

	sortedGroup:= TurnGroupToOrdered3DArray(groups)

	for _, paths := range sortedGroup {
		Sort2DArrayByLength(paths)
		smallestPaths = append(smallestPaths, paths[0])
	}

	Sort2DArrayByLength(smallestPaths)

	theSmallestOfAll = GiveTheOneWithMostCandidates(smallestPaths,groups)

	return smallestPaths, theSmallestOfAll
}


func MakeGroups(paths[][]string)Group{
	groups:= Group{}

	for i,path:= range paths{
		if len(path)>0{
			groups[string(paths[i][0])]= append(groups[string(paths[i][0])], path)
		}
	}
	return groups
}
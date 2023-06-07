package helper

import (
	"math"
)

func Elector(eligeables [][]string, thePromised []string, elected [][]string, end string, groups Group) [][]string {
	if len(eligeables) > 0 {

		found :=0
		min :=math.MaxInt32

		allPossibilities := GenerateAllPossibilities(thePromised, eligeables)
		for _, v := range allPossibilities {
			if !HasCommonElements2(v) {
				found++
				_, flat := Flat2DArray(v)
				if len(flat) < min {
					min = len(flat)
					elected = v
				}
			}
		}

		if found == 0 {
			min := math.MaxInt32
			choosen := [][]string{}
			for key, subPath := range groups {
				if key != thePromised[0] {
					_, flat := Flat2DArray(subPath)
					if len(flat) > 0 && len(flat) < min {
						for _,sub:=range subPath{
							if !HasCommonElements(sub,thePromised){
								choosen = [][]string{sub}
								break
							}
						}
					}
				}
			}

			elected = append(elected, choosen...)

		}
	}
	return elected
}
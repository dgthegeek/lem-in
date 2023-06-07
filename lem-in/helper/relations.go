package helper

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type RelationAndDistance struct {
	Name string
	Distance int
}

type Relation map[string][]RelationAndDistance

func GetRelations(filename string,rooms []Room)Relation {
	relations:= Relation{}
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		parts:= strings.Split(line, "-")
		if len(parts)==2{
			if IsValidRoom(parts[1],rooms) && IsValidRoom(parts[0],rooms){

				if parts[0]==parts[1]{
					fmt.Println("ERROR: invalid data format,a room can't link to himself")
					os.Exit(0)
				}

				room1:= peekRoom(parts[0],rooms)
				room2 := peekRoom(parts[1],rooms)
				distance := getDistance(room1,room2)

				relations[room1.Name]=  append(relations[room1.Name], RelationAndDistance{
					Name:room2.Name ,
					Distance: distance,
				})

				relations[room2.Name]= append(relations[room2.Name], RelationAndDistance{
					Name: room1.Name,
					Distance: distance,
				})

			}else{
				fmt.Println("ERROR: invalid data format,Link to unknown room")
				os.Exit(0)
			}
		}
	}
	return relations
}

func IsValidRoom(name string, rooms []Room)bool{
	for _,v:=range rooms{
		if v.Name==name{
			return true
		}
	}
	return false
}

func peekRoom(name string, rooms []Room)Room{
	for _,v:=range rooms{
		if v.Name==name{
			return v
		}
	}
	return Room{}
}

func PeekStartRoom(rooms []Room)Room{
	for _,v:=range rooms{
		if v.RoomType=="start"{
			return v
		}
	}
	return Room{}
}

func PeekEndRoom(rooms []Room)Room{
	for _,v:=range rooms{
		if v.RoomType=="end"{
			return v
		}
	}
	return Room{}
}

func getDistance(room1 Room,room2 Room) int{
	x1 := room1.X
	x2 := room2.X
	y1 := room1.Y
	y2 := room2.Y
	dx := math.Pow(float64(x1-x2),2)
	dy :=  math.Pow(float64(y1-y2),2)
	distance := math.Sqrt(dx+dy)
	return int(distance)
}
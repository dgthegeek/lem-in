package helper

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Room struct {
	Name     string
	RoomType string
	X int
	Y int
}

func GetRooms(filename string) []Room {
	rooms := []Room{}
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}

	lines := strings.Split(string(content), "\n")
	if !strings.Contains(strings.Join(lines, "\n"),"##start"){
		fmt.Println("ERROR: invalid data format,No start room found")
		os.Exit(0)
	}
	if !strings.Contains(strings.Join(lines,"\n"),"##end"){
		fmt.Println("ERROR: invalid data format,No end room found")
		os.Exit(0)
	}

	startCount :=strings.Count(strings.Join(lines, "\n"),"##start")
	endCount := strings.Count(strings.Join(lines, "\n"),"##end")

	if startCount >1 || endCount >1{
		fmt.Println("ERROR: invalid data format, More than one start/end command found")
		os.Exit(0)
	}

	for i, line := range lines {
		parts:= strings.Split(line, " ")
		if len(parts)==3{
			room := makeRoom(parts,lines,i)
			if !Contains(room,rooms){
				rooms = append(rooms, room)
			}else{
				fmt.Println("ERROR: invalid data format,Duplicated Room")
				os.Exit(0)
			}
		}
	}
	return rooms
}

func makeRoom(parts []string,lines []string,i int)Room{
	name:= parts[0]
	if string(name[0])=="L"{
		fmt.Println("ERROR: invalid data format,Name should never start with L")
		os.Exit(0)
	}
	x,err := strconv.Atoi(parts[1])
	if err !=nil{
		fmt.Println("ERROR: invalid data format,Invalid coordinates")
		os.Exit(0)
	}

	y,err := strconv.Atoi(parts[2])
	if err !=nil{
		fmt.Println("ERROR: invalid data format,Invalid coordinates")
		os.Exit(0)
	}
	roomType:="normal"

	if i-1 >=0 && i-1 <len(lines){
		if lines[i-1]== "##start"{
			roomType= "start"
		}
		if lines[i-1]=="##end"{
			roomType="end"
		}
	}
	return Room{
		Name  : name,
		X: x,
		Y: y,
		RoomType: roomType,
	}
}

func Contains (value Room, arr[]Room)bool{
	for _,v:=range arr{
		if value.Name == v.Name{
			return true
		}
	}
	return false
}

func GetAntsNumber(filename string)int{
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	lines := strings.Split(string(content), "\n")
	antNum,err := strconv.Atoi(lines[0])
	if err !=nil || antNum <=0 {
		fmt.Println("ERROR: invalid data format, Invalid Ants number")
		os.Exit(0)
	}
	return antNum
}
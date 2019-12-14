package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Maze struct {
	X      int  `json:"x"`
	Y      int  `json:"y"`
	Top    bool `json:"top"`
	Left   bool `json:"left"`
	Bottom bool `json:"bottom"`
	Right  bool `json:"right"`
}

type Coor struct {
	X int
	Y int
}

var (
	maze   [][]Maze
	arthur []Coor
	issac  []Coor
)

func readInput(path string) error {
	file, _ := ioutil.ReadFile(path)

	jsonError := json.Unmarshal([]byte(file), &maze)
	if jsonError != nil {
		fmt.Println("Error happend json", jsonError)
	}
	return jsonError
}

func isValidMove(position Coor) bool {
	return true
}

func main() {
	readInput("./MAZE.txt")
}

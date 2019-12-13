package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Maze struct {
	x      int  `json: "x"`
	y      int  `json: "y"`
	top    bool `json: "top"`
	left   bool `json: "left"`
	bottom bool `json: "bottom"`
	right  bool `json: "right"`
}

type MazeNodes struct {
	Nodes []Maze
}

var (
	maze []Maze = make([]Maze, 500)
)

func readInput(path string) {
	file, _ := ioutil.ReadFile(path)

	data := MazeNodes{}

	_ = json.Unmarshal([]byte(file), &data)

	for i := 0; i < len(data.Nodes); i++ {
		fmt.Println("x=", data.Nodes[i].x)
	}
}

func main() {
	readInput("./MAZE.txt")
}

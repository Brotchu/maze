package main

import (
	"fmt"
	"log"
	"os"

	"github.wtf/Brotchu/maze"
)

func main() {
	maze, err := maze.LoadMaze(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(maze)

}

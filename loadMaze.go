package maze

import (
	"fmt"
	"image/png"
	"os"
)

func loadMaze(filename string) ([][]bool, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	imageData, err := png.Decode(f)
	if err != nil {
		return nil, err
	}
	fmt.Println(imageData)
	return nil, nil
}

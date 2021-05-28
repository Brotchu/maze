package maze

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

func LoadMaze(filename string) ([]Junction, error) {
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	imageData, err := png.Decode(f)
	if err != nil {
		return nil, err
	}

	res := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	height := imageData.Bounds().Max.Y
	width := imageData.Bounds().Max.X

	// fmt.Println(imageData.At(4, 0))
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if r, _, _, _ := imageData.At(x, y).RGBA(); r == 0 {
				res[y][x] = 1
			} else {
				res[y][x] = 0
			}
		}
	}

	j := []Junction{}

	newImg := image.NewRGBA(image.Rect(0, 0, width, height))
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if res[x][y] == 0 {
				newImg.Set(y, x, color.RGBA{R: 255, G: 255, B: 255, A: 255})
			} else {
				newImg.Set(y, x, color.RGBA{R: 0, G: 0, B: 0, A: 255})
			}
		}
	}

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if x == 0 || x == width-1 || y == 0 || y == height-1 {
				if res[x][y] == 0 {
					j = append(j, Junction{
						X: y,
						Y: x,
					})
					newImg.Set(y, x, color.RGBA{R: 255, G: 0, B: 0, A: 255})
				}
			} else {
				if res[x][y] == 0 && (res[x][y-1] == 0 || res[x][y+1] == 0) && (res[x+1][y] == 0 || res[x-1][y] == 0) {
					j = append(j, Junction{
						X: y,
						Y: x,
					})
					newImg.Set(y, x, color.RGBA{R: 255, G: 0, B: 0, A: 255})
				}
			}
		}
	}
	// newImg.Set(3, 4, color.RGBA{R: 255, G: 0, B: 0, A: 255})
	// fmt.Printf("%+v\n", j)
	for _, junc := range j {
		fmt.Printf("%+v\n", junc)
	}

	out, _ := os.Create("junctions.png")
	png.Encode(out, newImg)
	out.Close()
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {

		}
	}

	return nil, nil
}

type Junction struct {
	X int
	Y int
}

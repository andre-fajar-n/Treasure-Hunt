package main

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	clearPathCoordinate = []string{
		"1,1", "2,1", "3,1", "4,1", "5,1", "6,1",
		"1,2", "5,2", "6,2",
		"1,3", "2,3", "3,3", "5,3",
		"3,4", "4,4", "5,4", "6,4",
	}
	startingPositionCoordinate = "1,4"
)

type gridModel struct {
	template    [][]string
	locTreasure string
}

func main() {
	// looping grid show
	for _, v := range clearPathCoordinate {
		template := createTemplate()
		grid := gridModel{
			template:    template,
			locTreasure: v,
		}

		grid.show()
	}
}

func createTemplate() [][]string {
	output := [][]string{}
	for y := 0; y < 6; y++ {
		temp := []string{}
		for x := 0; x < 8; x++ {
			coordinate := fmt.Sprintf("%d,%d", x, y)
			if coordinate == startingPositionCoordinate {
				temp = append(temp, "X")
			} else if strings.Contains(strings.Join(clearPathCoordinate, ";"), coordinate) {
				temp = append(temp, ".")
			} else {
				temp = append(temp, "#")
			}
		}
		output = append(output, temp)
	}

	return output
}

func (g *gridModel) show() {
	// replace clear path with treasure
	if g.locTreasure != "" {
		x, y := coordinateStringToInt(g.locTreasure)
		g.template[y][x] = "$"
		fmt.Println("Location Treasure:", g.locTreasure)
	}

	// show the grid
	var temp []string
	for _, v := range g.template {
		temp = append(temp, strings.Join(v, " "))
	}
	fmt.Println(strings.Join(temp, "\n"))
	fmt.Println("\n===================================\n")
}

func coordinateStringToInt(coordinate string) (int, int) {
	xByte, yByte := coordinate[0], coordinate[len(coordinate)-1]
	x, _ := strconv.Atoi(string([]byte{xByte}))
	y, _ := strconv.Atoi(string([]byte{yByte}))

	return x, y
}

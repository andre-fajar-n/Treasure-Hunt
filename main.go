package main

import (
	"fmt"
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

func main() {
	templateGrid()
}

func templateGrid() {
	for y := 0; y < 6; y++ {
		for x := 0; x < 8; x++ {
			coordinate := fmt.Sprintf("%d,%d", x, y)
			if coordinate == startingPositionCoordinate {
				fmt.Print("X")
			} else if strings.Contains(strings.Join(clearPathCoordinate, ";"), coordinate) {
				fmt.Print(".")
			} else {
				fmt.Print("#")
			}
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

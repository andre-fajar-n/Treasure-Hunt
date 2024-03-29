package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	width  = 6
	length = 8
)

var (
	// x,y
	clearPathCoordinate = []string{
		"1,1", "2,1", "3,1", "4,1", "5,1", "6,1",
		"1,2", "5,2", "6,2",
		"1,3", "2,3", "3,3", "5,3",
		"3,4", "4,4", "5,4", "6,4",
	}
	startingPositionCoordinate = "1,4"
)

type navigation struct {
	up            int
	right         int
	down          int
	validTreasure string
}

func main() {
	TreasureHunt()
}

func TreasureHunt() {
	nav := findRoute()

	for _, v := range nav {
		createGrid(v)
	}
}

func createGrid(nav navigation) {
	var xTrs, yTrs int
	if nav.validTreasure != "" {
		xTrs, yTrs = coordinateStringToInt(nav.validTreasure)
		fmt.Println("#######################################")
		fmt.Println("Location Treasure:", nav.validTreasure)
		fmt.Println(fmt.Sprintf("Up/North %d step(s)", nav.up))
		fmt.Println(fmt.Sprintf("Right/East %d step(s)", nav.right))
		fmt.Println(fmt.Sprintf("Down/South %d step(s)", nav.down))
	}

	output := []string{}
	for y := 0; y < 6; y++ {
		temp := []string{}
		for x := 0; x < 8; x++ {
			coordinate := fmt.Sprintf("%d,%d", x, y)
			if coordinate == startingPositionCoordinate {
				temp = append(temp, "X")
			} else if strings.Contains(strings.Join(clearPathCoordinate, ";"), coordinate) {
				// replace clear path with treasure
				if x == xTrs && y == yTrs {
					temp = append(temp, "$")
				} else {
					temp = append(temp, ".")
				}
			} else {
				temp = append(temp, "#")
			}
		}
		output = append(output, strings.Join(temp, " "))
	}

	showGrid := strings.Join(output, "\n")
	fmt.Println(showGrid)
	fmt.Println()
}

func coordinateStringToInt(coordinate string) (int, int) {
	xByte, yByte := coordinate[0], coordinate[len(coordinate)-1]
	x, _ := strconv.Atoi(string([]byte{xByte}))
	y, _ := strconv.Atoi(string([]byte{yByte}))

	return x, y
}

func findRoute() []navigation {
	var steps [width - 2][length - 2][width - 2]bool
	steps[0][0][0] = true
	xStart, yStart := coordinateStringToInt(startingPositionCoordinate)
	nav := make([]navigation, 0)
	for down := 0; down < width-2; down++ {
		for right := 0; right < length-2; right++ {
			for up := 0; up < width-2; up++ {
				// skip the loop with condition invalid step
				if right == 0 && down > 0 {
					continue
				}
				y := yStart - up + down
				x := xStart + right

				// check if coordinate is exist in clear paths
				if strings.Contains(strings.Join(clearPathCoordinate, ";"), fmt.Sprintf("%d,%d", x, y)) {
					var beforeUp, beforeRight, beforeDown int
					// fix the condition step and step before is valid
					if (up > 0 && right == 0 && down == 0) || (up == 0 && right > 0 && down == 0) || (up == 0 && right == 0 && down > 0) {
						beforeUp, beforeRight, beforeDown = 0, 0, 0
						if up > 1 {
							beforeUp = up - 1
						} else if right > 1 {
							beforeRight = right - 1
						} else if down > 1 {
							beforeDown = down - 1
						}
					}
					if up > 0 && right > 0 && down == 0 {
						beforeUp, beforeRight, beforeDown = up, right-1, 0
						if right == 1 {
							beforeRight = 0
						}
					}
					if up == 0 && right > 0 && down > 0 {
						beforeUp, beforeRight, beforeDown = 0, right, down-1
						if down == 1 {
							beforeDown = 0
						}
					}
					if up == 0 && right == 0 && down > 0 {
						beforeUp, beforeRight, beforeDown = 0, 0, down-1
					}
					if up > 0 && right > 0 && down > 0 {
						beforeUp, beforeRight, beforeDown = up, right, down-1
						if right > 1 && down == 1 {
							beforeRight, beforeDown = right, 0
						}
					}

					// check if step before is valid
					// it means, the previous step have been passed
					if steps[beforeUp][beforeRight][beforeDown] {
						steps[up][right][down] = true
						yBefore := yStart - beforeUp + beforeDown
						xBefore := xStart + beforeRight

						if strings.Contains(strings.Join(clearPathCoordinate, ";"), fmt.Sprintf("%d,%d", xBefore, yBefore)) || fmt.Sprintf("%d,%d", xBefore, yBefore) == startingPositionCoordinate {
							nav = append(nav, navigation{
								up:            up,
								right:         right,
								down:          down,
								validTreasure: fmt.Sprintf("%d,%d", x, y),
							})
						}
					}
				}
			}
		}
	}

	return nav
}

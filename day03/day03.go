package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

type vector struct {
	x int
	y int
}

type tobogganMap struct {
	width  int
	height int
	data   []byte

	wraps    int
	finished bool
	cursor   vector
}

func newTobogganMap(data []byte) *tobogganMap {
	width := 1
	for ; width < len(data); width++ {
		if data[width] == byte('\n') {
			break
		}
	}
	// +1 to width to account for \n, and +1 to len for missing last newline
	height := (len(data) + 1) / (width + 1)

	return &tobogganMap{width: width, height: height, data: data, cursor: vector{0, 0}}
}

// Advances cursor (which wraps), returns current position
func (t *tobogganMap) moveBy(v vector) vector {
	xMove := v.x + t.cursor.x
	yMove := v.y + t.cursor.y

	// Need to check if we are "past" the bounds
	if xMove >= t.width {
		xMove = xMove % t.width
	}
	t.cursor.x = xMove
	t.cursor.y = yMove
	if t.cursor.y >= t.height {
		t.finished = true
	}
	return t.cursor
}

func (t *tobogganMap) valueAt(position vector) string {
	index := (position.y*(t.width+1) + position.x)
	// fmt.Printf("pos: %+v, wid: %d, ind: %d\n", position, t.width, index)
	if index > len(t.data) {
		return "?"
	}
	val := t.data[index]
	// fmt.Printf("[%d, %d] =>%s\n", position.x, position.y, string(val))
	return string(val)
}

func main() {
	fmt.Println("Day3")

	var fileName string
	flag.StringVar(&fileName, "input", "inputs/day03", "Filename to read")
	flag.Parse()

	// We can treat the 1d array as a 2d array by "wrapping" every width
	mapAsBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	currMap := newTobogganMap(mapAsBytes)
	fmt.Printf("Len: %d, Height: %d, Width: %d\n", len(mapAsBytes), currMap.height, currMap.width)

	trees := 1
	moves := []vector{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}

	for _, move := range moves {
		t := plotCourse(currMap, move)
		trees = trees * t
		currMap.cursor = vector{0, 0}
		currMap.finished = false
	}

	fmt.Println("Trees", trees)
}

func plotCourse(currMap *tobogganMap, move vector) int {
	trees := 0

	for !currMap.finished {
		currMap.moveBy(move)
		spot := currMap.valueAt(currMap.cursor)
		if spot == "#" {
			trees++
		}
	}
	return trees
}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Grid [][]bool

const (
	width     = 50
	height    = 50
	sleepTime = 100
	escp      = "\033c\x0c"
	white     = "\xE2\xAC\x9C"
	red       = "\xF0\x9F\x9F\xA5"
)

func InitializeGrid() Grid {
	w := make(Grid, height)
	for i := range w {
		w[i] = make([]bool, width)
	}
	return w
}

func (g Grid) Seed() {
	for _, row := range g {
		for j := range row {
			if rand.Intn(5) == 1 {
				row[j] = true
			}
		}
	}
}

func (g Grid) PrintGrid() {
	for _, row := range g {
		for cell := range row {
			if cell == 1 {
				fmt.Print(red)
				continue
			}
			fmt.Print(white)
		}
	}
}

func (g Grid) Alive(x, y int) bool {
	return g[y][x]
}

func (g Grid) Neighbours(x, y int) int {
	var neighbours int

	for i := y - 1; i <= y+1; i++ {
		for j := x - 1; j <= x+1; j++ {
			if i == y && j == x {
				continue
			}
			if g.Alive(j, i) {
				neighbours++
			}
		}
	}
	return neighbours
}

func (g Grid) Next(x, y int) bool {
	n := g.Neighbours(x, y)
	alive := g.Alive(x, y)
	if n < 4 && n > 1 && alive {
		return true
	} else if n == 3 && !alive {
		return true
	} else {
		return false
	}
}

func NextGen(a, b Grid) {
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			b[i][j] = a.Next(j, i)
		}
	}
}

func main() {
	fmt.Println(escp)
	currentTime := time.Now().UTC().UnixNano()
	rand.Seed(currentTime)
	newWorld := InitializeGrid()
	nextWorld := InitializeGrid()
	newWorld.Seed()
	for {
		newWorld.PrintGrid()
		NextGen(newWorld, nextWorld)
		newWorld, nextWorld = nextWorld, newWorld
		time.Sleep(sleepTime * time.Millisecond)
		fmt.Println(escp)
	}
}

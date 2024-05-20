package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	
	"log"
	"image/color"
	"math/rand"
	"time"
)

type Grid [][]bool

const (
	WIDTH     = 100
	HEIGHT    = 100
	SCALE = 2
)

var (
	black color.RGBA=color.RGBA{255,255,255,255}
	red color.RGBA=color.RGBA{255,0,0,255}
	screen *ebiten.Image
	count int = 0
	newWorld = InitializeGrid()
	nextWorld = InitializeGrid()
)

func InitializeGrid() Grid {
	w := make(Grid, HEIGHT)
	for i := range w {
		w[i] = make([]bool, HEIGHT)
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
		return true //overcrowding or underpopulated
	} else if n == 3 && !alive {
		return true //Reproduction
	} else {
		return false //Survival
	}
}

func NextGen(a, b Grid) {
	for i := 0; i < HEIGHT; i++ {
		for j := 0; j < WIDTH; j++ {
			b[i][j] = a.Next(j, i)
		}
	}
}

func Frame (screen *ebiten.Image) error {
	//Update state
	NextGen(newWorld, nextWorld)
	//Render state
	newWorld.Render(screen)

	return nil
}

func (g Grid) Render (screen *ebiten.Image){
	screen.Fill(black)
	for i :=0 ; i < HEIGHT; i++ {
		for j:= 0; j < WIDTH; j++ {
			if g[i][j] {
			}
		}

	}
}

func main() {
	currentTime := time.Now().UTC().UnixNano()
	rand.Seed(currentTime)
	newWorld.Seed()

	if err := ebiten.Run(grid, WIDTH*scale, HEIGHT*scale, scale, "Conway's Game of Life"); err != nil {
		log.Fatal(err)
	}	
}

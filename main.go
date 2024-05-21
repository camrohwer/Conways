package main

import (
	"github.com/hajimehoshi/ebiten"
	
	"log"
	"image/color"
	"math/rand"
	"time"
)

const (
	WIDTH = 480
	HEIGHT = 480
	SCALE = 2
	TILESIZE = 16
)

var (
	black color.RGBA=color.RGBA{255,255,255,255}
	red color.RGBA=color.RGBA{255,0,0,255}
	screen *ebiten.Image
)

type Grid [][]bool
type Game struct{
	world Grid
	nextWorld Grid
}

func NewGame() *Game {
    world := ClearGrid()
	nextWorld := ClearGrid()

    return &Game{
        world: world,
        nextWorld: nextWorld,
    }
}

func ClearGrid() Grid {
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
	if x < 0 || x >= WIDTH || y < 0 || y >= HEIGHT {
        return false
    }
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
		return true //Overcrowding or Underpopulated
	} else if n == 3 && !alive {
		return true //Reproduction
	} else {
		return false //Survival
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
    return HEIGHT, WIDTH
}

func (g *Game) Update(*ebiten.Image) error {
	for i :=0 ; i < HEIGHT; i++ {
		for j:= 0; j < WIDTH; j++ {
			if g.world.Next(j, i) {
				g.nextWorld[i][j] = true
			}
		}
	}
	g.world = g.nextWorld
	g.nextWorld = ClearGrid()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for y, row := range g.world {
		for x, value := range row {
			var pixelColor color.RGBA
			if value {
				pixelColor = red
			} else {
				pixelColor = black
			}
			for i := 0; i < SCALE; i++ {
				for j := 0; j < SCALE; j++ {
					screen.Set(x*SCALE+i, y*SCALE+j, pixelColor)
				}
			}
		}
	}
}

func main() {
	currentTime := time.Now().UTC().UnixNano()
	rand.Seed(currentTime)

	game := NewGame()
	game.world.Seed()

    ebiten.SetWindowSize(WIDTH*SCALE, HEIGHT*SCALE)
    ebiten.SetWindowTitle("Conway's Game of Life")

    if err := ebiten.RunGame(game); err != nil {
        log.Fatal(err)
    }
}

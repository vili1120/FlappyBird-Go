package game

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {}

func (g *Game) Update() error {
  return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
  ebitenutil.DebugPrint(screen, "test: Works!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
  return 320, 240
}

func Run(g *Game) {
  ebiten.SetWindowSize(640, 480)
  ebiten.SetWindowTitle("Flappy Bird - Go Edition")
  if err := ebiten.RunGame(g); err != nil {
    log.Fatal(err)
  }
}

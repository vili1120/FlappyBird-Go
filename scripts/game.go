package game

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/vili1120/FlappyBird-Go.git/scripts/Bird"
)

type Game struct {
  Bird bird.Bird
}

func (g *Game) Update() error {
  g.Bird.Update()
  return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
  g.Bird.Draw(screen)
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

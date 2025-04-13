package game

import (
	"image/color"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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
  screen.Fill(color.RGBA{0,255,255,100})
  g.Bird.Draw(screen)
  if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
    os.Exit(0)
  }
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
  return outsideWidth, outsideHeight
}

func Run(g *Game) {
  ebiten.SetWindowSize(480, 640)
  ebiten.SetWindowTitle("Flappy Bird - Go Edition")
  if err := ebiten.RunGame(g); err != nil {
    log.Fatal(err)
  }
}

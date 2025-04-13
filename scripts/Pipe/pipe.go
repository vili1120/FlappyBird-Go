package pipe

import (
	"github.com/hajimehoshi/ebiten/v2"
	sprite "github.com/vili1120/FlappyBird-Go.git/scripts/Sprites"
)

func GeneratePipes(amount int, img *ebiten.Image) []Pipe {
  var pipes []Pipe

  for range amount {
    p := Pipe{
      Sprite: sprite.Sprite{
        Img: img,
        PosX: 320,
        PosY: 120,
      },
      VelX: 8,
    }
    pipes = append(pipes, p)
  }

  return pipes
}

type Pipe struct {
  sprite.Sprite
  VelX float64
}

func (p *Pipe) Update() error {
  p.PosX -= p.VelX
  return nil
}

func (p *Pipe) Draw(screen *ebiten.Image) {
  opts := &ebiten.DrawImageOptions{}
  opts.GeoM.Translate(p.PosX, p.PosY)
  screen.DrawImage(
    p.Img,
    opts,
  )
}

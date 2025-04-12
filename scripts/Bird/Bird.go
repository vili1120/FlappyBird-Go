package bird

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	sprite "github.com/vili1120/FlappyBird-Go.git/scripts/Sprites"
)

type Bird struct {
  sprite.Sprite
  Score int
}

func (b *Bird) Flap(screen *ebiten.Image) {
  op := &ebiten.DrawImageOptions{}
  op.GeoM.Translate(b.PosX, b.PosY)
  screen.DrawImage(
    b.Img.SubImage(
      image.Rect(16,0,32,16),
    ).(*ebiten.Image),
    op,
  )
}

func (b *Bird) Idle(screen *ebiten.Image) {
  op := &ebiten.DrawImageOptions{}
  op.GeoM.Translate(b.PosX, b.PosY)
  screen.DrawImage(
    b.Img.SubImage(
      image.Rect(0,0,16,16),
    ).(*ebiten.Image),
    op,
  )
}

func (b *Bird) Update(screen *ebiten.Image) error {
  if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
    b.Flap(screen)
    b.PosY -= 5
  } else {
    b.Idle(screen)
    b.PosY += 0.1
  }
  return nil
}

package bird

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	sprite "github.com/vili1120/FlappyBird-Go.git/scripts/Sprites"
)

type Bird struct {
  sprite.Sprite
  Gravity float64
  VelY float64
  Score int
}

func (b *Bird) Draw(screen *ebiten.Image) {
  op := &ebiten.DrawImageOptions{}
  rotation := b.VelY * 0.05
  op.GeoM.Rotate(rotation)
  op.GeoM.Translate(b.PosX, b.PosY)
  screen.DrawImage(
    b.Img.SubImage(
      image.Rect(b.TopX,b.TopY,b.BotX,b.BotY),
    ).(*ebiten.Image),
    op,
  )
}

func (b *Bird) Jump() {
  b.VelY = -8
}

func (b *Bird) Flap() {
  b.TopX = 16
  b.TopY = 0
  b.BotX = 32
  b.BotY = 16
}

func (b *Bird) Fall() {
  b.TopX = 0
  b.TopY = 0
  b.BotX = 16
  b.BotY = 16
}

func (b *Bird) Update() error {
  b.VelY += b.Gravity
  b.PosY += b.VelY
  if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
    b.Flap()
    b.Jump()
  } else {
    b.Fall()
  }
  return nil
}

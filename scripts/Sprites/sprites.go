package sprite

import "github.com/hajimehoshi/ebiten/v2"

type Sprite struct {
  PosX, PosY float64
  Img *ebiten.Image
}

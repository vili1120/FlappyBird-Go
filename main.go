package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/vili1120/FlappyBird-Go.git/scripts"
	bird "github.com/vili1120/FlappyBird-Go.git/scripts/Bird"
	sprite "github.com/vili1120/FlappyBird-Go.git/scripts/Sprites"
)

func main() {
  playerImg, _, err := ebitenutil.NewImageFromFile("assets/Bird/bird.png")
  if err != nil {
    log.Fatal(err)
  }

  g := game.Game{
    Bird: bird.Bird{
      Sprite: sprite.Sprite{
        PosX: 100,
        PosY: 120,
        Img: playerImg,
      },
    },
  }
  game.Run(&g)
}

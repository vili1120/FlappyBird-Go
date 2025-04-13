package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/vili1120/FlappyBird-Go.git/scripts"
	bird "github.com/vili1120/FlappyBird-Go.git/scripts/Bird"
	pipe "github.com/vili1120/FlappyBird-Go.git/scripts/Pipe"
	sprite "github.com/vili1120/FlappyBird-Go.git/scripts/Sprites"
)

func main() {
  playerImg, _, err := ebitenutil.NewImageFromFile("assets/Bird/bird.png")
  if err != nil {
    log.Fatal(err)
  }
  pipeImg, _, err := ebitenutil.NewImageFromFile("assets/Pipe/pipe.png")
  if err != nil {
    log.Fatal(err)
  }
  pipes := pipe.GeneratePipes(1, pipeImg)

  g := game.Game{
    Bird: bird.Bird{
      Sprite: sprite.Sprite{
        PosX: 100,
        PosY: 120,
        Img: playerImg,
      },
      Gravity: 0.4,
      VelY: 0,
    },
    Pipes: pipes,
  }
  game.Run(&g)
}

package main

import (
	"github.com/CyberTea0X/ebiui"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	HelloNode ebiui.UiNode
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(s *ebiten.Image) {
	ebiui.RenderNode(s, g.HelloNode)
}

func (g *Game) Layout(screenWidth, screenHeight int) (int, int) {
	return 128, 128
}

func main() {
	ebiten.SetWindowSize(640, 640)
	node := &ebiui.Node{}
	g := &Game{node}
	node.Style = ebiui.Style{
		Left:   ebiui.Px(0),
		Top:    ebiui.Px(320),
		Width:  ebiui.Px(320),
		Height: ebiui.Px(320),
	}
	node.Text = "Hello, ebiui"
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}

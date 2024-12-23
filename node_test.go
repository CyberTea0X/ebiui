package ebiui

import (
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
)

func TestNode(t *testing.T) {
	node := &Node{}
	node.Style.Left = Px(0)
	node.Style.Top = Px(320)
	node.Text = "t"
	image := ebiten.NewImage(32, 32)
	RenderNode(image, node)
}

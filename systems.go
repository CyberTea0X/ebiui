package ebiui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func RenderNode(s *ebiten.Image, node UiNode) {
	rect := node.GetRect()
	vector.DrawFilledRect(
		s,
		float32(rect.Min.X),
		float32(rect.Min.Y),
		float32(rect.Dx()),
		float32(rect.Dy()),
		color.Transparent,
		false,
	)
	text := node.GetText()
	fontSize := 14
	if text != "" {
		ebitenutil.DebugPrintAt(s, text, rect.Min.X, rect.Min.Y+fontSize)
	}
	for _, child := range node.GetChildren() {
		RenderNode(s, child)
	}
}

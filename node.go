package ebiui

import (
	"image"
)

type UiNode interface {
	GetRect() image.Rectangle
	SetParent(node UiNode)
	GetChildren() []UiNode
	GetParent() UiNode
	CalcValue(val StyleValue) int
	AppendChildren(nodes ...UiNode)
	GetText() string
}

func Position(u UiNode) image.Point {
	rect := u.GetRect()
	return image.Point{rect.Min.X, rect.Min.Y}
}

type Node struct {
	Style    Style
	parent   UiNode
	children []UiNode
	Text     string
}

func (n Node) GetText() string {
	return n.Text
}

func (n Node) Position() image.Point {
	return Position(&n)
}

func (n Node) CalcValue(sVal StyleValue) int {
	value := n.Style.getValue(sVal)
	switch {
	case value.Px != nil:
		return *value.Px
	case value.Percent != nil:
		if n.parent == nil {
			return 0
		}
		return (*value.Percent * n.parent.CalcValue(sVal)) / 100
	}
	return 0
}

func (n Node) GetRect() image.Rectangle {
	var x, y, width, height int
	if n.Style.Position == "absolute" {
		x = n.CalcValue(StyleLeft)
		y = n.CalcValue(StyleTop)
	}
	width = n.CalcValue(StyleWidth)
	height = n.CalcValue(StyleHeight)
	return image.Rect(x, y, width, height)
}

func (n *Node) AppendChildren(nodes ...UiNode) {
	for _, node := range nodes {
		node.SetParent(n)
	}
	n.children = append(n.children, nodes...)
}

func (n *Node) GetChildren() []UiNode {
	return n.children
}

func (n *Node) GetParent() UiNode {
	return n.parent
}

func (n *Node) SetParent(node UiNode) {
	n.parent = node
}

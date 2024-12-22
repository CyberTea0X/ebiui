package ebiui

type TextBox struct {
	Node
	Text string
}

type Button struct {
	Node
	Text      string
	onClicked func(b *Button)
}

type Div Node

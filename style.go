package ebiui

type Style struct {
	Width  *Value
	Height *Value
	// Example: 3px solid green
	Border *Border
	// Example: auto
	Margin          string
	Position        string
	Left            *Value
	Top             *Value
	BackgroundColor string
}

type PositionType int

const (
	PositionAbsolute PositionType = iota
)

type StyleValue int

const (
	StyleWidth StyleValue = iota
	StyleHeight
	StyleLeft
	StyleTop
)

func (s Style) getValue(val StyleValue) *Value {
	switch val {
	case StyleWidth:
		return s.Width
	case StyleHeight:
		return s.Height
	case StyleLeft:
		return s.Left
	case StyleTop:
		return s.Top
	}
	return nil
}

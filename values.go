package ebiui

import "image/color"

type Value struct {
	Px      *int
	Percent *int
}

type Border struct {
	Width *Value
	Color color.Color
}

func Px(val int) *Value {
	return &Value{Px: &val}
}

func Percent(val int) *Value {
	return &Value{Percent: &val}
}

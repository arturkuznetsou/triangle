package main
import (
	"image/color"
)

func ParseHex(hex string) color.RGBA {
	switch len(hex) {
	case 7:
	case 4:
	default:
		return color.RGBA{255, 255, 255}
	}
}

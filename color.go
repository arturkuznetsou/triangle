package main
import (
	"image/color"
	"strings"
	"strconv"
)

func ParseHex(hex string) color.RGBA {
	hex = strings.ToLower(hex)
	var r, g, b int64
	r, g, b = 255, 255, 255

	switch len(hex) {
	case 7:
		r, _ = strconv.ParseInt(string(hex[1:3]), 16, 64)
		g, _ = strconv.ParseInt(string(hex[3:5]), 16, 64)
		b, _ = strconv.ParseInt(string(hex[5:8]), 16, 64)
	case 4:
		r, _ = strconv.ParseInt(string(hex[1]), 16, 64)
		g, _ = strconv.ParseInt(string(hex[2]), 16, 64)
		b, _ = strconv.ParseInt(string(hex[3]), 16, 64)
		r *= 16
		g *= 16
		b *= 16
	default:
		println("Invalid color. Using black.")
	}
	return color.RGBA{uint8(r), uint8(g), uint8(b), 0xff}
}

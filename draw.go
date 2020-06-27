package main
import (
	"fmt"
)
/*
 * Draw sierpinski triangle
 */
func draw () {
	fmt.Printf("Generating sierpinski triangle with depth %d...\n", depth)
	hline(1, x * 8 - 1, x * 8)
	dline(1, x * 8, x * 8)
	dline(x * 4, 1, -x * 8)
	triangulate(1, x * 8, x * 8)
}








/*
 * Recursive function for splitting a triangle
 * into four smaller ones
 */
func triangulate (x1 int, x2 int, y int) {
	xma := x2 - x1 + 1
	if xma > 9 {
		xa := (x1 + x2) / 2
		hline(x1 + xma / 4, x2 - xma / 4, y - xma / 2)
		dline(x1 + xma / 4 - 1, y - xma / 2 + 1, -xma / 2)
		dline(x2 - xma / 2 + 1, y, xma / 2)
		triangulate(x1, x1 + (x2 - x1) / 2, y)
		triangulate(xa + 1, x2, y)
		triangulate(x1 + xma / 4, x2 - xma / 4, y - xma / 2)
	}
}


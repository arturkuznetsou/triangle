package main


/*
 * Function for drawing horizontal lines
 */
func hline (x1 int, x2 int, y int) {
	for n := x1; n < x2; n++ {
			img.Set(n, y, col)
	}
}

/*
 * Function for drawing diagonal lines.
 * Starts at x, y and moves diagonaly
 * for each integer in height
 * Always goes to the right.
 */
func dline (x int, y int, height int) {
	add := 1
	n := 0
	height += height % 2
	if height < 0 {
		add = -1
	}
	for n != height {
		img.Set(x, y - n, col)
		img.Set(x, y - n - add, col)
		x += 1;
		n += 2 * add
	}
}


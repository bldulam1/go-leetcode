package maxPointsOnALine

type Point struct {
	x int
	y int
}

type Line struct {
	A int
	B int
	C int
}

func (r Line) isCollinear(p Point) bool {
	return r.A*p.x+r.B*p.y+r.C == 0
}

func getLine(p1 Point, p2 Point) Line {
	A, B := p1.y-p2.y, -(p1.x - p2.x)
	if A < 0 {
		A, B = -A, -B
	}

	return Line{A, B, -(A*p1.x + B*p1.y)}
}

func MaxPoints(points [][]int) int {
	pointsLen := len(points)
	if pointsLen < 3 {
		return pointsLen
	}

	maxLineLen := 0
	maxPointsLen := 0
	lines := make(map[Line]bool)
	collinearPoints := make(map[Point]int)
	for i := 0; i < pointsLen; i++ {
		for j := i + 1; j < pointsLen; j++ {
			p1 := Point{points[i][0], points[i][1]}
			p2 := Point{points[j][0], points[j][1]}
			if p1 != p2 {
				line := getLine(p1, p2)
				if !lines[line] {
					numPoints := 0
					for _, point := range points {
						if line.isCollinear(Point{point[0], point[1]}) {
							numPoints++
						}
					}

					if numPoints > maxLineLen {
						maxLineLen = numPoints
					}
				}
			} else {
				collinearPoints[p1]++
				if maxPointsLen < collinearPoints[p1] {
					maxPointsLen = collinearPoints[p1]
				}
			}
		}
	}

	if maxPointsLen > maxLineLen {
		return maxPointsLen
	}

	return maxLineLen
}

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

func getPoint(point []int) Point {
	return Point{point[0], point[1]}
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
	maxPointsOnLine, maxPointsOnPoint := 0, 0

	coPoint := make(map[Point]int)
	coLine := make(map[Line]bool)

	for i := 0; i < pointsLen; i++ {
		for j := i + 1; j < pointsLen; j++ {
			p1 := getPoint(points[i])
			p2 := getPoint(points[j])

			if p1 == p2 {
				coPoint[p1]++

				if maxPointsOnPoint < coPoint[p1] {
					maxPointsOnPoint = coPoint[p1]
				}
			} else {
				line := getLine(p1, p2)
				if !coLine[line] {
					coLine[line] = true

					numPoints := 0
					for k := range points {
						p3 := getPoint(points[k])
						if line.isCollinear(p3) {
							numPoints++
						}
						if numPoints+pointsLen-k <= maxPointsOnLine {
							break
						}
					}

					if numPoints > maxPointsOnLine {
						maxPointsOnLine = numPoints
					}
				}
			}
		}
	}
	if maxPointsOnPoint > maxPointsOnLine {
		return maxPointsOnPoint
	}

	return maxPointsOnLine
}

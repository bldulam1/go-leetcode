package maxPointsOnALine

import (
	"fmt"
	"math"
)

type Point struct {
	x int
	y int
}

type Line struct {
	p1         Point
	p2         Point
	slope      float64
	yIntercept float64
}

func (r *Line) getSlope() float64 {
	if r.slope == 0 && r.p1.y != r.p2.y {
		deltaY := float64(r.p2.y - r.p1.y)
		deltaX := float64(r.p2.x - r.p1.x)

		if deltaY != 0 {
			if deltaX == 0 {
				return math.MaxFloat64
			}
			r.slope = deltaY / deltaX
		}
	}

	return r.slope
}

func (r *Line) getYIntercept() float64 {
	if r.yIntercept == 0 {
		r.yIntercept = float64(r.p1.y) - r.getSlope()*float64(r.p1.x)
	}

	return r.yIntercept
}

func MaxPointsOnALine(points [][]int) int {
	pointsLen := len(points)

	if pointsLen < 3 {
		return pointsLen
	}
	maxLen := 2

	pointsMap := make(map[int]map[int]int)
	linesMap := make(map[float64]map[float64][][]int)

	for i := 0; i < pointsLen; i++ {
		p1 := Point{points[i][0], points[i][1]}
		if pointsMap[p1.x] == nil {
			pointsMap[p1.x] = make(map[int]int)
		}
		pointsMap[p1.x][p1.y]++

		for j := i + 1; j < pointsLen; j++ {
			p2 := Point{points[j][0], points[j][1]}

			if !(p1.x == p2.x && p1.y == p2.y) {
				l := Line{p1: p1, p2: p2}
				slope, yInt := l.getSlope(), l.getYIntercept()

				if slope == -1 {
					fmt.Println(l)
				}

				if linesMap[slope] == nil {
					linesMap[slope] = make(map[float64][][]int)
				}
				linesMap[slope][yInt] = append(linesMap[slope][yInt], points[j])

				if maxLen < len(linesMap[slope][yInt]) {
					maxLen = len(linesMap[slope][yInt])
				}
				//fmt.Println(i, j, l, l.getSlope(), l.getYIntercept())
			}
		}
	}

	fmt.Println(pointsMap)
	fmt.Println(linesMap)
	return maxLen
}

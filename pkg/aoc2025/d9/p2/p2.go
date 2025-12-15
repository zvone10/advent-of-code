package aoc2025d9p2

import (
	"os"
	"strconv"
	"strings"

	"github.com/zvone10/advent-of-code/internal"
	"github.com/zvone10/advent-of-code/internal/math"
)

const DAY, PART = 9, 2

var polygonLinesXIndex, polygonLinesYIndex map[int][]line
var rp RectangleInPolygon

func Run() {
	data, err := os.ReadFile("./inputs/2025/d9.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")
	coordsList := make([]coords, 0, len(lines))
	for _, line := range lines {
		n := strings.Split(line, ",")
		x, _ := strconv.Atoi(n[0])
		y, _ := strconv.Atoi(n[1])
		coordsList = append(coordsList, coords{x: x, y: y})
	}

	polygonLinesXIndex = make(map[int][]line)
	polygonLinesYIndex = make(map[int][]line)
	polygonLines := make([]line, 0)
	for i := 0; i < len(coordsList); i++ {
		if coordsList[i].x == coordsList[(i+1)%len(coordsList)].x {
			polygonLinesXIndex[coordsList[i].x] = append(polygonLinesXIndex[coordsList[i].x], line{
				a: coordsList[i],
				b: coordsList[(i+1)%len(coordsList)],
			})
		}

		if coordsList[i].y == coordsList[(i+1)%len(coordsList)].y {
			polygonLinesYIndex[coordsList[i].y] = append(polygonLinesXIndex[coordsList[i].y], line{
				a: coordsList[i],
				b: coordsList[(i+1)%len(coordsList)],
			})
		}

		polygonLines = append(polygonLines, line{
			a: coordsList[i],
			b: coordsList[(i+1)%len(coordsList)],
		})
	}

	rp = &rectangeInPolygonGemini{}

	maxArea := 0
	for i := 0; i < len(coordsList); i++ {
		for j := i + 1; j < len(coordsList); j++ {

			ci, cj := coordsList[i], coordsList[j]
			a := (math.Abs(ci.x-cj.x) + 1) * (math.Abs(ci.y-cj.y) + 1)

			if a < maxArea {
				continue
			}

			if rp.IsPolygonInside(ci, cj, polygonLines) {

				if a > maxArea {
					maxArea = a
				}
			}

		}
	}

	internal.PrintResult(DAY, PART, maxArea)
}

type coords struct {
	x, y int
}

type line struct {
	a, b coords
}

// isRectangeInside calculates if a rectangle with opposite points c1 and c2 is inside a polygon
// defined by a list of vertical and horizontal lines.
func isRectangeInside(c1, c2 coords, polygon []line) bool {
	minX, maxX := math.Min(c1.x, c2.x), math.Max(c1.x, c2.x)
	minY, maxY := math.Min(c1.y, c2.y), math.Max(c1.y, c2.y)

	corners := []coords{
		{x: minX, y: minY},
		{x: maxX, y: minY},
		{x: minX, y: maxY},
		{x: maxX, y: maxY},
	}

	// Check 1: All corners must be inside or on the boundary of the polygon.
	for _, corner := range corners {
		if !pointOnBoundary(corner, polygon) && !pointInPolygon(corner, polygon) {
			return false
		}
	}

	// Check 2: No edge of the rectangle should cross an edge of the polygon.
	rectLines := []line{
		{a: coords{minX, minY}, b: coords{maxX, minY}}, // bottom
		{a: coords{minX, maxY}, b: coords{maxX, maxY}}, // top
		{a: coords{minX, minY}, b: coords{minX, maxY}}, // left
		{a: coords{maxX, minY}, b: coords{maxX, maxY}}, // right
	}

	for _, rl := range rectLines {
		for _, pl := range polygon {
			if segmentsCross(rl, pl) {
				return false
			}
		}
	}

	return true
}

// pointOnLine checks if a point p lies on a line segment l.
func pointOnLine(p coords, l line) bool {
	minX, maxX := math.Min(l.a.x, l.b.x), math.Max(l.a.x, l.b.x)
	minY, maxY := math.Min(l.a.y, l.b.y), math.Max(l.a.y, l.b.y)

	// Check if p is within the bounding box of the line segment
	isInsideBoundingBox := p.x >= minX && p.x <= maxX && p.y >= minY && p.y <= maxY

	// Check for collinearity for axis-aligned lines
	isCollinear := (l.a.x == l.b.x && p.x == l.a.x) || (l.a.y == l.b.y && p.y == l.a.y)

	return isInsideBoundingBox && isCollinear
}

// pointOnBoundary checks if a point p lies on any of the edges of the polygon.
func pointOnBoundary(p coords, polygon []line) bool {
	for _, l := range polygon {
		if pointOnLine(p, l) {
			return true
		}
	}
	return false
}

// pointInPolygon checks if a point is strictly inside a polygon using the Ray Casting algorithm.
func pointInPolygon(p coords, polygon []line) bool {
	intersections := 0
	for _, l := range polygon {
		// We are interested in vertical lines for a horizontal ray cast to the right.
		if l.a.x == l.b.x && l.a.x > p.x {
			minY := math.Min(l.a.y, l.b.y)
			maxY := math.Max(l.a.y, l.b.y)
			// To handle passing through a vertex correctly, count intersection if ray
			// passes through segment's y-range [minY, maxY).
			if p.y >= minY && p.y < maxY {
				intersections++
			}
		}
	}
	return intersections%2 == 1
}

// segmentsCross checks if two axis-aligned line segments, l1 and l2, cross each other.
func segmentsCross(l1, l2 line) bool {
	var h, v line
	if l1.a.y == l1.b.y && l2.a.x == l2.b.x {
		h, v = l1, l2
	} else if l2.a.y == l2.b.y && l1.a.x == l1.b.x {
		h, v = l2, l1
	} else {
		return false // Not a horizontal-vertical pair
	}

	// A strict crossing happens if the intersection point is not an endpoint for either segment.
	return v.a.x > math.Min(h.a.x, h.b.x) && v.a.x < math.Max(h.a.x, h.b.x) &&
		h.a.y > math.Min(v.a.y, v.b.y) && h.a.y < math.Max(v.a.y, v.b.y)
}

type RectangleInPolygon interface {
	IsPolygonInside(c1, c2 coords, polygon []line) bool
}

type rectangeInPolygon struct{}

func (r *rectangeInPolygon) IsPolygonInside(ci, cj coords, polygon []line) bool {
	c1 := coords{x: math.Min(ci.x, cj.x), y: math.Max(ci.y, cj.y)}
	c2 := coords{x: math.Max(ci.x, cj.x), y: math.Min(ci.y, cj.y)}
	c3 := coords{x: c1.x, y: c2.y}
	c4 := coords{x: c2.x, y: c1.y}

	// c1-----c4
	// |      |
	// c3-----c2
	hasCrossing := false
	for i := c4.x - 1; i > c1.x && !hasCrossing; i-- {
		for _, pl := range polygonLinesXIndex[i] {
			minY, maxY := math.Min(pl.a.y, pl.b.y), math.Max(pl.a.y, pl.b.y)
			if c1.y >= minY && c1.y <= maxY {
				hasCrossing = true
				break
			}
		}
	}

	if hasCrossing {
		return false
	}

	hasCrossing = false
	for i := c2.x - 1; i > c3.x && !hasCrossing; i-- {
		for _, pl := range polygonLinesXIndex[i] {
			minY, maxY := math.Min(pl.a.y, pl.b.y), math.Max(pl.a.y, pl.b.y)
			if c2.y >= minY && c2.y <= maxY {
				hasCrossing = true
				break
			}
		}
	}

	if hasCrossing {
		return false
	}

	hasCrossing = false
	for i := c3.y + 1; i < c1.y && !hasCrossing; i++ {
		for _, pl := range polygonLinesYIndex[i] {
			minX, maxX := math.Min(pl.a.x, pl.b.x), math.Max(pl.a.x, pl.b.x)
			if c1.x >= minX && c1.x <= maxX {
				hasCrossing = true
				break
			}
		}
	}

	if hasCrossing {
		return false
	}

	hasCrossing = false
	for i := c2.y + 1; i < c4.y && !hasCrossing; i++ {
		for _, pl := range polygonLinesYIndex[i] {
			minX, maxX := math.Min(pl.a.x, pl.b.x), math.Max(pl.a.x, pl.b.x)
			if c2.x >= minX && c2.x <= maxX {
				hasCrossing = true
				break
			}
		}
	}

	if hasCrossing {
		return false
	}

	return true
}

type rectangeInPolygonGemini struct{}

func (r *rectangeInPolygonGemini) IsPolygonInside(ci, cj coords, polygon []line) bool {
	return isRectangeInside(ci, cj, polygon)
}

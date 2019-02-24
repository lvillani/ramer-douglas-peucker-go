// MIT License
//
// Copyright (c) 2019 Lorenzo Villani
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package rdp

import (
	"math"
)

// Point represents a point in space.
type Point struct {
	X float64
	Y float64
}

// RamerDouglasPeucker is a recursive implementation of the curve decimation algorithm with the same
// name.
func RamerDouglasPeucker(points []Point, epsilon float64) []Point {
	var maxDistance float64

	index := -1
	last := len(points) - 1

	for i := 1; i < last; i++ {
		distance := PerpendicularDistance(&points[i], &points[0], &points[last])
		if distance > maxDistance {
			index = i
			maxDistance = distance
		}
	}

	var ret []Point

	if maxDistance > epsilon {
		ret1 := RamerDouglasPeucker(points[0:index+1], epsilon)
		ret2 := RamerDouglasPeucker(points[index:], epsilon)

		ret = append(ret, ret1[0:len(ret1)-1]...)
		ret = append(ret, ret2...)
	} else {
		ret = []Point{points[0], points[last]}
	}

	return ret
}

// PerpendicularDistance returns the distance between the given point `p` and a line that passes
// between `p1` and `p2`.
func PerpendicularDistance(p *Point, p1 *Point, p2 *Point) float64 {
	v1 := math.Abs((p2.Y-p1.Y)*p.X - (p2.X-p1.X)*p.Y + p2.X*p1.Y - p2.Y*p1.X)
	v2 := math.Sqrt(math.Pow(p2.Y-p1.Y, 2) + math.Pow(p2.X-p1.X, 2))

	return v1 / v2
}

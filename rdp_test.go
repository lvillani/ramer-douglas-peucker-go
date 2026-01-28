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

package rdp_test

import (
	"reflect"
	"testing"

	rdp "github.com/lvillani/ramer-douglas-peucker-go"
)

func TestRamerDouglasPeucker(t *testing.T) {
	type args struct {
		points  []rdp.Point
		epsilon float64
	}
	tests := []struct {
		name string
		args args
		want []rdp.Point
	}{
		{
			name: "Sinusoid",
			args: args{
				points: []rdp.Point{
					{X: 0.00000, Y: 0.00000},
					{X: 0.11111, Y: 0.03841},
					{X: 0.22222, Y: 0.04801},
					{X: 0.33333, Y: 0.03704},
					{X: 0.44444, Y: 0.01372},
					{X: 0.55556, Y: -0.01372},
					{X: 0.66667, Y: -0.03704},
					{X: 0.77778, Y: -0.04801},
					{X: 0.88889, Y: -0.03841},
					{X: 1.00000, Y: 0.00000},
				},
				epsilon: 0.01,
			},
			want: []rdp.Point{
				{X: 0.00000, Y: 0.00000},
				{X: 0.11111, Y: 0.03841},
				{X: 0.22222, Y: 0.04801},
				{X: 0.77778, Y: -0.04801},
				{X: 0.88889, Y: -0.03841},
				{X: 1.00000, Y: 0.00000},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rdp.RamerDouglasPeucker(tt.args.points, tt.args.epsilon); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RamerDouglasPeucker() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPerpendicularDistance(t *testing.T) {
	type args struct {
		p  *rdp.Point
		p1 *rdp.Point
		p2 *rdp.Point
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Simple",
			args: args{
				p:  &rdp.Point{X: 0, Y: 1},
				p1: &rdp.Point{X: 1, Y: 1},
				p2: &rdp.Point{X: 1, Y: 3},
			},
			want: 1,
		},
		{
			name: "45deg",
			args: args{
				p:  &rdp.Point{X: 1, Y: 2},
				p1: &rdp.Point{X: 1, Y: 1},
				p2: &rdp.Point{X: 3, Y: 3},
			},
			want: 0.7071067811865475,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rdp.PerpendicularDistance(tt.args.p, tt.args.p1, tt.args.p2); got != tt.want {
				t.Errorf("PerpendicularDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

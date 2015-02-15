package geom2d

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPointDistanceToPoint(t *testing.T) {
	cases := []struct {
		P, Q     Point
		expected float64
	}{
		{Point{1, 1}, Point{2, 2}, 1.414},
		{Point{1, 2}, Point{2, 2}, 1},
		{Point{0, 0}, Point{2, 2}, 2.828},
		{Point{2, 2}, Point{0, 2}, 2},
		{Point{0, 0}, Point{0.5, 0.5}, 0.707106781},
	}

	for _, tc := range cases {
		actual := tc.P.DistanceToPoint(tc.Q)
		assert.InDelta(t, tc.expected, actual, 0.01, "%v to %v should be %f but was %f", tc.P, tc.Q, tc.expected, actual)
	}
}

func TestPointDistanceToLineSegment(t *testing.T) {
	cases := []struct {
		P        Point
		S        LineSegment
		expected float64
	}{
		{Point{1, 1}, LineSegment{Point{1, 1}, Point{2, 2}}, 0},
		{Point{1, 1}, LineSegment{Point{1, 2}, Point{2, 2}}, 1},
		{Point{1, 1}, LineSegment{Point{0, 0}, Point{2, 2}}, 0},
		{Point{1, 1}, LineSegment{Point{2, 2}, Point{0, 2}}, 1},
		{Point{1, 1}, LineSegment{Point{0, 0}, Point{0.5, 0.5}}, 0.707106781},
	}

	for _, tc := range cases {
		actual := tc.P.DistanceToLineSegment(tc.S)
		assert.InDelta(t, tc.expected, actual, 0.01, "%v to %v should be %f but was %f", tc.P, tc.S, tc.expected, actual)
	}
}

func TestPointDistanceToPolygon(t *testing.T) {
	var poly = Polygon{
		Point{1, 1},
		Point{1, 2},
		Point{2, 2},
		Point{2, 1},
		Point{1, 1},
	}
	var cases = []struct {
		Pt       Point
		expected float64
	}{
		{Point{1, 1}, 0},
		{Point{1.5, 1.5}, 0},
		{Point{0, 0}, 1.414},
		{Point{2, 2.5}, 0.5},
		{Point{1.5, 0.9}, 0.1},
	}

	for _, tc := range cases {
		actual := tc.Pt.DistanceToPolygon(poly)
		assert.InDelta(t, tc.expected, actual, 0.01, "Point %v distance to poly should be %v but was %v", tc.Pt, tc.expected, actual)
	}
}

package geom2d

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPointDotProductPoint(t *testing.T) {
	cases := []struct {
		A, B     Point
		expected float64
	}{
		{Point{1, 1}, Point{2, 2}, 4},
		{Point{1, 2}, Point{3, 4}, 11},
	}

	for _, tc := range cases {
		assert.Equal(t, tc.expected, tc.A.Dot(tc.B))
	}
}

func TestPointMinusPoint(t *testing.T) {
	cases := []struct {
		A, B     Point
		expected Point
	}{
		{Point{1, 1}, Point{2, 2}, Point{-1, -1}},
		{Point{1, 6}, Point{3, 4}, Point{-2, 2}},
	}

	for _, tc := range cases {
		assert.Equal(t, tc.expected, tc.A.Minus(tc.B))
	}
}

func TestPointPlusPoint(t *testing.T) {
	cases := []struct {
		A, B     Point
		expected Point
	}{
		{Point{1, 1}, Point{2, 2}, Point{3, 3}},
		{Point{1, 6}, Point{3, 4}, Point{4, 10}},
	}

	for _, tc := range cases {
		assert.Equal(t, tc.expected, tc.A.Plus(tc.B))
	}
}

func TestPointMultScalar(t *testing.T) {
	cases := []struct {
		A        Point
		S        float64
		expected Point
	}{
		{Point{1, 1}, 3, Point{3, 3}},
		{Point{1, 6}, 1.5, Point{1.5, 9}},
	}

	for _, tc := range cases {
		assert.Equal(t, tc.expected, tc.A.Mult(tc.S))
	}
}

func TestPointLength(t *testing.T) {
	cases := []struct {
		p        Point
		expected float64
	}{
		{Point{1, 1}, 1.414},
		{Point{3, 4}, 5},
	}

	for _, tc := range cases {
		assert.InDelta(t, tc.expected, tc.p.Length(), 0.01)
	}
}

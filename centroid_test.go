package geom2d

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArea(t *testing.T) {
	poly := Polygon{
		Point{0, 1},
		Point{1, 1},
		Point{1, 0},
		Point{0, 0},
		Point{0, 1},
	}
	a := poly.area()

	// The polygon winds clockwise so the area is negative.
	assert.Equal(t, -1.0, a)
}

func TestCentroid(t *testing.T) {
	poly := Polygon{
		Point{0, 1},
		Point{1, 1},
		Point{1, 0},
		Point{0, 0},
		Point{0, 1},
	}

	c := poly.Centroid()

	assert.Equal(t, Point{0.5, 0.5}, c)

}

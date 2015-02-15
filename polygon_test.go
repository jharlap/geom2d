package geom2d

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPolygonContainsPoint(t *testing.T) {
	var poly = Polygon{
		Point{1, 1},
		Point{1, 2},
		Point{2, 2},
		Point{2, 1},
		Point{1, 1},
	}
	var cases = []struct {
		Pt       Point
		expected bool
	}{
		{Point{1, 1}, true},
		{Point{1.5, 1.5}, true},
		{Point{0, 0}, false},
		{Point{2, 2.5}, false},
		{Point{1.5, 1}, true},
	}

	for _, tc := range cases {
		actual := poly.ContainsPoint(tc.Pt)
		assert.Equal(t, tc.expected, actual, "Point %v contained in poly should be %v but was %v", tc.Pt, tc.expected, actual)
	}
}

func TestPolygonStraddlingDatelineContainsPoint(t *testing.T) {
	t.Skip("Skipped because geom2d knows nothing of the antimeridian")

	var poly = Polygon{
		Point{149.9921875, -31.653381399663985},
		Point{149.9921875, 52.696361078274485},
		Point{-76.9921875, 52.696361078274485},
		Point{-76.9921875, -31.653381399663985},
		Point{149.9921875, -31.653381399663985},
	}
	var cases = []struct {
		Pt       Point
		expected bool
	}{
		{Point{-144.140625, 14.264383087562662}, true},
		{Point{165.234375, 23.88583769986199}, true},
		{Point{0, 0}, false},
	}

	for _, tc := range cases {
		actual := poly.ContainsPoint(tc.Pt)
		assert.Equal(t, tc.expected, actual, "Point %v contained in poly should be %v but was %v", tc.Pt, tc.expected, actual)
	}
}

func BenchmarkPolygonContainsPoint(b *testing.B) {

	b.StopTimer()

	poly := Polygon{
		Point{-80.14185190200806, 25.9588742655546},
		Point{-80.14065027236938, 25.95785174405259},
		Point{-80.14030694961548, 25.956559109817576},
		Point{-80.13964176177979, 25.955710207970235},
		Point{-80.13938426971436, 25.95476483279832},
		Point{-80.13989925384521, 25.954147440832564},
		Point{-80.14142274856567, 25.953838743635337},
		Point{-80.14320373535156, 25.954552604675325},
		Point{-80.1453709602356, 25.956153952881426},
		Point{-80.14588594436646, 25.956732748077545},
		Point{-80.14590740203857, 25.958488409427705},
		Point{-80.14549970626831, 25.959125071358685},
		Point{-80.1445984840393, 25.95954951073254},
		Point{-80.14371871948242, 25.960089704085377},
		Point{-80.1427960395813, 25.959819607718924},
		Point{-80.14185190200806, 25.9588742655546},
	}
	point := Point{-80.14, 25.96}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		poly.ContainsPoint(point)
	}

}

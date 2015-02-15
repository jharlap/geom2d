package geom2d

import "math"

// DistanceToPoint computes the distance between two points
func (p *Point) DistanceToPoint(q Point) float64 {
	v := p.Minus(q)
	return v.Length()
}

// DistanceToLineSegment computes the distance between a point and a line segment
func (p *Point) DistanceToLineSegment(s LineSegment) float64 {
	v := s.B.Minus(s.A)
	w := p.Minus(s.A)

	c1 := w.Dot(v)
	if c1 <= 0 {
		return p.DistanceToPoint(s.A)
	}

	c2 := v.Dot(v)
	if c2 <= c1 {
		return p.DistanceToPoint(s.B)
	}

	b := c1 / c2
	pb := s.A.Plus(v.Mult(b))
	return p.DistanceToPoint(pb)
}

// DistanceToPolygon computes the distance between a point and a polygon
func (p *Point) DistanceToPolygon(poly Polygon) float64 {
	if poly.ContainsPoint(*p) {
		return 0
	}

	best := math.Inf(1)
	n := len(poly) - 1
	for i := 0; i < n; i++ {
		dist := p.DistanceToLineSegment(LineSegment{poly[i], poly[i+1]})
		if dist < best {
			best = dist
		}
	}
	return best
}

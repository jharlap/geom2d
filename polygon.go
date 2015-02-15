package geom2d

// Polygon is a linear ring of points - the last point should be the first point
type Polygon []Point

func isLeft(p0, p1, p2 Point) float64 {
	return (p1.X-p0.X)*(p2.Y-p0.Y) - (p2.X-p0.X)*(p1.Y-p0.Y)
}

// ContainsPoint uses the winding number to determine if a point is contained within the polygon.
// See http://geomalgorithms.com/a03-_inclusion.html for the algorithm
func (p Polygon) ContainsPoint(pt Point) bool {
	wn := 0
	n := len(p) - 1
	for i := 0; i < n; i++ {
		pi := p[i]
		pj := p[i+1]
		if pi.Y <= pt.Y { // start y <= pt.Y
			if pj.Y > pt.Y && isLeft(pi, pj, pt) > 0 { // an upward crossing and P left of  edge
				wn++ // have  a valid up intersect
			}
		} else { // start y > pt.Y (no test needed)
			if pj.Y <= pt.Y && isLeft(pi, pj, pt) < 0 { // a downward crossing and
				wn-- // have  a valid down intersect
			}
		}
	}
	return wn != 0
}

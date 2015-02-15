package geom2d

// Area computes the area of the exterior ring of the polygon, negative if the winding order is clockwise.
func (poly Polygon) area() float64 {
	s := 0.0

	for i := 0; i < len(poly)-1; i++ {
		s += float64(poly[i].X*poly[i+1].Y - poly[i+1].X*poly[i].Y)
	}

	return 0.5 * s
}

// Centroid calculates the centroid of the exterior ring of a polygon using
// the formula at http://en.wikipedia.org/wiki/Centroid#Centroid_of_polygon
// but be careful as this applies Euclidean principles to decidedly non-
// Euclidean geometry. In other words, it will fail miserably for polygons
// near the poles, polygons that straddle the dateline, and for large polygons
// where Euclidean approximations break down. For a true centroid algorithm,
// check page 54 of
// http://www.jennessent.com/downloads/Graphics_Shapes_Online.pdf
func (poly Polygon) Centroid() Point {
	c := Point{0, 0}

	for i := 0; i < len(poly)-1; i++ {
		c.X += (poly[i].X + poly[i+1].X) * (poly[i].X*poly[i+1].Y - poly[i+1].X*poly[i].Y)
		c.Y += (poly[i].Y + poly[i+1].Y) * (poly[i].X*poly[i+1].Y - poly[i+1].X*poly[i].Y)
	}

	a := poly.area()
	c.X /= a * 6
	c.Y /= a * 6
	return c
}

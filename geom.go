package geom2d

import "math"

// Point is a coordinate in 2D space
type Point struct {
	X, Y float64
}

// LineSegment is a line segment between two points
type LineSegment struct {
	A, B Point
}

// Dot product of two points
func (p *Point) Dot(q Point) float64 {
	return p.X*q.X + p.Y*q.Y
}

// Minus computes the vector difference between two points
func (p *Point) Minus(q Point) Point {
	return Point{p.X - q.X, p.Y - q.Y}
}

// Plus adds two points
func (p *Point) Plus(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}

// Mult scales a point by a scalar
func (p *Point) Mult(m float64) Point {
	return Point{p.X * m, p.Y * m}
}

// Length computes the length of the vector (or the distance from the origin to the point)
func (p *Point) Length() float64 {
	return math.Sqrt(p.Dot(*p))
}

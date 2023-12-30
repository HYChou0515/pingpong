package game

import "math"

type Point struct {
	X float64
	Y float64
}

func (p Point) dist(b *Point) float64 {
	return math.Sqrt((p.X-b.X)*(p.X-b.X) + (p.Y-b.Y)*(p.Y-b.Y))
}

func MakePoint(x float64, y float64) *Point {
	return &Point{
		X: x,
		Y: y,
	}
}

type Rectangle struct {
	TopLeft     *Point
	BottomRight *Point
}

type Circle struct {
	Center *Point
	Radius float64
}

func (r Rectangle) GetTopRight() *Point {
	return MakePoint(r.BottomRight.X, r.TopLeft.Y)
}
func (r Rectangle) GetBottomLeft() *Point {
	return MakePoint(r.TopLeft.X, r.BottomRight.Y)
}

func MakeRectangle(x1 float64, y1 float64, x2 float64, y2 float64) *Rectangle {
	if x1 > x2 {
		x1, x2 = x2, x1
	}
	if y1 > y2 {
		y1, y2 = y2, y1
	}
	return &Rectangle{
		TopLeft:     MakePoint(x1, y1),
		BottomRight: MakePoint(x2, y2),
	}
}

func MakeCircle(x float64, y float64, r float64) *Circle {
	return &Circle{
		Center: MakePoint(x, y),
		Radius: r,
	}
}

func (r Rectangle) Area() float64 {
	return (r.BottomRight.X - r.TopLeft.X) * (r.BottomRight.Y - r.TopLeft.Y)
}

func (r Rectangle) Intersect(r2 *Rectangle) bool {
	return r.TopLeft.X < r2.BottomRight.X &&
		r.BottomRight.X > r2.TopLeft.X &&
		r.TopLeft.Y < r2.BottomRight.Y &&
		r.BottomRight.Y > r2.TopLeft.Y
}

func (r Rectangle) IntersectArea(r2 *Rectangle) float64 {
	if !r.Intersect(r2) {
		return 0
	}
	x1 := max(r.TopLeft.X, r2.TopLeft.X)
	x2 := min(r.BottomRight.X, r2.BottomRight.X)
	y1 := max(r.TopLeft.Y, r2.TopLeft.Y)
	y2 := min(r.BottomRight.Y, r2.BottomRight.Y)
	return (x2 - x1) * (y2 - y1)
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Intersect(r *Rectangle) bool {
	if c.Center.X < r.TopLeft.X-c.Radius {
		return false
	}
	if c.Center.X > r.BottomRight.X+c.Radius {
		return false
	}
	if c.Center.Y < r.TopLeft.Y-c.Radius {
		return false
	}
	if c.Center.Y > r.BottomRight.Y+c.Radius {
		return false
	}
	if (c.Center.X > r.GetTopRight().X) && (c.Center.Y < r.GetTopRight().Y) && (c.Center.dist(r.GetTopRight()) > c.Radius) {
		return false
	}
	if (c.Center.X < r.GetBottomLeft().X) && (c.Center.Y > r.GetBottomLeft().Y) && (c.Center.dist(r.GetBottomLeft()) > c.Radius) {
		return false
	}
	if (c.Center.X < r.TopLeft.X) && (c.Center.Y < r.TopLeft.Y) && (c.Center.dist(r.TopLeft) > c.Radius) {
		return false
	}
	if (c.Center.X > r.BottomRight.X) && (c.Center.Y > r.BottomRight.Y) && (c.Center.dist(r.BottomRight) > c.Radius) {
		return false
	}
	return true
}

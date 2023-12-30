package game

type Brick struct {
	Shape *Rectangle
}

type Wall struct {
	Shape *Rectangle
}

type Ball struct {
	Shape *Circle
}

type Paddle struct {
	Shape *Rectangle
}

type Board struct {
	Width  float64
	Height float64
	Bricks []*Brick
	Wall   []*Wall
	Ball   *Ball
	Paddle *Paddle
	Fps    float64
}

func ResetBoard() *Board {
	Width := float64(20)
	Height := float64(40)

	var bricks []*Brick
	for i := float64(3); i < Width-3; i++ {
		for j := float64(3); j < 10; j++ {
			bricks = append(bricks, &Brick{
				Shape: MakeRectangle(i, j, i+0.8, j+0.8),
			})
		}
	}

	var walls []*Wall
	walls = append(walls, &Wall{
		Shape: MakeRectangle(0, 0, Width, 1),
	})
	walls = append(walls, &Wall{
		Shape: MakeRectangle(0, 1, 1, Height),
	})
	walls = append(walls, &Wall{
		Shape: MakeRectangle(Width-1, 1, Width, Height),
	})
	return &Board{
		Width:  Width,
		Height: Height,
		Bricks: bricks,
		Wall:   walls,
		Ball: &Ball{
			Shape: MakeCircle(Width/2, Height-1.5, 0.5),
		},
		Paddle: &Paddle{
			Shape: MakeRectangle(Width/2-2, Height-1, Width/2+2, Height),
		},
	}
}

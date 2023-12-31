package game

import "github.com/google/uuid"

type Brick struct {
	Shape *Rectangle
	Id    string
}

func MakeBrick(x1 float64, y1 float64, x2 float64, y2 float64) *Brick {
	return &Brick{
		Shape: MakeRectangle(x1, y1, x2, y2),
		Id:    uuid.New().String(),
	}
}

type Wall struct {
	Shape *Rectangle
	Id    string
}

func MakeWall(x1 float64, y1 float64, x2 float64, y2 float64) *Wall {
	return &Wall{
		Shape: MakeRectangle(x1, y1, x2, y2),
		Id:    uuid.New().String(),
	}
}

type Ball struct {
	Shape *Circle
	Id    string
}

func MakeBall(x float64, y float64, r float64) *Ball {
	return &Ball{
		Shape: MakeCircle(x, y, r),
		Id:    uuid.New().String(),
	}
}

type Paddle struct {
	Shape *Rectangle
	Id    string
}

func MakePaddle(x1 float64, y1 float64, x2 float64, y2 float64) *Paddle {
	return &Paddle{
		Shape: MakeRectangle(x1, y1, x2, y2),
		Id:    uuid.New().String(),
	}
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
			bricks = append(bricks,
				MakeBrick(i, j, i+0.8, j+0.8),
			)
		}
	}

	var walls []*Wall
	walls = append(walls,
		MakeWall(0, 0, Width, 1))
	walls = append(walls,
		MakeWall(0, 1, 1, Height))
	walls = append(walls,
		MakeWall(Width-1, 1, Width, Height))
	return &Board{
		Width:  Width,
		Height: Height,
		Bricks: bricks,
		Wall:   walls,
		Ball:   MakeBall(Width/2, Height-1.5, 0.5),
		Paddle: MakePaddle(Width/2-2, Height-1, Width/2+2, Height),
	}
}

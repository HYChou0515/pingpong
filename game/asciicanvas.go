package game

import "math"

type AsciiCanvas struct {
	Canvas     [][]string
	OffsetX    float64
	OffsetY    float64
	CharWidth  float64
	CharHeight float64
}

func MakeAsciiCanvas(
	windowWidth float64,
	windowHeight float64,
	charWidth float64,
	charHeight float64,
	OffsetX float64,
	OffsetY float64,
) *AsciiCanvas {
	charsPerRow := math.Floor(windowWidth / charWidth)
	charsPerCol := math.Floor(windowHeight / charHeight)
	canvas := make([][]string, int(charsPerCol))
	for i := range canvas {
		canvas[i] = make([]string, int(charsPerRow))
	}
	return &AsciiCanvas{
		Canvas:     canvas,
		OffsetX:    OffsetX,
		OffsetY:    OffsetY,
		CharWidth:  charWidth,
		CharHeight: charHeight,
	}
}

func (c *AsciiCanvas) RenderRectangle(rect *Rectangle, char string) float64 {
	xStart := math.Floor((rect.TopLeft.X-c.OffsetX)/c.CharWidth) - 1
	xEnd := math.Ceil((rect.BottomRight.X-c.OffsetX)/c.CharWidth) + 1
	yStart := math.Floor((rect.TopLeft.Y-c.OffsetY)/c.CharHeight) - 1
	yEnd := math.Ceil((rect.BottomRight.Y-c.OffsetY)/c.CharHeight) + 1
	for col := xStart; col <= xEnd; col++ {
		for row := yStart; row <= yEnd; row++ {
			if col < 0 || row < 0 || col >= float64(len(c.Canvas[0])) || row >= float64(len(c.Canvas)) {
				continue
			}
			charRect := MakeRectangle(
				col*c.CharWidth+c.OffsetX,
				row*c.CharHeight+c.OffsetY,
				(col+1)*c.CharWidth+c.OffsetX,
				(row+1)*c.CharHeight+c.OffsetY,
			)
			if rect.IntersectArea(charRect) > 0 {
				c.Canvas[int(row)][int(col)] = char
			}
		}
	}
	return 0
}

func (c *AsciiCanvas) RenderCircle(circle *Circle, char string) {
	xStart := math.Floor((circle.Center.X-circle.Radius-c.OffsetX)/c.CharWidth) - 1
	xEnd := math.Ceil((circle.Center.X+circle.Radius-c.OffsetX)/c.CharWidth) + 1
	yStart := math.Floor((circle.Center.Y-circle.Radius-c.OffsetY)/c.CharHeight) - 1
	yEnd := math.Ceil((circle.Center.Y+circle.Radius-c.OffsetY)/c.CharHeight) + 1
	for col := xStart; col <= xEnd; col++ {
		for row := yStart; row <= yEnd; row++ {
			if col < 0 || row < 0 || col >= float64(len(c.Canvas[0])) || row >= float64(len(c.Canvas)) {
				continue
			}
			charRect := MakeRectangle(
				col*c.CharWidth+c.OffsetX,
				row*c.CharHeight+c.OffsetY,
				(col+1)*c.CharWidth+c.OffsetX,
				(row+1)*c.CharHeight+c.OffsetY,
			)
			smallerCircle := MakeCircle(
				circle.Center.X,
				circle.Center.Y,
				circle.Radius*0.5,
			)
			if smallerCircle.Intersect(charRect) {
				c.Canvas[int(row)][int(col)] = char
			}
		}
	}
}

func (c *AsciiCanvas) RenderBrick(brick *Brick) {
	c.RenderRectangle(brick.Shape, "B")
}

func (c *AsciiCanvas) RenderBall(ball *Ball) {
	c.RenderCircle(ball.Shape, "*")
}

func (c *AsciiCanvas) RenderPaddle(paddle *Paddle) {
	c.RenderRectangle(paddle.Shape, "=")
}

func (c *AsciiCanvas) RenderWall(wall *Wall) {
	c.RenderRectangle(wall.Shape, "W")
}

func (c *AsciiCanvas) RenderBoard(board *Board) {

	//c := MakeAsciiCanvas(
	//	board.Width+6.6,
	//	board.Height+6.6,
	//	0.4,
	//	1.2,
	//	-3.3,
	//	-3.3,
	//)

	for _, brick := range board.Bricks {
		c.RenderBrick(brick)
	}
	for _, wall := range board.Wall {
		c.RenderWall(wall)
	}
	c.RenderPaddle(board.Paddle)
	c.RenderBall(board.Ball)

	for _, row := range c.Canvas {
		for _, char := range row {
			if char == "" {
				print(" ")
			} else {
				print(char)
			}
		}
		print("\n")
	}
}

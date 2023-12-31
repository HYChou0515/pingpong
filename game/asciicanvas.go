package game

import "math"

type AsciiCanvasPlotOptions struct {
	Width      float64
	Height     float64
	CharWidth  float64
	CharHeight float64
	OffsetX    float64
	OffsetY    float64

	BrickColor  string
	WallColor   string
	BallColor   string
	PaddleColor string
}

type AsciiCanvas struct {
	Canvas [][]string

	PlotOptions *AsciiCanvasPlotOptions
}

func MakeAsciiCanvas(
	plotOptions *AsciiCanvasPlotOptions,
) *AsciiCanvas {
	charsPerRow := math.Floor(plotOptions.Width / plotOptions.CharWidth)
	charsPerCol := math.Floor(plotOptions.Height / plotOptions.CharHeight)
	canvas := make([][]string, int(charsPerCol))
	for i := range canvas {
		canvas[i] = make([]string, int(charsPerRow))
	}
	return &AsciiCanvas{
		Canvas:      canvas,
		PlotOptions: plotOptions,
	}
}

func (c *AsciiCanvas) renderRectangle(rect *Rectangle, char string) float64 {
	xStart := math.Floor((rect.TopLeft.X-c.PlotOptions.OffsetX)/c.PlotOptions.CharWidth) - 1
	xEnd := math.Ceil((rect.BottomRight.X-c.PlotOptions.OffsetX)/c.PlotOptions.CharWidth) + 1
	yStart := math.Floor((rect.TopLeft.Y-c.PlotOptions.OffsetY)/c.PlotOptions.CharHeight) - 1
	yEnd := math.Ceil((rect.BottomRight.Y-c.PlotOptions.OffsetY)/c.PlotOptions.CharHeight) + 1
	for col := xStart; col <= xEnd; col++ {
		for row := yStart; row <= yEnd; row++ {
			if col < 0 || row < 0 || col >= float64(len(c.Canvas[0])) || row >= float64(len(c.Canvas)) {
				continue
			}
			charRect := MakeRectangle(
				col*c.PlotOptions.CharWidth+c.PlotOptions.OffsetX,
				row*c.PlotOptions.CharHeight+c.PlotOptions.OffsetY,
				(col+1)*c.PlotOptions.CharWidth+c.PlotOptions.OffsetX,
				(row+1)*c.PlotOptions.CharHeight+c.PlotOptions.OffsetY,
			)
			if rect.IntersectArea(charRect) > 0 {
				c.Canvas[int(row)][int(col)] = char
			}
		}
	}
	return 0
}

func (c *AsciiCanvas) renderCircle(circle *Circle, char string) {
	xStart := math.Floor((circle.Center.X-circle.Radius-c.PlotOptions.OffsetX)/c.PlotOptions.CharWidth) - 1
	xEnd := math.Ceil((circle.Center.X+circle.Radius-c.PlotOptions.OffsetX)/c.PlotOptions.CharWidth) + 1
	yStart := math.Floor((circle.Center.Y-circle.Radius-c.PlotOptions.OffsetY)/c.PlotOptions.CharHeight) - 1
	yEnd := math.Ceil((circle.Center.Y+circle.Radius-c.PlotOptions.OffsetY)/c.PlotOptions.CharHeight) + 1
	for col := xStart; col <= xEnd; col++ {
		for row := yStart; row <= yEnd; row++ {
			if col < 0 || row < 0 || col >= float64(len(c.Canvas[0])) || row >= float64(len(c.Canvas)) {
				continue
			}
			charRect := MakeRectangle(
				col*c.PlotOptions.CharWidth+c.PlotOptions.OffsetX,
				row*c.PlotOptions.CharHeight+c.PlotOptions.OffsetY,
				(col+1)*c.PlotOptions.CharWidth+c.PlotOptions.OffsetX,
				(row+1)*c.PlotOptions.CharHeight+c.PlotOptions.OffsetY,
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
	c.renderRectangle(brick.Shape, c.PlotOptions.BrickColor)
}

func (c *AsciiCanvas) RenderBall(ball *Ball) {
	c.renderCircle(ball.Shape, c.PlotOptions.BallColor)
}

func (c *AsciiCanvas) RenderPaddle(paddle *Paddle) {
	c.renderRectangle(paddle.Shape, c.PlotOptions.PaddleColor)
}

func (c *AsciiCanvas) RenderWall(wall *Wall) {
	c.renderRectangle(wall.Shape, c.PlotOptions.WallColor)
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

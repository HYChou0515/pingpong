package game

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"time"
)

type FyneCanvas struct {
	Canvas  fyne.Window
	OffsetX float64
	OffsetY float64
}

//func (c *FyneCanvas) RenderRectangle(rect *Rectangle, char string) float64 {
//	xStart := math.Floor((rect.TopLeft.X-c.OffsetX)/c.CharWidth) - 1
//	xEnd := math.Ceil((rect.BottomRight.X-c.OffsetX)/c.CharWidth) + 1
//	yStart := math.Floor((rect.TopLeft.Y-c.OffsetY)/c.CharHeight) - 1
//	yEnd := math.Ceil((rect.BottomRight.Y-c.OffsetY)/c.CharHeight) + 1
//	for col := xStart; col <= xEnd; col++ {
//		for row := yStart; row <= yEnd; row++ {
//			if col < 0 || row < 0 || col >= float64(len(c.Canvas[0])) || row >= float64(len(c.Canvas)) {
//				continue
//			}
//			charRect := MakeRectangle(
//				col*c.CharWidth+c.OffsetX,
//				row*c.CharHeight+c.OffsetY,
//				(col+1)*c.CharWidth+c.OffsetX,
//				(row+1)*c.CharHeight+c.OffsetY,
//			)
//			if rect.IntersectArea(charRect) > 0 {
//				c.Canvas[int(row)][int(col)] = char
//			}
//		}
//	}
//	return 0
//}
//
//func (c *FyneCanvas) RenderCircle(circle *Circle, char string) {
//	xStart := math.Floor((circle.Center.X-circle.Radius-c.OffsetX)/c.CharWidth) - 1
//	xEnd := math.Ceil((circle.Center.X+circle.Radius-c.OffsetX)/c.CharWidth) + 1
//	yStart := math.Floor((circle.Center.Y-circle.Radius-c.OffsetY)/c.CharHeight) - 1
//	yEnd := math.Ceil((circle.Center.Y+circle.Radius-c.OffsetY)/c.CharHeight) + 1
//	for col := xStart; col <= xEnd; col++ {
//		for row := yStart; row <= yEnd; row++ {
//			if col < 0 || row < 0 || col >= float64(len(c.Canvas[0])) || row >= float64(len(c.Canvas)) {
//				continue
//			}
//			charRect := MakeRectangle(
//				col*c.CharWidth+c.OffsetX,
//				row*c.CharHeight+c.OffsetY,
//				(col+1)*c.CharWidth+c.OffsetX,
//				(row+1)*c.CharHeight+c.OffsetY,
//			)
//			smallerCircle := MakeCircle(
//				circle.Center.X,
//				circle.Center.Y,
//				circle.Radius*0.5,
//			)
//			if smallerCircle.Intersect(charRect) {
//				c.Canvas[int(row)][int(col)] = char
//			}
//		}
//	}
//}
//
//func (c *FyneCanvas) RenderBrick(brick *Brick) {
//	c.RenderRectangle(brick.Shape, "B")
//}
//
//func (c *FyneCanvas) RenderBall(ball *Ball) {
//	c.RenderCircle(ball.Shape, "*")
//}
//
//func (c *FyneCanvas) RenderPaddle(paddle *Paddle) {
//	c.RenderRectangle(paddle.Shape, "=")
//}
//
//func (c *FyneCanvas) RenderWall(wall *Wall) {
//	c.RenderRectangle(wall.Shape, "W")
//}
//
//func (c *FyneCanvas) RenderBoard(board *Board) {
//
//	//c := MakeFyneCanvas(
//	//	board.Width+6.6,
//	//	board.Height+6.6,
//	//	0.4,
//	//	1.2,
//	//	-3.3,
//	//	-3.3,
//	//)
//
//	for _, brick := range board.Bricks {
//		c.RenderBrick(brick)
//	}
//	for _, wall := range board.Wall {
//		c.RenderWall(wall)
//	}
//	c.RenderPaddle(board.Paddle)
//	c.RenderBall(board.Ball)
//
//	for _, row := range c.Canvas {
//		for _, char := range row {
//			if char == "" {
//				print(" ")
//			} else {
//				print(char)
//			}
//		}
//		print("\n")
//	}
//}

func Aaa() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Canvas and Status")
	myWindow.Resize(fyne.NewSize(800, 600))

	// Circle
	circle := canvas.NewCircle(color.NRGBA{R: 255, A: 255})
	circle.Resize(fyne.NewSize(100, 100))
	circle.Move(fyne.NewPos(50, 50))

	// Rectangle
	rectangle := canvas.NewRectangle(color.NRGBA{G: 255, A: 255})
	rectangle.Resize(fyne.NewSize(100, 50))
	rectangle.Move(fyne.NewPos(200, 50))

	// Another Rectangle
	anotherRectangle := canvas.NewRectangle(color.NRGBA{B: 255, A: 255})
	anotherRectangle.Resize(fyne.NewSize(150, 100))
	anotherRectangle.Move(fyne.NewPos(100, 200))

	// Canvas Container
	canvasContainer := container.NewWithoutLayout(circle, rectangle, anotherRectangle)
	canvasContainer.Resize(fyne.NewSize(400, 400))

	// Status Panel
	statusLabel := widget.NewLabel("Status Info")
	statusPanel := container.NewVBox(
		statusLabel,
		// Add more status elements here
	)

	// Layout with an 80-20 split
	splitLayout := container.NewHSplit(canvasContainer, statusPanel)
	splitLayout.Offset = 0.8 // 80% for the canvas, 20% for the status
	go animateCircle(circle, 50, 50, 100, 100, 20*time.Second)
	myWindow.SetContent(splitLayout)
	myWindow.ShowAndRun()
}

func animateCircle(circle *canvas.Circle, startX, startY, endX, endY float32, duration time.Duration) {
	deltaX := (endX - startX) / float32(duration/time.Millisecond)
	deltaY := (endY - startY) / float32(duration/time.Millisecond)

	for i := 0; i < int(duration/time.Millisecond); i++ {
		time.Sleep(time.Millisecond)
		circle.Move(fyne.NewPos(startX+deltaX*float32(i), startY+deltaY*float32(i)))
		canvas.Refresh(circle)
	}
}

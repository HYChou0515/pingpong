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
	ShapeMap   map[string]fyne.CanvasObject
	app        fyne.App
	Window     fyne.Window
	OffsetX    float64
	OffsetY    float64
	WidthUnit  float64 // one unit in the game is how many pixels in the canvas
	HeightUnit float64 // one unit in the game is how many pixels in the canvas
}

func (c *FyneCanvas) PositionTransform(gameX float64, gameY float64) (float32, float32) {
	return float32(gameX*c.WidthUnit - c.OffsetX), float32(gameY*c.HeightUnit - c.OffsetY)
}

func (c *FyneCanvas) LengthTransform(gameX float64, gameY float64) (float32, float32) {
	return float32(gameX * c.WidthUnit), float32(gameY * c.HeightUnit)
}

func (c *FyneCanvas) renderRectangle(canvasObject fyne.CanvasObject, rect *Rectangle) {
	canvasObject.Move(fyne.NewPos(c.PositionTransform(rect.TopLeft.X, rect.TopLeft.Y)))
	canvasObject.Resize(fyne.NewSize(c.LengthTransform(rect.BottomRight.X-rect.TopLeft.X, rect.BottomRight.Y-rect.TopLeft.Y)))
}

func (c *FyneCanvas) renderCircle(canvasObject fyne.CanvasObject, circle *Circle) {
	canvasObject.Move(fyne.NewPos(c.PositionTransform(circle.Center.X, circle.Center.Y)))
	canvasObject.Resize(fyne.NewSize(c.LengthTransform(circle.Radius, circle.Radius)))
}

func (c *FyneCanvas) RenderBrick(brick *Brick) {
	if _, ok := c.ShapeMap[brick.Id]; !ok {
		c.ShapeMap[brick.Id] = canvas.NewRectangle(color.NRGBA{R: 255, A: 255})
	}
	rect := c.ShapeMap[brick.Id]
	c.renderRectangle(rect, brick.Shape)
}

func (c *FyneCanvas) RenderPaddle(paddle *Paddle) {
	if _, ok := c.ShapeMap[paddle.Id]; !ok {
		c.ShapeMap[paddle.Id] = canvas.NewRectangle(color.NRGBA{G: 255, A: 255})
	}
	rect := c.ShapeMap[paddle.Id]
	c.renderRectangle(rect, paddle.Shape)
}

func (c *FyneCanvas) RenderWall(wall *Wall) {
	if _, ok := c.ShapeMap[wall.Id]; !ok {
		c.ShapeMap[wall.Id] = canvas.NewRectangle(color.NRGBA{B: 255, A: 255})
	}
	rect := c.ShapeMap[wall.Id]
	c.renderRectangle(rect, wall.Shape)
}

func (c *FyneCanvas) RenderBall(ball *Ball) {
	if _, ok := c.ShapeMap[ball.Id]; !ok {
		c.ShapeMap[ball.Id] = canvas.NewCircle(color.NRGBA{R: 255, G: 255, B: 255, A: 255})
	}
	rect := c.ShapeMap[ball.Id]
	c.renderCircle(rect, ball.Shape)
}

func (c *FyneCanvas) RenderBoard(board *Board) {
	for _, brick := range board.Bricks {
		c.RenderBrick(brick)
	}
	for _, wall := range board.Wall {
		c.RenderWall(wall)
	}
	c.RenderPaddle(board.Paddle)
	c.RenderBall(board.Ball)
}

func (c *FyneCanvas) Show() {
	// put all ShapeMap objects into container.NewWithoutLayout
	objects := make([]fyne.CanvasObject, 0)
	for _, shape := range c.ShapeMap {
		objects = append(objects, shape)
	}
	canvasContainer := container.NewWithoutLayout(objects...)
	canvasContainer.Resize(fyne.NewSize(400, 400))
	c.Window.SetContent(canvasContainer)
	c.Window.ShowAndRun()
}

func MakeFyneCanvas() *FyneCanvas {
	myApp := app.New()
	myWindow := myApp.NewWindow("Window and Status")
	myWindow.Resize(fyne.NewSize(800, 600))

	return &FyneCanvas{
		ShapeMap:   make(map[string]fyne.CanvasObject),
		Window:     myWindow,
		app:        myApp,
		OffsetX:    -30,
		OffsetY:    -30,
		WidthUnit:  10,
		HeightUnit: 10,
	}
}

func Aaa() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Window and Status")
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

	// Window Container
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

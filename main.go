package main

import (
	"pingpong/game"
)

func main() {
	canvas := game.MakeFyneCanvas()
	board := game.ResetBoard()
	canvas.RenderBoard(board)
	canvas.Show()
}

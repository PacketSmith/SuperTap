package main

import rl "github.com/gen2brain/raylib-go/raylib"
import (
	"fmt"
	"math/rand/v2"
	"strconv"
)

func main() {
	var fontSize int32 = 24

	var width int = 600
	var height int = 600
	rl.InitWindow(int32(width), int32(height), "SuperTap! Native")
	rl.SetWindowMinSize(600, 600)
	icon := rl.LoadImage("icon.ico")
	rl.SetWindowIcon(*icon)
	rl.UnloadImage(icon)

	defer rl.CloseWindow()

	rl.SetTargetFPS(int32(rl.GetMonitorRefreshRate(rl.GetCurrentMonitor())))

	fmt.Printf(strconv.Itoa(rl.GetMonitorRefreshRate(rl.GetCurrentMonitor())) + "\n")
	var boxWidth int = 50
	var boxHeight int = 50
	var boxX int = rand.IntN(width-boxWidth) + 1
	var boxY int = 1
	var speed int = 500
	var score int = 0     // Reset starting score
	var gameRound int = 0 // Reset starting round
	var gameRounds int = 5

	var text = ""

	// Possible gameState(s): "start", "game", "roundOver"
	var gameState = "start"

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		switch gameState {

		case "start":
			fontSize = 40
			text = "SuperTap!"
			rl.DrawText(text, int32(width/2)-int32(rl.MeasureText(text, fontSize)/2), 275, fontSize, rl.DarkGray)
			// text = "Click to start"
			// fontSize = 24
			// rl.DrawText(text, int32(width/2)-int32(rl.MeasureText(text, fontSize)/2), 325, fontSize, rl.DarkGray)
			Button("Click to start", int32(width/2), int32(float32(height)*1.5/2), 40, func() {
				gameState = "game"
			})

		case "game":
			box := rl.Rectangle{X: float32(boxX), Y: float32(boxY), Width: float32(boxWidth), Height: float32(boxHeight)}
			rl.DrawRectangleRec(box, rl.LightGray)
			boxY = boxY + int((rl.GetFrameTime() * float32(speed)))
			if rl.IsMouseButtonReleased(rl.MouseLeftButton) && rl.CheckCollisionPointRec(rl.GetMousePosition(), box) {
				println("Box clicked!")
				score++
			}

			if boxY >= height {
				boxY = 1 - boxHeight
				boxX = rand.IntN(width-boxWidth) + 1
				gameRound++
			}

			if gameRound == gameRounds {
				gameState = "roundOver"
			}
			// Draw score last so it's always visible
			rl.DrawText(strconv.Itoa(score), int32(width/2)-int32(rl.MeasureText(strconv.Itoa(score), fontSize)/2), 100, fontSize, rl.DarkGray)

		case "roundOver":
			text = "SuperTap! Over"
			fontSize = 30
			rl.DrawText(text, int32(width/2)-int32(rl.MeasureText(text, fontSize)/2), 260, fontSize, rl.DarkGray)

			text = "Final Score: " + strconv.Itoa(score)
			fontSize = 24
			rl.DrawText(text, int32(width/2)-int32(rl.MeasureText(text, fontSize)/2), 290, fontSize, rl.DarkGray)

			text = "Click to start over"
			fontSize = 28
			rl.DrawText(text, int32(width/2)-int32(rl.MeasureText(text, fontSize)/2), 340, fontSize, rl.DarkGray)

			if rl.IsMouseButtonReleased(rl.MouseLeftButton) {
				gameState = "game"
				score = 0
				gameRound = 0
			}
		}

		rl.EndDrawing()
	}
}

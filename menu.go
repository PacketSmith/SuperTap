package main

import rl "github.com/gen2brain/raylib-go/raylib"

func Button(text string, xCenter int32, yCenter int32, height int32, onClick func()) {
	fontSize := int32(float64(height) * 0.75)
	width := height
	if rl.MeasureText(text, fontSize)/2 > height {
		width = int32(float32(rl.MeasureText(text, fontSize)) * 1.25)
	}
	x := xCenter - width/2
	y := yCenter - height/2
	box := rl.Rectangle{X: float32(x), Y: float32(y), Width: float32(width), Height: float32(height)}
	rl.DrawRectangleRec(box, rl.LightGray)
	rl.DrawText(text, x+(width-rl.MeasureText(text, fontSize))/2, y+height/5, fontSize, rl.DarkGray)

	if rl.IsMouseButtonReleased(rl.MouseLeftButton) && rl.CheckCollisionPointRec(rl.GetMousePosition(), box) {
		println("Button clicked!")
		onClick()
	}
}

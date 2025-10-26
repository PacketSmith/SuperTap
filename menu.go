// I'm avoiding converting int to int32 till the last moment cause it makes development easier

package main

import rl "github.com/gen2brain/raylib-go/raylib"

func Button(text string, xCenter int, yCenter int, height int, onClick func()) {
	fontSize := int(float64(height) * 0.75)
	width := height
	if rl.MeasureText(text, int32(fontSize))/2 > int32(height) {
		width = int((rl.MeasureText(text, int32(float64(fontSize)*float64(1.25))))) // ignore type shenanigans, it made the error demons go away
	}
	x := xCenter - width/2
	y := yCenter - height/2
	box := rl.Rectangle{X: float32(x), Y: float32(y), Width: float32(width), Height: float32(height)}
	rl.DrawRectangleRec(box, rl.LightGray)
	rl.DrawText(text, int32(x)+(int32(width)-rl.MeasureText(text, int32(fontSize)))/2, int32(y+height/5), int32(fontSize), rl.DarkGray)

	if rl.IsMouseButtonReleased(rl.MouseLeftButton) && rl.CheckCollisionPointRec(rl.GetMousePosition(), box) {
		println("Button clicked!")
		onClick()
	}
}

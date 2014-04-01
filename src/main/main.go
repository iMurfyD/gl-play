package main

import glfw "github.com/go-gl/glfw3"
import gl "github.com/go-gl/gl"
import "fmt"

func main() {
	if !glfw.Init() {
		fmt.Println("No glfw.Init()\nMuch Error\nSo fail")
	}

	glfw.SetErrorCallback(error_callback)

	monitor, _ := glfw.GetMonitors()
	m := monitor[0]
	vid, _ := m.GetVideoMode()
	x_res := vid.Width
	y_res := vid.Height
	window, _ := glfw.CreateWindow(x_res, y_res, "Hello World", m, nil)
	window.MakeContextCurrent()
	window.SetKeyCallback(key_callback)

	for !window.ShouldClose() {
		var ratio float64
		width, height := window.GetFramebufferSize()
		ratio = float64(width) / float64(height)

		gl.Viewport(0, 0, width, height)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		gl.MatrixMode(gl.PROJECTION)
		gl.LoadIdentity()
		gl.Ortho(-ratio, ratio, -1, 1, 1, -1)
		gl.MatrixMode(gl.MODELVIEW)

		gl.LoadIdentity()
		//original rotate
		//gl.Rotatef(float32(glfw.GetTime()*50), 0, 0, 1)

		gl.Rotatef(float32(glfw.GetTime()*250), -1, 0, 1)

		//original triangles
		gl.Begin(gl.TRIANGLES)
		gl.Color3f(1, 0, 0)
		gl.Vertex3f(-.6, -.4, 0)
		gl.Color3f(0, 1, 0)
		gl.Vertex3f(.6, -.4, 0)
		gl.Color3f(0, 0, 1)
		gl.Vertex3f(0, .6, 0)
		gl.End()

		window.SwapBuffers()
		glfw.PollEvents()
	}
	glfw.Terminate()
}

func key_callback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if key == glfw.KeyQ {
		w.SetShouldClose(true)
	}
}

func error_callback(code glfw.ErrorCode, desc string) {
	fmt.Println("Something went wrong...")
	fmt.Println(desc)
}

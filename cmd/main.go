package main

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"log"
	"runtime"
)

const (
	width  = 800
	height = 600
)

func init() {
	runtime.LockOSThread()
}

func main() {
	err := glfw.Init()
	if err != nil {
		log.Fatal("Failed to initialize GLFW:", err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.Resizable, glfw.False)

	window, err := glfw.CreateWindow(width, height, "OpenGL Point", nil, nil)
	if err != nil {
		log.Fatal("Failed to create GLFW window:", err)
	}
	window.MakeContextCurrent()

	err = gl.Init()
	if err != nil {
		log.Fatal("Failed to initialize OpenGL:", err)
	}

	gl.Viewport(0, 0, width, height)
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(0, float64(width), 0, float64(height), -1, 1)

	for !window.ShouldClose() {
		gl.ClearColor(0.0, 0.0, 0.0, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		gl.Begin(gl.POINTS)
		gl.Color3f(1.0, 0.0, 0.0) // Устанавливаем цвет точки (красный)
		gl.Vertex2f(400, 300)     // Устанавливаем координаты точки
		gl.End()

		window.SwapBuffers()
		glfw.PollEvents()
	}
}

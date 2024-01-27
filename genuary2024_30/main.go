// Package main is the starting point of the app.
package main

//revive:disable:unused-parameter

import (
	"fmt"
	"log"
	"math"
	"os"
	"runtime"
	"strings"
	"unsafe"

	"github.com/bit101/bitlib/bitmap"
	"github.com/bit101/blcairo/render"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	width  = 500
	height = 500
)

var (
	triangle = []float32{
		// -0.5, -0.5, 0.0,
		// 0.5, -0.5, 0.0,
		// 0.0, 0.5, 0.0,
		-1, 1, 0,
		1, 1, 0,
		-1, -1, 0,
		1, -1, 0,
	}
	frame = 0
)

func main() {
	// doit()
	renderit()
}

func doit() {
	runtime.LockOSThread()

	window := initGlfw()
	defer glfw.Terminate()
	program := initOpenGL()

	vao := makeVao(triangle)
	for !window.ShouldClose() {
		draw(vao, window, program)
	}
}

func renderit() {
	render.UseBMP(true)
	render.ConvertToVideo("out", "out.mp4", 500, 500, 60, 10)
	render.PlayVideo("out.mp4")
}

func draw(vao uint32, window *glfw.Window, program uint32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	timeValue := glfw.GetTime()
	greenValue := float32((math.Sin(timeValue) / 2.0) + 0.5)
	blueValue := float32((math.Sin(timeValue) * 2.0) + 0.5)
	vertexColorLocation := gl.GetUniformLocation(program, gl.Str("ourColor\x00"))
	gl.UseProgram(program)
	gl.Uniform4f(vertexColorLocation, 1.0-greenValue, greenValue, blueValue, 1.0)

	gl.BindVertexArray(vao)
	gl.DrawArrays(gl.TRIANGLE_STRIP, 0, int32(len(triangle)/3))

	processInput(window)
	glfw.PollEvents()
	window.SwapBuffers()

	if frame < 600 {
		var bytes [width * height * 3]byte
		pointer := unsafe.Pointer(&(bytes[0]))

		bmp := bitmap.NewBitmap(width, height)

		gl.ReadPixels(0, 0, width, height, gl.RGB, gl.UNSIGNED_BYTE, pointer)
		x := 0
		y := 0
		for i := 0; i < width*height*3; i += 3 {
			r := float64(bytes[i]) / 255.0
			g := float64(bytes[i+1]) / 255.0
			b := float64(bytes[i+2]) / 255.0
			bmp.SetPixel(x, y, r, g, b)
			x++
			if x >= width {
				x = 0
				y++
			}
		}
		bmp.SaveImage(fmt.Sprintf("out/frame_%04d.bmp", frame))
		fmt.Println(frame)
		frame++
	}
}

// initGlfw initializes glfw and returns a Window to use.
func initGlfw() *glfw.Window {
	if err := glfw.Init(); err != nil {
		panic(err)
	}
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(width, height, "Conway's Game of Life", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	return window
}

// initOpenGL initializes OpenGL and returns an intiialized program.
func initOpenGL() uint32 {
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	vertexShader, err := compileShader("shaders/vertex_shader.gl", gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}

	fragmentShader, err := compileShader("shaders/fragment_shader.gl", gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}

	prog := gl.CreateProgram()
	gl.AttachShader(prog, vertexShader)
	gl.AttachShader(prog, fragmentShader)
	gl.LinkProgram(prog)
	return prog
}

// makeVao initializes and returns a vertex array from the points provided.
func makeVao(points []float32) uint32 {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(points), gl.Ptr(points), gl.STATIC_DRAW)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	return vao
}

func compileShader(path string, shaderType uint32) (uint32, error) {
	sourceBytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	source := string(sourceBytes)
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}

func processInput(window *glfw.Window) {
	if glfw.GetCurrentContext().GetKey(glfw.KeyEscape) == glfw.Press {
		window.SetShouldClose(true)
	}
}

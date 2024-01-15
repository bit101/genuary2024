// Package main is main
package main

import (
	"fmt"
	"log"
	"runtime"
	"strings"
	"unsafe"

	"github.com/bit101/bitlib/bitmap"
	"github.com/bit101/bitlib/random"
	"github.com/bit101/blcairo/render"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	width  = 500
	height = 500
)

var (
	frame              = 0
	points             = []float32{}
	vel                = []float32{}
	vertexShaderSource = `
    #version 410
    in vec3 vp;
    void main() {
        gl_Position = vec4(vp, 1.0);
    }
` + "\x00"

	fragmentShaderSource = `
    #version 410
    out vec4 frag_colour;
    void main() {
        frag_colour = vec4(1, 1, 1, 1);
    }
` + "\x00"
)

func main() {
	doit()
	// renderit()
}

func renderit() {
	render.UseBMP(true)
	render.ConvertToVideo("out", "out.mp4", 500, 500, 60, 10)
	render.PlayVideo("out.mp4")
}

func doit() {
	runtime.LockOSThread()
	for i := 0; i < 1000; i++ {
		points = append(points, float32(random.FloatRange(-1, 1)))
		points = append(points, float32(random.FloatRange(-1, 1)))
		points = append(points, float32(random.FloatRange(-1, 1)))
		vel = append(vel, float32(random.FloatRange(-0.01, 0.01)))
		vel = append(vel, float32(random.FloatRange(-0.01, 0.01)))
		vel = append(vel, 0)
	}

	window := initGlfw()
	defer glfw.Terminate()

	program := initOpenGL()

	for !window.ShouldClose() {
		update()
		vao := makeVao(points)
		draw(vao, window, program)
	}
}

func update() {
	for i := 0; i < 3000; i += 3 {
		points[i] += vel[i]
		if points[i] > 1 {
			points[i] -= 2
		}
		if points[i] < -1 {
			points[i] += 2
		}
		if points[i+1] > 1 {
			points[i+1] -= 2
		}
		if points[i+1] < -1 {
			points[i+1] += 2
		}
		points[i+1] += vel[i+1]
		vel[i] += float32(random.FloatRange(-0.001, 0.001))
		vel[i] *= 0.99
		vel[i+1] += float32(random.FloatRange(-0.001, 0.001))
		vel[i+1] *= 0.99
	}
}

func draw(vao uint32, window *glfw.Window, program uint32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)

	gl.BindVertexArray(vao)
	gl.DrawArrays(gl.POINTS, 0, int32(len(points)/3))

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

	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.ContextVersionMajor, 4) // OR 2
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

	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}
	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
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

func compileShader(source string, shaderType uint32) (uint32, error) {
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

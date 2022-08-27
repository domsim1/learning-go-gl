package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/domsim1/go-gecko/gecko"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func init() {
	runtime.LockOSThread()
}

var vertexShader = `#version 330 core

layout(location = 0) in vec2 position;
layout(location = 1) in vec4 color;

uniform vec2 u_Resolution;

out vec4 v_Color;

void main()
{
  vec2 zeroToOne = position / u_Resolution;
  vec2 zeroToTwo = zeroToOne * 2.0;
  vec2 clipSpace = zeroToTwo - 1.0;
  gl_Position = vec4(clipSpace * vec2(1, -1), 0, 1);
  v_Color = color;
}
` + "\x00"

var fragmentShader = `#version 330 core

layout(location = 0) out vec4 color;

in vec4 v_Color;

void main()
{
  color = v_Color;
}
` + "\x00"

type Vec2 struct {
	x float32
	y float32
}

type Color struct {
	R float32
	G float32
	B float32
	A float32
}

type Vertex struct {
	Position Vec2
	Color    Color
}

func main() {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(640, 480, "Testing", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()

	glfw.SwapInterval(1)

	if err := gl.Init(); err != nil {
		panic(err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Printf("Init OpenGL: v%s\n", version)

	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	bufferSize := 2024

	indices := make([]uint32, 2024*4)

	for i := 0; i < bufferSize; i += 6 {
		indices[i] = uint32(i)
		indices[i+1] = uint32(i + 1)
		indices[i+2] = uint32(i + 2)
		indices[i+3] = uint32(i + 2)
		indices[i+4] = uint32(i + 1)
		indices[i+5] = uint32(i + 3)
	}

	va := gecko.NewVertexArray()
	vb := gecko.NewDynamicVertexBuffer((6 * 4) * bufferSize)

	layout := gecko.NewVertexBufferLayout()
	layout.PushFloat32(2)
	layout.PushFloat32(4)
	va.AddBuffer(vb, layout)

	ib := gecko.NewIndexBuffer(&indices, len(indices))
	ib.Bind()

	red := Color{
		R: 255,
		G: 0,
		B: 0,
		A: 255,
	}

	vertices := []Vertex{
		{
			Position: Vec2{
				x: 0, y: 0,
			},
			Color: red,
		},
		{
			Position: Vec2{
				x: 30, y: 0,
			},
			Color: red,
		},
		{
			Position: Vec2{
				x: 0, y: 30,
			},
			Color: red,
		},
		{
			Position: Vec2{
				x: 30, y: 30,
			},
			Color: red,
		},
	}

	shader := gecko.NewShader(vertexShader, fragmentShader)
	shader.Bind()

	shader.SetUniform2f("u_Resolution", 640, 480)

	lastTime := time.Now()
	framecount := 0
	for !window.ShouldClose() {
		framecount++
		if time.Since(lastTime) > time.Second {
			lastTime = time.Now()
			println(framecount)
			framecount = 0
		}
		gecko.Clear()

		vb.Push(vertices, ((6 * 4) * len(vertices)))

		gecko.Draw(vb, ib, shader)
		window.SwapBuffers()
		glfw.PollEvents()
	}
}

func NewRect() {

}

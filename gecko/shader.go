package gecko

import (
	"fmt"

	"github.com/go-gl/gl/v3.3-core/gl"
)

type Shader struct {
	rendererID           uint32
	uniformLocationCache map[string]int32
}

func NewShader(vertexShaderSource string, fragmentShaderSource string) *Shader {
	shader := &Shader{
		rendererID:           createShader(vertexShaderSource, fragmentShaderSource),
		uniformLocationCache: make(map[string]int32),
	}
	return shader
}

func (s *Shader) Delete() {
	GLClearError()
	gl.DeleteProgram(s.rendererID)
	GLCheckError()
}

func (s *Shader) Bind() {
	GLClearError()
	gl.UseProgram(s.rendererID)
	GLCheckError()
}

func (s *Shader) Unbind() {
	GLClearError()
	gl.UseProgram(0)
	GLCheckError()
}

func (s *Shader) SetUniform2f(name string, v1 float32, v2 float32) {
	location := s.getUniformLocation(name)
	GLClearError()
	gl.Uniform2f(location, v1, v2)
	GLCheckError()
}

func (s *Shader) SetUniform4f(name string, v1 float32, v2 float32, v3 float32, v4 float32) {
	location := s.getUniformLocation(name)
	GLClearError()
	gl.Uniform4f(location, v1, v2, v3, v4)
	GLCheckError()
}

func (s *Shader) SetUniform1i(name string, v1 int32) {
	location := s.getUniformLocation(name)
	GLClearError()
	gl.Uniform1i(location, v1)
	GLCheckError()
}

func (s *Shader) getUniformLocation(name string) int32 {
	if location, ok := s.uniformLocationCache[name]; ok {
		return location
	}
	cname := name + "\x00"
	GLCheckError()
	location := gl.GetUniformLocation(s.rendererID, gl.Str(cname))
	GLCheckError()
	if location == -1 {
		fmt.Printf("Warning: could not get location of uniform %s\n", name)
	}
	s.uniformLocationCache[name] = location
	return location
}

func compileShader(shaderType uint32, source string) uint32 {
	id := gl.CreateShader(shaderType)
	src := gl.Str(source)
	gl.ShaderSource(id, 1, &src, nil)
	gl.CompileShader(id)

	var result int32
	gl.GetShaderiv(id, gl.COMPILE_STATUS, &result)
	if result == gl.FALSE {
		var length int32
		gl.GetShaderiv(id, gl.INFO_LOG_LENGTH, &length)
		message := make([]byte, length)
		message[length-1] = '\x00'
		msg := string(message)

		gl.GetShaderInfoLog(id, length, &length, gl.Str(msg))
		shaderTypeName := "fragment"
		if shaderType == gl.VERTEX_SHADER {
			shaderTypeName = "vertex"
		}
		gl.DeleteShader(id)
		panic(fmt.Sprintf("Failed to compile %s shader!\n%s", shaderTypeName, msg))
	}
	return id
}

func createShader(vertexShader string, fragmentShader string) uint32 {
	program := gl.CreateProgram()
	vs := compileShader(gl.VERTEX_SHADER, vertexShader)
	fs := compileShader(gl.FRAGMENT_SHADER, fragmentShader)

	gl.AttachShader(program, vs)
	gl.AttachShader(program, fs)
	gl.LinkProgram(program)
	gl.ValidateProgram(program)

	gl.DeleteShader(vs)
	gl.DeleteShader(fs)

	return program
}

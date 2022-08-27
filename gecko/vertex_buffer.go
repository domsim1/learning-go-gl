package gecko

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

type VertexBuffer struct {
	rendererID uint32
}

func NewVertexBuffer(data any, size int) *VertexBuffer {
	buffer := &VertexBuffer{}
	GLClearError()
	gl.GenBuffers(1, &buffer.rendererID)
	gl.BindBuffer(gl.ARRAY_BUFFER, buffer.rendererID)
	gl.BufferData(gl.ARRAY_BUFFER, size, gl.Ptr(data), gl.STATIC_DRAW)
	GLCheckError()
	return buffer
}

func NewDynamicVertexBuffer(max_size int) *VertexBuffer {
	buffer := &VertexBuffer{}
	GLClearError()
	gl.GenBuffers(1, &buffer.rendererID)
	gl.BindBuffer(gl.ARRAY_BUFFER, buffer.rendererID)
	gl.BufferData(gl.ARRAY_BUFFER, max_size, nil, gl.DYNAMIC_DRAW)
	GLCheckError()
	return buffer
}

func (vb *VertexBuffer) Delete() {
	GLCheckError()
	gl.DeleteBuffers(1, &vb.rendererID)
	GLCheckError()
}

func (vb *VertexBuffer) Bind() {
	GLClearError()
	gl.BindBuffer(gl.ARRAY_BUFFER, vb.rendererID)
	GLCheckError()
}

func (vb *VertexBuffer) Unbind() {
	GLClearError()
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	GLCheckError()
}

func (vb *VertexBuffer) Push(data any, size int) {
	GLClearError()
	vb.Bind()
	gl.BufferSubData(gl.ARRAY_BUFFER, 0, size, gl.Ptr(data))
	GLCheckError()
}

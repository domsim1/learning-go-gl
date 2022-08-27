package gecko

import "github.com/go-gl/gl/v3.3-core/gl"

type IndexBuffer struct {
	rendererID uint32
	count      int32
}

func NewIndexBuffer(data *[]uint32, count int) *IndexBuffer {
	buffer := &IndexBuffer{
		count: int32(count),
	}
	GLClearError()
	gl.GenBuffers(1, &buffer.rendererID)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, buffer.rendererID)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, count*4, gl.Ptr(*data), gl.STATIC_DRAW)
	GLCheckError()
	return buffer
}

func (ib *IndexBuffer) Delete() {
	GLClearError()
	gl.DeleteBuffers(1, &ib.rendererID)
	GLCheckError()
}

func (ib *IndexBuffer) Bind() {
	GLClearError()
	gl.BindBuffer(gl.ARRAY_BUFFER, ib.rendererID)
	GLCheckError()
}

func (ib *IndexBuffer) Unbind() {
	GLClearError()
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	GLCheckError()
}

func (ib *IndexBuffer) Count() int32 {
	return ib.count
}

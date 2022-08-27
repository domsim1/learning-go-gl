package gecko

import "github.com/go-gl/gl/v3.3-core/gl"

type VertexArray struct {
	rendererID uint32
}

func NewVertexArray() *VertexArray {
	buffer := &VertexArray{}
	GLClearError()
	gl.GenVertexArrays(1, &buffer.rendererID)
	GLCheckError()
	return buffer
}

func (va *VertexArray) Delete() {
	GLCheckError()
	gl.DeleteVertexArrays(1, &va.rendererID)
	GLCheckError()
}

func (va *VertexArray) AddBuffer(vb *VertexBuffer, layout *VertexBufferLayout) {
	GLClearError()
	va.Bind()
	vb.Bind()
	elements := layout.Elements()
	offset := 0
	for i, element := range elements {
		gl.EnableVertexAttribArray(uint32(i))
		gl.VertexAttribPointer(uint32(i), int32(element.count), element.xtype, element.normalized, int32(layout.Stride()), gl.PtrOffset(offset))
		offset += int(element.count * uint32(GetSizeOfType(element.xtype)))
	}
	GLCheckError()
}

func (va *VertexArray) Bind() {
	GLClearError()
	gl.BindVertexArray(va.rendererID)
	GLCheckError()
}

func (va *VertexArray) Unbind() {
	GLClearError()
	gl.BindVertexArray(0)
	GLCheckError()
}

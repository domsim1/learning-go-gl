package gecko

import "github.com/go-gl/gl/v3.3-core/gl"

type VertexBufferLayoutElement struct {
	count      uint32
	xtype      uint32
	normalized bool
}

func GetSizeOfType(xtype uint32) uint32 {
	switch xtype {
	case gl.FLOAT:
		return 4
	case gl.UNSIGNED_INT:
		return 4
	case gl.UNSIGNED_BYTE:
		return 1
	}
	panic("unsupported type")
}

type VertexBufferLayout struct {
	elements []*VertexBufferLayoutElement
	stride   uint32
}

func NewVertexBufferLayout() *VertexBufferLayout {
	return &VertexBufferLayout{
		elements: make([]*VertexBufferLayoutElement, 0, 1),
		stride:   0,
	}
}

func (vbl *VertexBufferLayout) PushFloat32(count uint32) {
	vbl.elements = append(vbl.elements, &VertexBufferLayoutElement{
		count:      count,
		xtype:      gl.FLOAT,
		normalized: false,
	})
	vbl.stride += count * GetSizeOfType(gl.FLOAT)
}

func (vbl *VertexBufferLayout) PushUint32(count uint32) {
	vbl.elements = append(vbl.elements, &VertexBufferLayoutElement{
		count:      count,
		xtype:      gl.UNSIGNED_INT,
		normalized: false,
	})
	vbl.stride += count * GetSizeOfType(gl.UNSIGNED_INT)
}

func (vbl *VertexBufferLayout) PushUint8(count uint32) {
	vbl.elements = append(vbl.elements, &VertexBufferLayoutElement{
		count:      count,
		xtype:      gl.UNSIGNED_BYTE,
		normalized: true,
	})
	vbl.stride += count * GetSizeOfType(gl.UNSIGNED_BYTE)
}

func (vbl *VertexBufferLayout) Elements() []*VertexBufferLayoutElement {
	return vbl.elements
}

func (vbl *VertexBufferLayout) Stride() uint32 {
	return vbl.stride
}

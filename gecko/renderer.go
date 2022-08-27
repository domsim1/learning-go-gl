package gecko

import (
	"fmt"

	"github.com/go-gl/gl/v3.3-core/gl"
)

func GLClearError() {
	for gl.GetError() != gl.NO_ERROR {
	}
}

func GLCheckError() {
	errs := make([]uint32, 0)
	for {
		err := gl.GetError()
		if err == gl.NO_ERROR {
			break
		}
		errs = append(errs, err)
	}
	for _, err := range errs {
		fmt.Printf("[OpenGL Error!] %d\n", err)
	}
	if len(errs) > 0 {
		panic("OpenGL Errors!")
	}
}

func Draw(va *VertexBuffer, ib *IndexBuffer, shader *Shader) {
	shader.Bind()
	va.Bind()
	ib.Bind()
	GLClearError()
	gl.DrawElements(gl.TRIANGLES, ib.Count(), gl.UNSIGNED_INT, nil)
	GLCheckError()
}

func Clear() {
	GLClearError()
	gl.Clear(gl.COLOR_BUFFER_BIT)
	GLCheckError()
}

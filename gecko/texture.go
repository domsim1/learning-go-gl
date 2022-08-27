package gecko

import (
	"image/png"
	"os"
	"sync"

	"github.com/go-gl/gl/v3.3-core/gl"
)

type Texture struct {
	rendererID  uint32
	filePath    string
	localBuffer []uint8
	width       int32
	height      int32
}

func NewTextureFromPNG(path string) *Texture {
	t := &Texture{
		filePath: path,
	}
	imgFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	img, err := png.Decode(imgFile)
	if err != nil {
		panic(err)
	}
	var wg sync.WaitGroup
	yMax := img.Bounds().Max.Y
	xMax := img.Bounds().Max.X
	t.width = int32(xMax)
	t.height = int32(yMax)
	t.localBuffer = make([]uint8, (yMax * (xMax * 4)))
	for y := 0; y < yMax; y++ {
		wg.Add(1)
		go func(y int) {
			defer wg.Done()
			for x := 0; x < xMax; x++ {
				r, g, b, a := img.At(x, -(y - yMax)).RGBA()
				p := (y * (xMax * 4)) + x*4
				t.localBuffer[p] = uint8(r)
				t.localBuffer[p+1] = uint8(g)
				t.localBuffer[p+2] = uint8(b)
				t.localBuffer[p+3] = uint8(a)
			}
		}(y)
	}
	wg.Wait()
	GLClearError()
	gl.GenTextures(1, &t.rendererID)
	gl.BindTexture(gl.TEXTURE_2D, t.rendererID)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)

	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA8, t.width, t.height, 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(t.localBuffer))
	gl.BindTexture(gl.TEXTURE_2D, 0)
	GLCheckError()

	if len(t.localBuffer) > 0 {
		t.localBuffer = nil
	}
	return t
}

func (t *Texture) Delete() {
	gl.DeleteTextures(1, &t.rendererID)
}

func (t *Texture) Bind(slot uint32) {
	gl.ActiveTexture(gl.TEXTURE0 + slot)
	gl.BindTexture(gl.TEXTURE_2D, t.rendererID)
}

func (t *Texture) Unbind() {
	gl.BindTexture(gl.TEXTURE_2D, 0)
}

func (t *Texture) Width() int32 {
	return t.width
}

func (t *Texture) Height() int32 {
	return t.height
}

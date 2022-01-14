package interfaces

type Fuzable interface {
	Fuzz()
}

type Image struct {
	Width    int
	Height   int
	Contents [][]byte
	fuzzed   bool
}

func (c *Image) Fuzz() {
	c.fuzzed = true
}

func New(width int, height int, contents [][]byte) Image {
	return Image{width, height, contents, false}
}

func (c Image) IsFuzzy() bool {
	return c.fuzzed
}

func FuzzTool(fz Fuzable) {
	fz.Fuzz()
}

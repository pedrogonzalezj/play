package interfaces_test

import (
	"testing"

	"github.com/pedrogonzalezj/play/interfaces"
)

func TestNewImage(t *testing.T) {
	data := make([][]byte, 4)
	img := interfaces.New(300, 600, data)

	if img.Width != 300 {
		t.Errorf("got %d want %d", img.Width, 300)
	}
	if img.Height != 600 {
		t.Errorf("got %d want %d", img.Height, 600)
	}
	if img.IsFuzzy() {
		t.Errorf("got %v want %v", img.IsFuzzy(), false)
	}
}

func TestFuzzTool(t *testing.T) {
	data := make([][]byte, 4)
	img := interfaces.New(300, 600, data)
	interfaces.FuzzTool(&img)
	got := img.IsFuzzy()
	if !got {
		t.Errorf("got %v want %v", got, true)
	}
}

package structs

import (
	"testing"
)

func TestButton(t *testing.T) {

	t.Run("Button Click", func(t *testing.T) {
		want := true
		btn := Button{Clicked: want}
		if !btn.Clicked {
			t.Errorf("got %v want %v", btn.Clicked, want)
		}
	})

	t.Run("Toogle", func(t *testing.T) {
		click := true
		btn := Button{Clicked: click}
		got := btn.Toggle()
		want := "unclicked"
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}

		btn.Clicked = false
		got = btn.Toggle()
		want = "clicked"
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}

func TestOverWrittenWindow(t *testing.T) {

	t.Run("Window Hide", func(t *testing.T) {
		want := true
		wdn := OverWrittenWindow{Hidden: want}
		if !wdn.Hidden {
			t.Errorf("got %v want %v", wdn.Hidden, want)
		}
	})

	t.Run("Toogle", func(t *testing.T) {
		wdn := OverWrittenWindow{
			Button: Button{
				Clicked: true,
			},
			Hidden: false,
		}
		got := wdn.Toggle()
		want := "hidden"
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}

		wdn.Clicked = false
		got = wdn.Toggle()
		want = "visible"
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}

func TestCompositeWindowWindow(t *testing.T) {

	t.Run("Toogle", func(t *testing.T) {
		wdn := CompositeWindow{
			Button: Button{
				Clicked: true,
			},
		}
		got := wdn.Toggle()
		want := "unclicked"
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}

		wdn.Clicked = false
		got = wdn.Toggle()
		want = "clicked"
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}

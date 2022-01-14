package structs

type CompositeWindow struct {
	Button
}

type OverWrittenWindow struct {
	Button
	Hidden bool
}

func (w *OverWrittenWindow) Toggle() string {
	w.Hidden = !w.Hidden
	if w.Hidden {
		return "hidden"
	}
	return "visible"
}

type Button struct {
	Clicked bool
}

func (b *Button) Toggle() string {
	b.Clicked = !b.Clicked
	if b.Clicked {
		return "clicked"
	}
	return "unclicked"
}

package gtk

import (
	"github.com/ziutek/gdk"
	"testing"
	"fmt"
)

func cbDelete(w *Widget, ev *gdk.Event) bool {
	fmt.Println("Delete")
	return false
}

func cbDestroy(w *Widget) {
	fmt.Println("Destroy")
	MainQuit()
}

func TestHello(t *testing.T) {
	w := NewWindow(WINDOW_TOPLEVEL)
	w.Connect("delete-event", cbDelete, nil)
	w.Connect("destroy", cbDestroy, nil)
	w.SetBorderWidth(10)
	w.Show()

	a := A{"Hello World"}

	b := NewButtonWithLabel("Hello World")
	b.Connect("clicked", (*A).Hello, &a)
	w.Add(b.AsWidget())
	b.Show()

	Main()
}

type A struct {
	s string
}

func (a *A) Hello(w *Widget) {
	fmt.Printf(a.s)
}

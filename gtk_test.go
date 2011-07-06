package gtk

import (
	"github.com/ziutek/gdk"
	"testing"
	"fmt"
)

func cbDelete(e *gdk.Event) bool {
	fmt.Println("Delete")
	return false
}

func cbDestroy(w *Widget) {
	fmt.Println("Destroy")
	MainQuit()
}

func TestHello(t *testing.T) {
	w := NewWindow(WINDOW_TOPLEVEL)
	w.ConnectByName("delete-event", cbDelete)
	//w.ConnectByName("destroy", cbDestroy)
	w.Show()

	Main()
}

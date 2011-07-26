package gtk

/*
#include <gtk/gtk.h>
*/
import "C"

import (
	"github.com/ziutek/glib"
)

type WindowType C.GtkWindowType

const (
	WINDOW_TOPLEVEL = WindowType(C.GTK_WINDOW_TOPLEVEL)
	WINDOW_POPUP = WindowType(C.GTK_WINDOW_POPUP)
)

type Window struct {
	Bin
}

func (w *Window) g() *C.GtkWindow {
	return (*C.GtkWindow)(w.GetPtr())
}

func (w *Window) AsWindow() *Window {
	return w
}

// Returns C pointer
func NewWindow(t WindowType) *Window {
	w := new(Window)
	w.SetPtr(glib.Pointer(C.gtk_window_new(C.GtkWindowType(t))))
	return w
}

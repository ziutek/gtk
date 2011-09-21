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

func (w *Window) SetTitle(title string) {
	C.gtk_window_set_title(w.g(), (*C.gchar)(C.CString(title)))
}

func (w *Window) SetDefaultSize(width, height int) {
	C.gtk_window_set_default_size(w.g(), C.gint(width), C.gint(height))
}

// Returns C pointer
func NewWindow(t WindowType) *Window {
	w := new(Window)
	w.SetPtr(glib.Pointer(C.gtk_window_new(C.GtkWindowType(t))))
	return w
}

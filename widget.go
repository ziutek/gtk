package gtk

/*
#include <gtk/gtk.h>
*/
import "C"

import (
	"github.com/ziutek/glib"
)

type Widget struct {
	glib.Object
}

func (w *Widget) g() *C.GtkWidget {
	return (*C.GtkWidget)(w.Pointer())
}

func (w *Widget) AsWidget() *Widget {
		return w
}

func (w *Widget) Destroy() {
	C.gtk_widget_destroy(w.g())
}

func (w *Widget) Show() {
	C.gtk_widget_show(w.g())
}

func (w *Widget) ShowNow() {
	C.gtk_widget_show_now(w.g())
}

func (w *Widget) ShowAll() {
	C.gtk_widget_show_all(w.g())
}

func (w *Widget) Hide() {
	C.gtk_widget_hide(w.g())
}

func (w *Widget) HideAll() {
	C.gtk_widget_hide_all(w.g())
}

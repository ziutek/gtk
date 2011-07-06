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

func (w *Widget) GtkWidget() *C.GtkWidget {
	return (*C.GtkWidget)(w.GPointer())
}

func (w *Widget) Destroy() {
	C.gtk_widget_destroy(w.GtkWidget())
}

func (w *Widget) Show() {
	C.gtk_widget_show(w.GtkWidget())
}

func (w *Widget) ShowNow() {
	C.gtk_widget_show_now(w.GtkWidget())
}

func (w *Widget) ShowAll() {
	C.gtk_widget_show_all(w.GtkWidget())
}

func (w *Widget) Hide() {
	C.gtk_widget_hide(w.GtkWidget())
}

func (w *Widget) HideAll() {
	C.gtk_widget_hide_all(w.GtkWidget())
}

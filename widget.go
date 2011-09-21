package gtk

/*
#include <gtk/gtk.h>
*/
import "C"

import (
	"github.com/ziutek/glib"
	"github.com/ziutek/gdk"
)

type Widget struct {
	glib.Object
}

func (w *Widget) g() *C.GtkWidget {
	return (*C.GtkWidget)(w.GetPtr())
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

func (w *Widget) Realize() {
	C.gtk_widget_realize(w.g())
}

func (w *Widget) GetWindow() *gdk.Window {
	gw := new(gdk.Window)
	gw.SetPtr(glib.Pointer(C.gtk_widget_get_window(w.g())))
	//gw.SetPtr(glib.Pointer(w.g().window))
	return gw
}

func (w *Widget) SetDoubleBuffered(double_buffered bool) {
	var db C.gboolean
	if double_buffered {
		db = 1
	}
	C.gtk_widget_set_double_buffered(w.g(), db)
}

func (w *Widget) SetSizeRequest(width, height int) {
	C.gtk_widget_set_size_request(w.g(), C.gint(width), C.gint(height))
}

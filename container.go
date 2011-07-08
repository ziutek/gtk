package gtk

/*
#include <gtk/gtk.h>
*/
import "C"

type Container struct {
	Widget
}

func (k *Container) g() *C.GtkContainer {
	return (*C.GtkContainer)(k.Pointer())
}

func (k *Container) AsContainer() *Container {
	return k
}

func (k *Container) Add(w *Widget) {
	C.gtk_container_add(k.g(), w.g())
}

func (k *Container) SetBorderWidth(width uint) {
	C.gtk_container_set_border_width(k.g(), C.guint(width))
}

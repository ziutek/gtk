package gtk

/*
#include <gtk/gtk.h>
*/
import "C"

type Container struct {
	Widget
}

func (k *Container) GtkContainer() *C.GtkContainer {
	return (*C.GtkContainer)(k.GPointer())
}

func (k *Container) Add(w *Widget) {
	C.gtk_container_add(k.GtkContainer(), w.GtkWidget())
}


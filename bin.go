package gtk

/*
#include <gtk/gtk.h>
*/
import "C"

import (
	"github.com/ziutek/glib"
)

type Bin struct {
	Container
}

func (b *Bin) GtkBin() *C.GtkBin {
	return (*C.GtkBin)(b.GPointer())
}

func (b *Bin) GetChild() *Widget {
	w := new(Widget)
	w.Set(glib.Pointer(C.gtk_bin_get_child(b.GtkBin())))
	return w
}

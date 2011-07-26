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

func (b *Bin) g() *C.GtkBin {
	return (*C.GtkBin)(b.GetPtr())
}

func (b *Bin) AsBin() *Bin {
	return b
}

func (b *Bin) GetChild() *Widget {
	w := new(Widget)
	w.SetPtr(glib.Pointer(C.gtk_bin_get_child(b.g())))
	return w
}

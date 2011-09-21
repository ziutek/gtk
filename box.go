package gtk

/*
#include <gtk/gtk.h>
*/
import "C"

import (
	"github.com/ziutek/glib"
)

type Box struct {
	Container
}

func (b *Box) g() *C.GtkBox {
	return (*C.GtkBox)(b.GetPtr())
}

func (b *Box) AsBox() *Box {
	return b
}

func (b *Box) PackStart(child *Widget, expand, fill bool, padding uint) {
	var e, f C.gboolean
	if expand {
		e = 1
	}
	if fill {
		f = 1
	}
	C.gtk_box_pack_start(b.g(), child.g(), e, f, C.guint(padding))
}

func (b *Box) PackEnd(child *Widget, expand, fill bool, padding uint) {
	var e, f C.gboolean
	if expand {
		e = 1
	}
	if fill {
		f = 1
	}
	C.gtk_box_pack_end(b.g(), child.g(), e, f, C.guint(padding))
}

func (b *Box) SetSpecing(spacing int) {
	C.gtk_box_set_spacing(b.g(), C.gint(spacing))
}

func (b *Box) GetSpecing() int {
	return int(C.gtk_box_get_spacing(b.g()))
}

type VBox struct {
	Box
}

func (b *VBox) g() *C.GtkVBox {
	return (*C.GtkVBox)(b.GetPtr())
}

func (b *VBox) AsVBox() *VBox {
	return b
}

func NewVBox(homogeneous bool, spacing int) *VBox {
	b := new(VBox)
	var h C.gboolean
	if homogeneous {
		h = 1
	}
	b.SetPtr(glib.Pointer(C.gtk_vbox_new(h, C.gint(spacing))))
	return b
}

type HBox struct {
	Box
}

func (b *HBox) g() *C.GtkHBox {
	return (*C.GtkHBox)(b.GetPtr())
}

func (b *HBox) AsHBox() *HBox {
	return b
}

func NewHBox(homogeneous bool, spacing int) *HBox {
	b := new(HBox)
	var h C.gboolean
	if homogeneous {
		h = 1
	}
	b.SetPtr(glib.Pointer(C.gtk_hbox_new(h, C.gint(spacing))))
	return b
}

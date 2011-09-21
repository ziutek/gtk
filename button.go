package gtk

/*
#include <gtk/gtk.h>
*/
import "C"

import (
	"github.com/ziutek/glib"
)

type Button struct {
	Bin
}

func (b *Button) g() *C.GtkButton {
	return (*C.GtkButton)(b.GetPtr())
}

func (b *Button) AsButton() *Button {
	return b
}

func (b *Button) SetLabel(label string) {
	C.gtk_button_set_label(b.g(), (*C.gchar)(C.CString(label)))
}

func (b *Button) GetLabel() string {
	 return C.GoString((*C.char)(C.gtk_button_get_label(b.g())))
}

func (b *Button) SetUseStock(use_stock bool) {
	var us C.gboolean	
	if use_stock {
		us = 1
	}
	C.gtk_button_set_use_stock(b.g(), us)
}

func (b *Button) GetUseStock() bool {
	return C.gtk_button_get_use_stock(b.g()) != 0
}

func NewButton() *Button {
	b := new(Button)
	b.SetPtr(glib.Pointer(C.gtk_button_new()))
	return b
}

func NewButtonWithLabel(label string) *Button {
	b := new(Button)
	b.SetPtr(
		glib.Pointer(C.gtk_button_new_with_label((*C.gchar)(C.CString(label)))),
	)
	return b
}

func NewButtonFromStock(stock string) *Button {
	b := new(Button)
	b.SetPtr(
		glib.Pointer(C.gtk_button_new_from_stock((*C.gchar)(C.CString(stock)))),
	)
	return b
}

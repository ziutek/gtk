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
	return (*C.GtkButton)(b.Pointer())
}

// Returns C pointer
func NewButton() *Button {
	b := new(Button)
	b.Set(glib.Pointer(C.gtk_button_new()))
	return b
}

// Returns C pointer
func NewButtonWithLabel(l string) *Button {
	b := new(Button)
	b.Set(glib.Pointer(C.gtk_button_new_with_label((*C.gchar)(C.CString(l)))))
	return b
}

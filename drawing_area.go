package gtk

/*
#include <gtk/gtk.h>
*/
import "C"

import (
	"github.com/ziutek/glib"
)

type DrawingArea struct {
	Widget
}

func (d *DrawingArea) g() *C.GtkDrawingArea {
	return (*C.GtkDrawingArea)(d.GetPtr())
}

func (d *DrawingArea) AsDrawingArea() *DrawingArea {
	return d
}

func NewDrawingArea() *DrawingArea {
	d := new(DrawingArea)
	d.SetPtr(glib.Pointer(C.gtk_drawing_area_new()))
	return d
}

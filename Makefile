include $(GOROOT)/src/Make.inc

TARG = github.com/ziutek/gtk
CGOFILES = functions.go \
	   widget.go \
	   container.go \
	   bin.go \
	   window.go \
	   button.go \
	   box.go \
	   entry.go \
	   drawing_area.go \
	   file_chooser.go \
	   file_chooser_button.go \
	   settings.go \

include $(GOROOT)/src/Make.pkg

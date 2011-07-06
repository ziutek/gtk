include $(GOROOT)/src/Make.inc

TARG = github.com/ziutek/gtk
CGOFILES = functions.go widget.go container.go bin.go window.go

include $(GOROOT)/src/Make.pkg

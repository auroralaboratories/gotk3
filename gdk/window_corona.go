package gdk

// #cgo pkg-config: gdk-3.0
// #include "gdk.go.h"
import "C"

import (
	"unsafe"
)

var gdkWindowFilters = make([]*FilterCallback, 0)

type XEvent struct {
	GdkXEvent *C.GdkXEvent
}

// native returns a pointer to the underlying GdkXEvent.
func (v *XEvent) native() *C.GdkXEvent {
	if v == nil {
		return nil
	}

	return v.GdkXEvent
}

// Native returns a pointer to the underlying GdkXEvent.
func (v *XEvent) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}


type FilterCallback struct {
	Window   *Window
	Callback FilterFunc
	UserData unsafe.Pointer
}

type FilterReturn int
const (
	FilterReturnContinue  FilterReturn = C.GDK_FILTER_CONTINUE
	FilterReturnTranslate              = C.GDK_FILTER_TRANSLATE
	FilterReturnRemove                 = C.GDK_FILTER_REMOVE
)

type FilterFunc func(xevent XEvent, event Event, userdata unsafe.Pointer) FilterReturn


func AddWindowEventFilter(filter FilterFunc, userdata unsafe.Pointer) *FilterCallback {
	newFilterCallback := &FilterCallback{
		Callback: filter,
		UserData: userdata,
	}

	gdkWindowFilters = append(gdkWindowFilters, newFilterCallback)

	C.gdk_window_add_filter(nil, (C.GdkFilterFunc)(C.gdk_window_filter_func_callback), C.gpointer(newFilterCallback))

	return newFilterCallback
}

func RemoveWindowEventFilter(filterCallback *FilterCallback) bool {
	for i, fcb := range gdkWindowFilters {
	//	if this filter matches the one we're removing...
		if fcb == filterCallback {
		//	call the remove
			C.gdk_window_remove_filter(nil, (C.GdkFilterFunc)(C.gdk_window_filter_func_callback), nil)

		//	remove the golang-side tracking element
			gdkWindowFilters = append(gdkWindowFilters[:i], gdkWindowFilters[(i+1):]...)

			return true
		}
	}

	return false
}


func (v *Window) AddEventFilter(filter FilterFunc, userdata unsafe.Pointer) *FilterCallback {
	newFilterCallback := &FilterCallback{
		Window:   v,
		Callback: filter,
		UserData: userdata,
	}

	gdkWindowFilters = append(gdkWindowFilters, newFilterCallback)

	C.gdk_window_add_filter(v.native(), (C.GdkFilterFunc)(C.gdk_window_filter_func_callback), C.gpointer(newFilterCallback))

	return newFilterCallback
}

func (v *Window) RemoveFilter(filterCallback *FilterCallback) bool {
	for i, fcb := range gdkWindowFilters {
	//	if this filter matches the one we're removing...
		if fcb == filterCallback {
		//	call the remove
			C.gdk_window_remove_filter(v.native(), (C.GdkFilterFunc)(C.gdk_window_filter_func_callback), nil)

		//	remove the golang-side tracking element
			gdkWindowFilters = append(gdkWindowFilters[:i], gdkWindowFilters[(i+1):]...)

			return true
		}
	}

	return false
}

//export go_genericGtkWindowFilterFuncCallback
func go_genericGtkWindowFilterFuncCallback(callback unsafe.Pointer, xevent *C.GdkXEvent, event *C.GdkEvent) int {
	for _, fcb := range gdkWindowFilters {
		if unsafe.Pointer(fcb) == callback {
			xev := XEvent{ xevent }
			gev := Event{ event }

			return int(fcb.Callback(xev, gev, fcb.UserData))
		}
	}

	return int(FilterReturnContinue)
}

package gdk

// #cgo pkg-config: gdk-3.0
// #include "gdk.go.h"
import "C"

import (
	"math/rand"
	"time"
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
	ID       uint32
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

func addWindowEventFilter(window *Window, filter FilterFunc, userdata unsafe.Pointer) *FilterCallback {
	var gdkWindow *C.GdkWindow

	if window != nil {
		gdkWindow = window.native()
	}

	newFilterCallback := &FilterCallback{
		ID:       rand.New(rand.NewSource(time.Now().UnixNano())).Uint32(),
		Window:   window,
		Callback: filter,
		UserData: userdata,
	}

	gdkWindowFilters = append(gdkWindowFilters, newFilterCallback)

	C.gdk_window_add_filter(gdkWindow, (C.GdkFilterFunc)(C.gdk_window_filter_func_callback), C.uint32_to_gpointer(C.uint32_t(newFilterCallback.ID)))

	return newFilterCallback
}

func removeWindowEventFilter(window *Window, filterCallback *FilterCallback) bool {
	var gdkWindow *C.GdkWindow

	if window != nil {
		gdkWindow = window.native()
	}

	for i, fcb := range gdkWindowFilters {
		//	if this filter matches the one we're removing...
		if fcb.ID == filterCallback.ID {
			//	call the remove
			C.gdk_window_remove_filter(gdkWindow, (C.GdkFilterFunc)(C.gdk_window_filter_func_callback), nil)

			//	remove the golang-side tracking element
			gdkWindowFilters = append(gdkWindowFilters[:i], gdkWindowFilters[(i+1):]...)

			return true
		}
	}

	return false
}

func AddGlobalEventFilter(filter FilterFunc, userdata unsafe.Pointer) *FilterCallback {
	return addWindowEventFilter(nil, filter, userdata)
}

func RemoveGlobalEventFilter(filterCallback *FilterCallback) bool {
	return removeWindowEventFilter(nil, filterCallback)
}

func (v *Window) AddEventFilter(filter FilterFunc, userdata unsafe.Pointer) *FilterCallback {
	return addWindowEventFilter(v, filter, userdata)
}

func (v *Window) RemoveFilter(filterCallback *FilterCallback) bool {
	return removeWindowEventFilter(v, filterCallback)
}

func (v *Window) GetEventMask() EventMask {
	return EventMask(C.gdk_window_get_events(v.native()))
}

func (v *Window) SetEventMask(mask EventMask) {
	C.gdk_window_set_events(v.native(), (C.GdkEventMask)(mask))
}

//export go_genericGtkWindowFilterFuncCallback
func go_genericGtkWindowFilterFuncCallback(filterID uint32, xevent *C.GdkXEvent, event *C.GdkEvent) int {
	for _, fcb := range gdkWindowFilters {
		if fcb.ID == filterID {
			xev := XEvent{xevent}
			gev := Event{event}

			return int(fcb.Callback(xev, gev, fcb.UserData))
		}
	}

	return int(FilterReturnContinue)
}

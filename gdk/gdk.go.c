#include "gdk.go.h"

GdkAtom toGdkAtom(void *p) {
	return ((GdkAtom)p);
}

GdkDevice* toGdkDevice(void *p) {
	return (GDK_DEVICE(p));
}

GdkCursor* toGdkCursor(void *p) {
	return (GDK_CURSOR(p));
}

GdkDeviceManager* toGdkDeviceManager(void *p) {
	return (GDK_DEVICE_MANAGER(p));
}

GdkDisplay* toGdkDisplay(void *p) {
	return (GDK_DISPLAY(p));
}

GdkDragContext* toGdkDragContext(void *p) {
	return (GDK_DRAG_CONTEXT(p));
}

GdkPixbuf* toGdkPixbuf(void *p) {
	return (GDK_PIXBUF(p));
}

gboolean _gdk_pixbuf_save_png(GdkPixbuf *pixbuf,
const char *filename, GError ** err, const char *compression) {
	return gdk_pixbuf_save(pixbuf, filename, "png", err, "compression", compression, NULL);
}

gboolean _gdk_pixbuf_save_jpeg(GdkPixbuf *pixbuf,
const char *filename, GError ** err, const char *quality) {
	return gdk_pixbuf_save(pixbuf, filename, "jpeg", err, "quality", quality, NULL);
}

GdkPixbufLoader* toGdkPixbufLoader(void *p) {
	return (GDK_PIXBUF_LOADER(p));
}

GdkScreen* toGdkScreen(void *p) {
	return (GDK_SCREEN(p));
}

GdkVisual* toGdkVisual(void *p) {
	return (GDK_VISUAL(p));
}

GdkWindow* toGdkWindow(void *p) {
	return (GDK_WINDOW(p));
}

gpointer uint32_to_gpointer(uint32_t in) {
	return GUINT_TO_POINTER(in);
}

GdkFilterReturn gdk_window_filter_func_callback(GdkXEvent *xevent, GdkEvent *event, gpointer goFilterID) {
	XEvent *xlibEvent = (XEvent*)(xevent);

	return (GdkFilterReturn)(go_genericGtkWindowFilterFuncCallback((uint32_t)(GPOINTER_TO_UINT(goFilterID)), xevent, event));
}

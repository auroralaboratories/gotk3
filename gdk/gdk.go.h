#ifndef GDK_GO_H
#define GDK_GO_H
/*
 * Copyright (c) 2013-2014 Conformal Systems <info@conformal.com>
 *
 * This file originated from: http://opensource.conformal.com/
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

#include <stdlib.h>
#include <stdint.h>
#include <X11/Xlib.h>
#include <glib.h>
#include <gdk/gdk.h>


gboolean          _gdk_pixbuf_save_jpeg(GdkPixbuf *pixbuf, const char *filename, GError ** err, const char *quality);
gboolean          _gdk_pixbuf_save_png(GdkPixbuf *pixbuf, const char *filename, GError ** err, const char *compression);
GdkAtom           toGdkAtom(void *p);
GdkCursor*        toGdkCursor(void *p);
GdkDevice*        toGdkDevice(void *p);
GdkDeviceManager* toGdkDeviceManager(void *p);
GdkDisplay*       toGdkDisplay(void *p);
GdkDragContext*   toGdkDragContext(void *p);
GdkPixbuf*        toGdkPixbuf(void *p);
GdkPixbufLoader*  toGdkPixbufLoader(void *p);
GdkScreen*        toGdkScreen(void *p);
GdkVisual*        toGdkVisual(void *p);
GdkWindow*        toGdkWindow(void *p);
GdkWindow*        toGdkWindow(void *p);
gpointer          uint32_to_gpointer(uint32_t);
GdkFilterReturn   gdk_window_filter_func_callback(GdkXEvent *xevent, GdkEvent *event, gpointer goFilterID);

#endif

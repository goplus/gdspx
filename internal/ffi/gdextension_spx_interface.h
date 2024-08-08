#ifndef GDEXTENSION_SPX_INTERFACE_H
#define GDEXTENSION_SPX_INTERFACE_H

#include <stddef.h>
#include <stdint.h>
#include <stdlib.h>
#include "gdextension_interface.h"

#ifndef __cplusplus
typedef uint32_t char32_t;
typedef uint16_t char16_t;
#endif

#ifdef __cplusplus
extern "C" {
#endif

typedef uintptr_t pointer;
typedef const char* string;
typedef char32_t rune;
typedef uint8_t * bytes;
typedef GDExtensionBool bool;

extern void initialize(void *userdata, GDExtensionInitializationLevel p_level);
extern void deinitialize(void *userdata, GDExtensionInitializationLevel p_level);

static inline void initialization(GDExtensionInitialization *p_init) {
	p_init->initialize = initialize;
	p_init->deinitialize = deinitialize;
}

static inline void *get_proc_address(pointer fn, string p_name) {
	return (void *)((GDExtensionInterfaceGetProcAddress)fn)(p_name);
}
static inline void print_error(pointer fn, string p_description, string p_function, string p_file, int32_t p_line, GDExtensionBool p_notify_editor) {
	((GDExtensionInterfacePrintError)fn)(p_description, p_function, p_file, p_line, p_notify_editor);
}


#ifdef __cplusplus
}
#endif

#endif // GDEXTENSION_SPX_INTERFACE_H
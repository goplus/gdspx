#ifndef GDEXTENSION_SPX_WRAP_H
#define GDEXTENSION_SPX_WRAP_H

#include <stddef.h>
#include <stdint.h>
#include <stdlib.h>
#define NOT_GODOT_ENGINE
#include "gdextension_spx_ext.h"
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

extern void func_on_engine_start();  
extern void func_on_engine_update(gdfloat delta);  
extern void func_on_engine_destroy();  
extern void func_on_sprite_ready(gdint id);  
extern void func_on_sprite_updated(gdint id);  
extern void func_on_sprite_destroyed(gdint id);  
extern void func_on_mouse_pressed(gdint keyid);  
extern void func_on_mouse_released(gdint keyid);  
extern void func_on_key_pressed(gdint keyid);  
extern void func_on_key_released(gdint keyid);  
extern void func_on_action_pressed(gdstring action_name);  
extern void func_on_action_just_pressed(gdstring action_name);  
extern void func_on_action_just_released(gdstring action_name);  
extern void func_on_axis_changed(gdstring action_name, gdfloat value);  
extern void func_on_collision_enter(gdint self_id, gdint other_id);  
extern void func_on_collision_stay(gdint self_id, gdint other_id);  
extern void func_on_collision_exit(gdint self_id, gdint other_id);  
extern void func_on_trigger_enter(gdint self_id, gdint other_id);  
extern void func_on_trigger_stay(gdint self_id, gdint other_id);  
extern void func_on_trigger_exit(gdint self_id, gdint other_id);  
extern void func_on_ui_pressed(gdint id);  
extern void func_on_ui_released(gdint id);  
extern void func_on_ui_hovered(gdint id);  
extern void func_on_ui_clicked(gdint id);  
extern void func_on_ui_toggle(gdint id, gdbool is_on);  
extern void func_on_ui_text_changed(gdint id, gdstring text);  

static inline void spx_global_register_callbacks(pointer fn) {
	SpxCallbackInfo info;
	SpxCallbackInfo* p_extension_funcs = &info;
	p_extension_funcs->func_on_engine_start = func_on_engine_start;
	p_extension_funcs->func_on_engine_update = func_on_engine_update;
	p_extension_funcs->func_on_engine_destroy = func_on_engine_destroy;
	p_extension_funcs->func_on_sprite_ready = func_on_sprite_ready;
	p_extension_funcs->func_on_sprite_updated = func_on_sprite_updated;
	p_extension_funcs->func_on_sprite_destroyed = func_on_sprite_destroyed;
	p_extension_funcs->func_on_mouse_pressed = func_on_mouse_pressed;
	p_extension_funcs->func_on_mouse_released = func_on_mouse_released;
	p_extension_funcs->func_on_key_pressed = func_on_key_pressed;
	p_extension_funcs->func_on_key_released = func_on_key_released;
	p_extension_funcs->func_on_action_pressed = func_on_action_pressed;
	p_extension_funcs->func_on_action_just_pressed = func_on_action_just_pressed;
	p_extension_funcs->func_on_action_just_released = func_on_action_just_released;
	p_extension_funcs->func_on_axis_changed = func_on_axis_changed;
	p_extension_funcs->func_on_collision_enter = func_on_collision_enter;
	p_extension_funcs->func_on_collision_stay = func_on_collision_stay;
	p_extension_funcs->func_on_collision_exit = func_on_collision_exit;
	p_extension_funcs->func_on_trigger_enter = func_on_trigger_enter;
	p_extension_funcs->func_on_trigger_stay = func_on_trigger_stay;
	p_extension_funcs->func_on_trigger_exit = func_on_trigger_exit;
	p_extension_funcs->func_on_ui_pressed = func_on_ui_pressed;
	p_extension_funcs->func_on_ui_released = func_on_ui_released;
	p_extension_funcs->func_on_ui_hovered = func_on_ui_hovered;
	p_extension_funcs->func_on_ui_clicked = func_on_ui_clicked;
	p_extension_funcs->func_on_ui_toggle = func_on_ui_toggle;
	p_extension_funcs->func_on_ui_text_changed = func_on_ui_text_changed;
	((GDExtensionSpxGlobalRegisterCallbacks)fn)((GDExtensionSpxCallbackInfoPtr)p_extension_funcs);
}



#ifdef __cplusplus
}
#endif

#endif // GDEXTENSION_SPX_WRAP_H
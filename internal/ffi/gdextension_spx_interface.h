#ifndef GDEXTENSION_SPX_INTERFACE_H
#define GDEXTENSION_SPX_INTERFACE_H

// use 64-bit real_t
#define REAL_T_IS_DOUBLE 1 
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

#ifdef REAL_T_IS_DOUBLE
typedef double real_t;
#else
typedef float real_t;
#endif

typedef struct {
    real_t x;
    real_t y;
    real_t z;
    real_t w;
} Vector4;

typedef struct {
    real_t x;
    real_t y;
    real_t z;
} Vector3;

typedef struct {
    real_t x;
    real_t y;
} Vector2;

typedef struct {
    float x;
    float y;
    float z;
    float w;
} Color;

typedef struct {
	Vector2 position;
	Vector2 size;
} Rect2;

typedef real_t GDReal;

#ifdef __cplusplus
}
#endif

#define NOT_GODOT_ENGINE
#include "gdextension_spx_ext.h"
#include "gdextension_spx_wrap.h"
 
#endif // GDEXTENSION_SPX_INTERFACE_H
#ifndef GDEXTENSION_SPX_PRE_DEFINE_H
#define GDEXTENSION_SPX_PRE_DEFINE_H

#include <stddef.h>
#include <stdint.h>
#include <stdlib.h>
#ifndef __cplusplus
typedef uint32_t char32_t;
typedef uint16_t char16_t;
#endif

#ifdef __cplusplus
extern "C" {
#endif


typedef double real_t;

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

typedef void *GDExtensionStringPtr;
typedef int64_t GDExtensionInt;
typedef uint8_t GDExtensionBool;


#ifdef __cplusplus
}
#endif

#endif // GDEXTENSION_SPX_PRE_DEFINE_H
#pragma once

#include <stdint.h>
#include "include/jni.h"

#define ARG_TYPE_UINT32     0
#define ARG_TYPE_UINT64     1
#define ARG_TYPE_DOUBLE     2
#define ARG_TYPE_PTR        3
#define ARG_TYPE_JCLASS     4
#define ARG_TYPE_JOBJECT    5
#define ARG_TYPE_JSTRING    6

/**
 *  MethodArgType represents the gojni.TypedMethodArgs in C side
 *  it is used for representing a single JNI Function / Method Argument
 **/
typedef struct MethodArgType {
   int type;
   void *ptrValue;
   uint64_t intValue;
   double doubleValue;
} MethodArgType;

void *GetJNIEnv();
void *GetJVM();

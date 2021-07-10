#pragma once
#include "types.h"

extern jint        JVM_DestroyJavaVM(JavaVM*);
extern jint        JVM_AttachCurrentThread(JavaVM*, JNIEnv**, void*);
extern jint        JVM_DetachCurrentThread(JavaVM*);
extern jint        JVM_GetEnv(JavaVM*, void**, jint);
extern jint        JVM_AttachCurrentThreadAsDaemon(JavaVM*, JNIEnv**, void*);
extern void        JVM_LogDebug(const char *);

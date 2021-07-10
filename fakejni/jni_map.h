#pragma once
#include "types.h"

extern jint             JNI_GetVersion(JNIEnv *);

extern jclass           JNI_DefineClass(JNIEnv*, const char*, jobject, const jbyte*, jsize);
extern jclass           JNI_FindClass(JNIEnv*, const char*);

extern jmethodID        JNI_FromReflectedMethod(JNIEnv*, jobject);
extern jfieldID         JNI_FromReflectedField(JNIEnv*, jobject);
/* spec doesn't show jboolean parameter */
extern jobject          JNI_ToReflectedMethod(JNIEnv*, jclass, jmethodID, jboolean);

extern jclass           JNI_GetSuperclass(JNIEnv*, jclass);
extern jboolean         JNI_IsAssignableFrom(JNIEnv*, jclass, jclass);

/* spec doesn't show jboolean parameter */
extern jobject          JNI_ToReflectedField(JNIEnv*, jclass, jfieldID, jboolean);

extern jint             JNI_Throw(JNIEnv*, jthrowable);
extern jint             JNI_ThrowNew(JNIEnv *, jclass, const char *);
extern jthrowable       JNI_ExceptionOccurred(JNIEnv*);
extern void             JNI_ExceptionDescribe(JNIEnv*);
extern void             JNI_ExceptionClear(JNIEnv*);
extern void             JNI_FatalError(JNIEnv*, const char*);

extern jint             JNI_PushLocalFrame(JNIEnv*, jint);
extern jobject          JNI_PopLocalFrame(JNIEnv*, jobject);

extern jobject          JNI_NewGlobalRef(JNIEnv*, jobject);
extern void             JNI_DeleteGlobalRef(JNIEnv*, jobject);
extern void             JNI_DeleteLocalRef(JNIEnv*, jobject);
extern jboolean         JNI_IsSameObject(JNIEnv*, jobject, jobject);

extern jobject          JNI_NewLocalRef(JNIEnv*, jobject);
extern jint             JNI_EnsureLocalCapacity(JNIEnv*, jint);

extern jobject          JNI_AllocObject(JNIEnv*, jclass);
extern jobject          JNI_NewObject(JNIEnv*, jclass, jmethodID, ...);
extern jobject          JNI_NewObjectV(JNIEnv*, jclass, jmethodID, va_list);
extern jobject          JNI_NewObjectA(JNIEnv*, jclass, jmethodID, const jvalue*);

extern jclass           JNI_GetObjectClass(JNIEnv*, jobject);
extern jboolean         JNI_IsInstanceOf(JNIEnv*, jobject, jclass);
extern jmethodID        JNI_GetMethodID(JNIEnv*, jclass, const char*, const char*);

extern jobject          JNI_CallObjectMethodV(JNIEnv*, jobject, jmethodID, va_list);
extern jobject          JNI_CallObjectMethodA(JNIEnv*, jobject, jmethodID, const jvalue*);
extern jboolean         JNI_CallBooleanMethod(JNIEnv*, jobject, jmethodID, ...);
extern jboolean         JNI_CallBooleanMethodV(JNIEnv*, jobject, jmethodID, va_list);
extern jboolean         JNI_CallBooleanMethodA(JNIEnv*, jobject, jmethodID, const jvalue*);
extern jbyte            JNI_CallByteMethod(JNIEnv*, jobject, jmethodID, ...);
extern jbyte            JNI_CallByteMethodV(JNIEnv*, jobject, jmethodID, va_list);
extern jbyte            JNI_CallByteMethodA(JNIEnv*, jobject, jmethodID, const jvalue*);
extern jchar            JNI_CallCharMethod(JNIEnv*, jobject, jmethodID, ...);
extern jchar            JNI_CallCharMethodV(JNIEnv*, jobject, jmethodID, va_list);
extern jchar            JNI_CallCharMethodA(JNIEnv*, jobject, jmethodID, const jvalue*);
extern jshort           JNI_CallShortMethod(JNIEnv*, jobject, jmethodID, ...);
extern jshort           JNI_CallShortMethodV(JNIEnv*, jobject, jmethodID, va_list);
extern jshort           JNI_CallShortMethodA(JNIEnv*, jobject, jmethodID, const jvalue*);
extern jint             JNI_CallIntMethod(JNIEnv*, jobject, jmethodID, ...);
extern jint             JNI_CallIntMethodV(JNIEnv*, jobject, jmethodID, va_list);
extern jint             JNI_CallIntMethodA(JNIEnv*, jobject, jmethodID, const jvalue*);
extern jlong            JNI_CallLongMethod(JNIEnv*, jobject, jmethodID, ...);
extern jlong            JNI_CallLongMethodV(JNIEnv*, jobject, jmethodID, va_list);
extern jlong            JNI_CallLongMethodA(JNIEnv*, jobject, jmethodID, const jvalue*);
extern jfloat           JNI_CallFloatMethod(JNIEnv*, jobject, jmethodID, ...);
extern jfloat           JNI_CallFloatMethodV(JNIEnv*, jobject, jmethodID, va_list);
extern jfloat           JNI_CallFloatMethodA(JNIEnv*, jobject, jmethodID, const jvalue*);
extern jdouble          JNI_CallDoubleMethod(JNIEnv*, jobject, jmethodID, ...);
extern jdouble          JNI_CallDoubleMethodV(JNIEnv*, jobject, jmethodID, va_list);
extern jdouble          JNI_CallDoubleMethodA(JNIEnv*, jobject, jmethodID, const jvalue*);
extern void             JNI_CallVoidMethod(JNIEnv*, jobject, jmethodID, ...);
extern void             JNI_CallVoidMethodV(JNIEnv*, jobject, jmethodID, va_list);
extern void             JNI_CallVoidMethodA(JNIEnv*, jobject, jmethodID, const jvalue*);

extern jobject          JNI_CallNonvirtualObjectMethod(JNIEnv*, jobject, jclass, jmethodID, ...);
extern jobject          JNI_CallNonvirtualObjectMethodV(JNIEnv*, jobject, jclass, jmethodID, va_list);
extern jobject          JNI_CallNonvirtualObjectMethodA(JNIEnv*, jobject, jclass, jmethodID, const jvalue*);
extern jboolean         JNI_CallNonvirtualBooleanMethod(JNIEnv*, jobject, jclass, jmethodID, ...);
extern jboolean         JNI_CallNonvirtualBooleanMethodV(JNIEnv*, jobject, jclass,  jmethodID, va_list);
extern jboolean         JNI_CallNonvirtualBooleanMethodA(JNIEnv*, jobject, jclass,  jmethodID, const jvalue*);
extern jbyte            JNI_CallNonvirtualByteMethod(JNIEnv*, jobject, jclass, jmethodID, ...);
extern jbyte            JNI_CallNonvirtualByteMethodV(JNIEnv*, jobject, jclass, jmethodID, va_list);
extern jbyte            JNI_CallNonvirtualByteMethodA(JNIEnv*, jobject, jclass, jmethodID, const jvalue*);
extern jchar            JNI_CallNonvirtualCharMethod(JNIEnv*, jobject, jclass, jmethodID, ...);
extern jchar            JNI_CallNonvirtualCharMethodV(JNIEnv*, jobject, jclass, jmethodID, va_list);
extern jchar            JNI_CallNonvirtualCharMethodA(JNIEnv*, jobject, jclass, jmethodID, const jvalue*);
extern jshort           JNI_CallNonvirtualShortMethod(JNIEnv*, jobject, jclass, jmethodID, ...);
extern jshort           JNI_CallNonvirtualShortMethodV(JNIEnv*, jobject, jclass, jmethodID, va_list);
extern jshort           JNI_CallNonvirtualShortMethodA(JNIEnv*, jobject, jclass, jmethodID, const jvalue*);
extern jint             JNI_CallNonvirtualIntMethod(JNIEnv*, jobject, jclass, jmethodID, ...);
extern jint             JNI_CallNonvirtualIntMethodV(JNIEnv*, jobject, jclass, jmethodID, va_list);
extern jint             JNI_CallNonvirtualIntMethodA(JNIEnv*, jobject, jclass, jmethodID, const jvalue*);
extern jlong            JNI_CallNonvirtualLongMethod(JNIEnv*, jobject, jclass, jmethodID, ...);
extern jlong            JNI_CallNonvirtualLongMethodV(JNIEnv*, jobject, jclass, jmethodID, va_list);
extern jlong            JNI_CallNonvirtualLongMethodA(JNIEnv*, jobject, jclass, jmethodID, const jvalue*);
extern jfloat           JNI_CallNonvirtualFloatMethod(JNIEnv*, jobject, jclass, jmethodID, ...);
extern jfloat           JNI_CallNonvirtualFloatMethodV(JNIEnv*, jobject, jclass, jmethodID, va_list);
extern jfloat           JNI_CallNonvirtualFloatMethodA(JNIEnv*, jobject, jclass, jmethodID, const jvalue*);
extern jdouble          JNI_CallNonvirtualDoubleMethod(JNIEnv*, jobject, jclass, jmethodID, ...);
extern jdouble          JNI_CallNonvirtualDoubleMethodV(JNIEnv*, jobject, jclass, jmethodID, va_list);
extern jdouble          JNI_CallNonvirtualDoubleMethodA(JNIEnv*, jobject, jclass, jmethodID, const jvalue*);
extern void             JNI_CallNonvirtualVoidMethod(JNIEnv*, jobject, jclass, jmethodID, ...);
extern void             JNI_CallNonvirtualVoidMethodV(JNIEnv*, jobject, jclass, jmethodID, va_list);
extern void             JNI_CallNonvirtualVoidMethodA(JNIEnv*, jobject, jclass, jmethodID, const jvalue*);

extern jfieldID         JNI_GetFieldID(JNIEnv*, jclass, const char*, const char*);

extern jobject          JNI_GetObjectField(JNIEnv*, jobject, jfieldID);
extern jboolean         JNI_GetBooleanField(JNIEnv*, jobject, jfieldID);
extern jbyte            JNI_GetByteField(JNIEnv*, jobject, jfieldID);
extern jchar            JNI_GetCharField(JNIEnv*, jobject, jfieldID);
extern jshort           JNI_GetShortField(JNIEnv*, jobject, jfieldID);
extern jint             JNI_GetIntField(JNIEnv*, jobject, jfieldID);
extern jlong            JNI_GetLongField(JNIEnv*, jobject, jfieldID);
extern jfloat           JNI_GetFloatField(JNIEnv*, jobject, jfieldID);
extern jdouble          JNI_GetDoubleField(JNIEnv*, jobject, jfieldID);

extern void             JNI_SetObjectField(JNIEnv*, jobject, jfieldID, jobject);
extern void             JNI_SetBooleanField(JNIEnv*, jobject, jfieldID, jboolean);
extern void             JNI_SetByteField(JNIEnv*, jobject, jfieldID, jbyte);
extern void             JNI_SetCharField(JNIEnv*, jobject, jfieldID, jchar);
extern void             JNI_SetShortField(JNIEnv*, jobject, jfieldID, jshort);
extern void             JNI_SetIntField(JNIEnv*, jobject, jfieldID, jint);
extern void             JNI_SetLongField(JNIEnv*, jobject, jfieldID, jlong);
extern void             JNI_SetFloatField(JNIEnv*, jobject, jfieldID, jfloat);
extern void             JNI_SetDoubleField(JNIEnv*, jobject, jfieldID, jdouble);

extern jmethodID        JNI_GetStaticMethodID(JNIEnv*, jclass, const char*, const char*);

extern jobject          JNI_CallStaticObjectMethodV(JNIEnv*, jclass, jmethodID, va_list);
extern jobject          JNI_CallStaticObjectMethodA(JNIEnv*, jclass, jmethodID, const jvalue*);
extern jboolean         JNI_CallStaticBooleanMethod(JNIEnv*, jclass, jmethodID, ...);
extern jboolean         JNI_CallStaticBooleanMethodV(JNIEnv*, jclass, jmethodID, va_list);
extern jboolean         JNI_CallStaticBooleanMethodA(JNIEnv*, jclass, jmethodID, const jvalue*);
extern jbyte            JNI_CallStaticByteMethod(JNIEnv*, jclass, jmethodID, ...);
extern jbyte            JNI_CallStaticByteMethodV(JNIEnv*, jclass, jmethodID, va_list);
extern jbyte            JNI_CallStaticByteMethodA(JNIEnv*, jclass, jmethodID, const jvalue*);
extern jchar            JNI_CallStaticCharMethod(JNIEnv*, jclass, jmethodID, ...);
extern jchar            JNI_CallStaticCharMethodV(JNIEnv*, jclass, jmethodID, va_list);
extern jchar            JNI_CallStaticCharMethodA(JNIEnv*, jclass, jmethodID, const jvalue*);
extern jshort           JNI_CallStaticShortMethod(JNIEnv*, jclass, jmethodID, ...);
extern jshort           JNI_CallStaticShortMethodV(JNIEnv*, jclass, jmethodID, va_list);
extern jshort           JNI_CallStaticShortMethodA(JNIEnv*, jclass, jmethodID, const jvalue*);
extern jint             JNI_CallStaticIntMethod(JNIEnv*, jclass, jmethodID, ...);
extern jint             JNI_CallStaticIntMethodV(JNIEnv*, jclass, jmethodID, va_list);
extern jint             JNI_CallStaticIntMethodA(JNIEnv*, jclass, jmethodID, const jvalue*);
extern jlong            JNI_CallStaticLongMethod(JNIEnv*, jclass, jmethodID, ...);
extern jlong            JNI_CallStaticLongMethodV(JNIEnv*, jclass, jmethodID, va_list);
extern jlong            JNI_CallStaticLongMethodA(JNIEnv*, jclass, jmethodID, const jvalue*);
extern jfloat           JNI_CallStaticFloatMethod(JNIEnv*, jclass, jmethodID, ...);
extern jfloat           JNI_CallStaticFloatMethodV(JNIEnv*, jclass, jmethodID, va_list);
extern jfloat           JNI_CallStaticFloatMethodA(JNIEnv*, jclass, jmethodID, const jvalue*);
extern jdouble          JNI_CallStaticDoubleMethod(JNIEnv*, jclass, jmethodID, ...);
extern jdouble          JNI_CallStaticDoubleMethodV(JNIEnv*, jclass, jmethodID, va_list);
extern jdouble          JNI_CallStaticDoubleMethodA(JNIEnv*, jclass, jmethodID, const jvalue*);
extern void             JNI_CallStaticVoidMethod(JNIEnv*, jclass, jmethodID, ...);
extern void             JNI_CallStaticVoidMethodV(JNIEnv*, jclass, jmethodID, va_list);
extern void             JNI_CallStaticVoidMethodA(JNIEnv*, jclass, jmethodID, const jvalue*);

extern jfieldID         JNI_GetStaticFieldID(JNIEnv*, jclass, const char*, const char*);

extern jobject          JNI_GetStaticObjectField(JNIEnv*, jclass, jfieldID);
extern jboolean         JNI_GetStaticBooleanField(JNIEnv*, jclass, jfieldID);
extern jbyte            JNI_GetStaticByteField(JNIEnv*, jclass, jfieldID);
extern jchar            JNI_GetStaticCharField(JNIEnv*, jclass, jfieldID);
extern jshort           JNI_GetStaticShortField(JNIEnv*, jclass, jfieldID);
extern jint             JNI_GetStaticIntField(JNIEnv*, jclass, jfieldID);
extern jlong            JNI_GetStaticLongField(JNIEnv*, jclass, jfieldID);
extern jfloat           JNI_GetStaticFloatField(JNIEnv*, jclass, jfieldID);
extern jdouble          JNI_GetStaticDoubleField(JNIEnv*, jclass, jfieldID);

extern void             JNI_SetStaticObjectField(JNIEnv*, jclass, jfieldID, jobject);
extern void             JNI_SetStaticBooleanField(JNIEnv*, jclass, jfieldID, jboolean);
extern void             JNI_SetStaticByteField(JNIEnv*, jclass, jfieldID, jbyte);
extern void             JNI_SetStaticCharField(JNIEnv*, jclass, jfieldID, jchar);
extern void             JNI_SetStaticShortField(JNIEnv*, jclass, jfieldID, jshort);
extern void             JNI_SetStaticIntField(JNIEnv*, jclass, jfieldID, jint);
extern void             JNI_SetStaticLongField(JNIEnv*, jclass, jfieldID, jlong);
extern void             JNI_SetStaticFloatField(JNIEnv*, jclass, jfieldID, jfloat);
extern void             JNI_SetStaticDoubleField(JNIEnv*, jclass, jfieldID, jdouble);

extern jstring          JNI_NewString(JNIEnv*, const jchar*, jsize);
extern jsize            JNI_GetStringLength(JNIEnv*, jstring);
extern const jchar*     JNI_GetStringChars(JNIEnv*, jstring, jboolean*);
extern void             JNI_ReleaseStringChars(JNIEnv*, jstring, const jchar*);
extern jstring          JNI_NewStringUTF(JNIEnv*, const char*);
extern jsize            JNI_GetStringUTFLength(JNIEnv*, jstring);
/* JNI spec says this returns const jbyte*, but that's inconsistent */
extern const char*      JNI_GetStringUTFChars(JNIEnv*, jstring, jboolean*);
extern void             JNI_ReleaseStringUTFChars(JNIEnv*, jstring, const char*);
extern jsize            JNI_GetArrayLength(JNIEnv*, jarray);
extern jobjectArray     JNI_NewObjectArray(JNIEnv*, jsize, jclass, jobject);
extern jobject          JNI_GetObjectArrayElement(JNIEnv*, jobjectArray, jsize);
extern void             JNI_SetObjectArrayElement(JNIEnv*, jobjectArray, jsize, jobject);

extern jbooleanArray    JNI_NewBooleanArray(JNIEnv*, jsize);
extern jbyteArray       JNI_NewByteArray(JNIEnv*, jsize);
extern jcharArray       JNI_NewCharArray(JNIEnv*, jsize);
extern jshortArray      JNI_NewShortArray(JNIEnv*, jsize);
extern jintArray        JNI_NewIntArray(JNIEnv*, jsize);
extern jlongArray       JNI_NewLongArray(JNIEnv*, jsize);
extern jfloatArray      JNI_NewFloatArray(JNIEnv*, jsize);
extern jdoubleArray     JNI_NewDoubleArray(JNIEnv*, jsize);

extern jboolean*        JNI_GetBooleanArrayElements(JNIEnv*, jbooleanArray, jboolean*);
extern jbyte*           JNI_GetByteArrayElements(JNIEnv*, jbyteArray, jboolean*);
extern jchar*           JNI_GetCharArrayElements(JNIEnv*, jcharArray, jboolean*);
extern jshort*          JNI_GetShortArrayElements(JNIEnv*, jshortArray, jboolean*);
extern jint*            JNI_GetIntArrayElements(JNIEnv*, jintArray, jboolean*);
extern jlong*           JNI_GetLongArrayElements(JNIEnv*, jlongArray, jboolean*);
extern jfloat*          JNI_GetFloatArrayElements(JNIEnv*, jfloatArray, jboolean*);
extern jdouble*         JNI_GetDoubleArrayElements(JNIEnv*, jdoubleArray, jboolean*);

extern void             JNI_ReleaseBooleanArrayElements(JNIEnv*, jbooleanArray, jboolean*, jint);
extern void             JNI_ReleaseByteArrayElements(JNIEnv*, jbyteArray, jbyte*, jint);
extern void             JNI_ReleaseCharArrayElements(JNIEnv*, jcharArray, jchar*, jint);
extern void             JNI_ReleaseShortArrayElements(JNIEnv*, jshortArray, jshort*, jint);
extern void             JNI_ReleaseIntArrayElements(JNIEnv*, jintArray, jint*, jint);
extern void             JNI_ReleaseLongArrayElements(JNIEnv*, jlongArray, jlong*, jint);
extern void             JNI_ReleaseFloatArrayElements(JNIEnv*, jfloatArray, jfloat*, jint);
extern void             JNI_ReleaseDoubleArrayElements(JNIEnv*, jdoubleArray, jdouble*, jint);

extern void             JNI_GetBooleanArrayRegion(JNIEnv*, jbooleanArray, jsize, jsize, jboolean*);
extern void             JNI_GetByteArrayRegion(JNIEnv*, jbyteArray, jsize, jsize, jbyte*);
extern void             JNI_GetCharArrayRegion(JNIEnv*, jcharArray, jsize, jsize, jchar*);
extern void             JNI_GetShortArrayRegion(JNIEnv*, jshortArray, jsize, jsize, jshort*);
extern void             JNI_GetIntArrayRegion(JNIEnv*, jintArray, jsize, jsize, jint*);
extern void             JNI_GetLongArrayRegion(JNIEnv*, jlongArray, jsize, jsize, jlong*);
extern void             JNI_GetFloatArrayRegion(JNIEnv*, jfloatArray, jsize, jsize, jfloat*);
extern void             JNI_GetDoubleArrayRegion(JNIEnv*, jdoubleArray, jsize, jsize, jdouble*);

/* spec shows these without const; some jni.h do, some don't */
extern void             JNI_SetBooleanArrayRegion(JNIEnv*, jbooleanArray, jsize, jsize, const jboolean*);
extern void             JNI_SetByteArrayRegion(JNIEnv*, jbyteArray, jsize, jsize, const jbyte*);
extern void             JNI_SetCharArrayRegion(JNIEnv*, jcharArray, jsize, jsize, const jchar*);
extern void             JNI_SetShortArrayRegion(JNIEnv*, jshortArray, jsize, jsize, const jshort*);
extern void             JNI_SetIntArrayRegion(JNIEnv*, jintArray, jsize, jsize, const jint*);
extern void             JNI_SetLongArrayRegion(JNIEnv*, jlongArray, jsize, jsize, const jlong*);
extern void             JNI_SetFloatArrayRegion(JNIEnv*, jfloatArray, jsize, jsize, const jfloat*);
extern void             JNI_SetDoubleArrayRegion(JNIEnv*, jdoubleArray, jsize, jsize, const jdouble*);

extern jint             JNI_RegisterNatives(JNIEnv*, jclass, const JNINativeMethod*, jint);
extern jint             JNI_UnregisterNatives(JNIEnv*, jclass);
extern jint             JNI_MonitorEnter(JNIEnv*, jobject);
extern jint             JNI_MonitorExit(JNIEnv*, jobject);
extern jint             JNI_GetJavaVM(JNIEnv*, JavaVM**);

extern void             JNI_GetStringRegion(JNIEnv*, jstring, jsize, jsize, jchar*);
extern void             JNI_GetStringUTFRegion(JNIEnv*, jstring, jsize, jsize, char*);

extern void*            JNI_GetPrimitiveArrayCritical(JNIEnv*, jarray, jboolean*);
extern void             JNI_ReleasePrimitiveArrayCritical(JNIEnv*, jarray, void*, jint);

extern const jchar*     JNI_GetStringCritical(JNIEnv*, jstring, jboolean*);
extern void             JNI_ReleaseStringCritical(JNIEnv*, jstring, const jchar*);

extern jweak            JNI_NewWeakGlobalRef(JNIEnv*, jobject);
extern void             JNI_DeleteWeakGlobalRef(JNIEnv*, jweak);

extern jboolean         JNI_ExceptionCheck(JNIEnv*);

extern jobject          JNI_NewDirectByteBuffer(JNIEnv*, void*, jlong);
extern void*            JNI_GetDirectBufferAddress(JNIEnv*, jobject);
extern jlong            JNI_GetDirectBufferCapacity(JNIEnv*, jobject);

/* added in JNI 1.6 */
extern jobjectRefType   JNI_GetObjectRefType(JNIEnv*, jobject);

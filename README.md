# FakeJNI
Golang FakeJNI for calling JNI Libraries from golang (with android support)

# Example

See `example` folder

# TODO

* Implement remaining JNI Methods
* Cross-platform (now it only works on linux / android)
* Better documentation
* Better API 
* Better README

# How it works

The library works by redirecting **ALL** JNI Calls that a JNI library / app can call to golang side, and dealing with that
inside golang. The project is subdivided in few parts:

* *gojni* => A Golang mock implementation of JNI. This deals with high level stuff like allocating objects, calling methods and so
* *fakejni* => The proxy call from native JNI C to *gojni* package. This implements all methods from JNI and forwards to a singleton from *gojni* package
* *android* => A small implementation of android classes mock. It doesnt do much but implements some functionality.
* *java* => A really small implementation of basic java classes

All the calls are currently being redirected, but not all of them actually have a implementation. 
You can check the implemented JNI calls at [fakejni/fakejni.go](fakejni/fakejni.go) and the uninimplemented calls at [fakejni/fakejni_nimpl.go](fakejni/fakejni_nimpl.go) 

# Adding a class mock

You can add java classes mocks in gojni during the runtime by calling `AddClass` method with the class name. 

```go 
jni :=  gojni.GetJNI()
androidContextClass := jni.AddClass("android/content/Context")
androidContextClass.
    AddStaticFieldGetter("TELEPHONY_SERVICE", androidContentContextGetTELEPHONY_SERVICE).
    AddStaticFieldGetter("WIFI_SERVICE", androidContentContextGetWIFI_SERVICE).
    AddMethod("getSystemService", "(Ljava/lang/String;)Ljava/lang/Object;", androidContentContextGetSystemService).
    AddMethod("getContentResolver", "()Landroid/content/ContentResolver;", androidContentContextGetContentResolver)
```

The handler function is called when the action is called by JNI. The handler is called with the JNI State and arguments from the caller.
For a more detailed example, check the `android` and `java` packages.

# Android

Android uses a different linker than linux, so the binaries are "incompatible" with linux. 
Luckily golang compiles fine for android and you can use this for calling Android JNI libraries as well,
but you will need to do everything inside a android environment. You can create a docker image using the contents
of a x86-64 android vm and it will work fine. I have a working android docker at https://hub.docker.com/repository/docker/racerxdl/android

Besides that, you can write the golang app normally and when building:

```bash
// Set the path correctly to your NDK Bundle
export CC=/opt/google/android/android-sdk-linux/ndk-bundle/toolchains/llvm/prebuilt/linux-x86_64/bin/i686-linux-android21-clang
CGO_ENABLED=1 GOOS=android CGO_CXXFLAGS=-std=c++17 CGO_CFLAGS="-I$NDK/sysroot/usr/include" go build -o androidexec
```

Then you can send `androidexec` file to the docker machine and run it from there.
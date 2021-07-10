package gojni

import "fmt"

// JNIState represents the current JNI State
// this object is called by the native JNI functions to lookup java stuff
type JNIState struct {
	classMap              map[string]*JClassHolder
	classRevMap           map[JClass]string
	allocatedClassObjects map[JObject]*jniObj
	allocatedObjects      map[JObject]*jniObj

	lastClass        JClass
	lastObject       JObject
	exceptionOccured bool
	localFrames      []*localFrame
}

var fakeJniEnv = &JNIState{
	classMap:              make(map[string]*JClassHolder),
	classRevMap:           make(map[JClass]string),
	allocatedObjects:      make(map[JObject]*jniObj),
	allocatedClassObjects: make(map[JObject]*jniObj),
}

// SetExceptionOccured sets if an exception has been ocurred in JNI
func (j *JNIState) SetExceptionOccured() {
	j.exceptionOccured = true
}

// UnsetExceptionOccured unsets if an exception has been ocurred in JNI
func (j *JNIState) UnsetExceptionOccured() {
	j.exceptionOccured = false
}

// AllocObject allocates an object inside JNI Context with the specified type and class.
// The content field is acessible by the handlers that receives this object
// Returns JObject
func (j *JNIState) AllocObject(objType ArgType, class JClass, content interface{}) JObject {
	objectId := JObject(int(j.lastObject) + 1)
	j.lastObject++

	j.allocatedObjects[objectId] = &jniObj{
		id:      objectId,
		objType: objType,
		content: content,
		class:   class,
	}

	return JObject(nonClassObjectOffset) + objectId
}

// AllocClassObject allocates an class object inside JNI for usage in reflection primitives
func (j *JNIState) AllocClassObject(objType ArgType, class JClass, content interface{}) {
	objectId := JObject(int(j.lastObject) + 1)
	j.lastObject++

	j.allocatedClassObjects[JObject(class)] = &jniObj{
		id:      objectId,
		objType: objType,
		content: content,
		class:   j.classMap["java/lang/Class"].id,
	}
}

// GetObjectClass returns the JClass for the specified JObject
func (j *JNIState) GetObjectClass(id JObject) (JClass, error) {
	obj := j.getRawObject(id)
	if obj == nil {
		return InvalidClass, fmt.Errorf("object %d not found", id)
	}
	return obj.class, nil
}

// GetObject returns the golang object data for the specified JObject
func (j *JNIState) GetObject(id JObject) (objType ArgType, content interface{}, err error) {
	obj := j.getRawObject(id)
	if obj == nil {
		return 0, nil, fmt.Errorf("object %d not found", int(id))
	}
	return obj.objType, obj.content, nil
}

// DeallocObject deallocates an object from JNI State
// If the object is pinned as global, it only mark as local deleted
// This does not deallocate Class Objects
func (j *JNIState) DeallocObject(id JObject) {
	if id > nonClassObjectOffset {
		if obj, ok := j.allocatedObjects[id]; ok {
			if !obj.markedAsGlobal || obj.deletedAsGlobal {
				delete(j.allocatedObjects, id)
			} else {
				obj.deletedAsLocal = true
			}
		}
	} // Do not dealloc class objects
}

// AddClass adds a new class to the JNI State returning a class holder
// The returned class holder is used for adding method and field handlers
func (j *JNIState) AddClass(className string) *JClassHolder {
	if _, ok := j.classMap[className]; ok {
		log.Warnf("warning: class %s already exists.\n", className)
		return j.classMap[className]
	}

	currentClass := JClass(int(j.lastClass) + 1)

	j.classMap[className] = newJClassHolder(currentClass, className)
	j.classRevMap[currentClass] = className
	j.lastClass = currentClass

	// Don't allocate reflection data for base Class and Null Class
	if className != "java.lang.Class" && className != "fullyNullClass" {
		j.AllocClassObject(ArgTypeJObject, j.classMap[className].id, className)
	}

	return j.classMap[className]
}

// FindClass returns a JClass ID by it's name
func (j *JNIState) FindClass(className string) (JClass, error) {
	if h, ok := j.classMap[className]; ok {
		return h.id, nil
	}
	return -1, fmt.Errorf("class %s not found", className)
}

// GetStaticMethodID returns the JMethod ID for the specified class static method and args
func (j *JNIState) GetStaticMethodID(classId JClass, methodName string, methodArgs string) (JMethodID, error) {
	className, ok := j.classRevMap[classId]
	if !ok {
		return -1, fmt.Errorf("class with id %d not found", classId)
	}
	return j.classMap[className].GetStaticMethod(methodName, methodArgs)
}

// GetMethodIDArgs returns the arguments required for calling the specified method from the specified class
func (j *JNIState) GetMethodIDArgs(classId JClass, methodId JMethodID) []TypedMethodArgs {
	className, ok := j.classRevMap[classId]
	if !ok {
		return nil
	}
	args, _ := j.classMap[className].GetMethodArgs(methodId)
	return args
}

// GetStaticMethodIDArgs returns the arguments required for calling the specified static method from the specified class
func (j *JNIState) GetStaticMethodIDArgs(classId JClass, methodId JMethodID) []TypedMethodArgs {
	className, ok := j.classRevMap[classId]
	if !ok {
		return nil
	}
	args, _ := j.classMap[className].GetStaticMethodArgs(methodId)
	return args
}

// GetStaticFieldID returns the java field ID for the static method in the specified class
func (j *JNIState) GetStaticFieldID(classId JClass, name, args string) (JFieldID, error) {
	className, ok := j.classRevMap[classId]
	if !ok {
		return 0, fmt.Errorf("class with id %d not found", classId)
	}
	return j.classMap[className].GetStaticField(name, args)
}

// GetMethodID returns the method id from the specified class
func (j *JNIState) GetMethodID(classId JClass, name, args string) (JMethodID, error) {
	className, ok := j.classRevMap[classId]
	if !ok {
		return 0, fmt.Errorf("class with id %d not found", classId)
	}
	return j.classMap[className].GetMethod(name, args)
}

// GetStaticObjectField returns the java object for the specified static field in class
func (j *JNIState) GetStaticObjectField(classId JClass, fieldId JFieldID) (JObject, error) {
	className, ok := j.classRevMap[classId]
	if !ok {
		return 0, fmt.Errorf("class with id %d not found", classId)
	}
	return j.classMap[className].StaticFieldGetter(fieldId)
}

// CallObjectMethod calls the specified object method with the specified arguments
func (j *JNIState) CallObjectMethod(obj JObject, methodId JMethodID, args []TypedMethodArgs) JObject {
	class, err := j.GetObjectClass(obj)
	if err != nil {
		log.Errorf("object %d not found: %s\n", obj, err)
		return -1
	}
	className, ok := j.classRevMap[class]
	if !ok {
		log.Errorf("class %d not found \n", class)
		return -1
	}
	obj, err = j.classMap[className].CallMethod(methodId, args)
	if err != nil {
		log.Errorf(err.Error())
	}
	return obj
}

// CallStaticObjectMethod calls the static method in the specified object with the specified arguments
func (j *JNIState) CallStaticObjectMethod(classId JClass, methodId JMethodID, args []TypedMethodArgs) JObject {
	className := j.classRevMap[classId]
	if className == "" {
		log.Errorf("class %d not found \n", classId)
		return -1
	}
	obj, err := j.classMap[className].CallStaticMethod(methodId, args)
	if err != nil {
		log.Errorf(err.Error())
		return -1
	}
	return obj
}

// NewStringUTF allocates a new java string
func (j *JNIState) NewStringUTF(data string) (JString, error) {
	obj := j.AllocObject(ArgTypeJString, -100, data)
	return JString(obj), nil
}

// GetString returns the string stored as java string
func (j *JNIState) GetString(str JString) (string, error) {
	t, data, err := j.GetObject(JObject(str))
	if err != nil {
		return "", err
	}
	if t != ArgTypeJString {
		return "", fmt.Errorf("object %d is not string", str)
	}

	return data.(string), nil
}

// GetStringUTFLength returns the length of the stored java string
func (j *JNIState) GetStringUTFLength(str JString) (JSize, error) {
	t, data, err := j.GetObject(JObject(str))
	if err != nil {
		return 0, err
	}
	if t != ArgTypeJString {
		return 0, fmt.Errorf("object %d is not string", str)
	}

	return JSize(len(data.(string))), nil
}

// GetStringUTFRegion gets a golang string from the specified range in the java string
func (j *JNIState) GetStringUTFRegion(str JString, start, l int) (string, error) {
	t, data, err := j.GetObject(JObject(str))
	if err != nil {
		return "", err
	}
	if t != ArgTypeJString {
		return "", fmt.Errorf("object %d is not string", str)
	}

	v := data.(string)
	if len(v) > start {
		v = v[start:]
	}
	if len(v) > l {
		v = v[:l]
	}

	return v, nil
}

// Local Frame

// PushLocalFrame creates a new local frame with the specified length in number of objects
func (j *JNIState) PushLocalFrame(length int) {
	j.localFrames = append(j.localFrames, &localFrame{
		length: length,
	})
}

// PopLocalFrame removes the current frame from the frame stack
// It deallocates local objects in the process
func (j *JNIState) PopLocalFrame(result JObject) {
	if len(j.localFrames) == 0 {
		log.Errorf("No local frame to pop!")
		return
	}
	currentFrame := j.localFrames[len(j.localFrames)-1]
	j.localFrames = j.localFrames[:len(j.localFrames)-1]
	if result != 0 {
		for _, obj := range currentFrame.allocatedObjects {
			if obj != result {
				j.DeallocObject(obj)
			}
		}
	}
}

// AddToLocalFrame adds an object to the local frame of allocated objects
func (j *JNIState) AddToLocalFrame(obj JObject) {
	if len(j.localFrames) == 0 {
		log.Errorf("No local frame to add!")
		return
	}
	frame := j.localFrames[len(j.localFrames)-1]
	frame.allocatedObjects = append(frame.allocatedObjects, obj)
}

// DelFromLocalFrame removes the object from the local frame
func (j *JNIState) DelFromLocalFrame(obj JObject) {
	if len(j.localFrames) == 0 {
		log.Errorf("No local frame to add!")
		return
	}
	frame := j.localFrames[len(j.localFrames)-1]
	for i, v := range frame.allocatedObjects {
		if v == obj {
			frame.allocatedObjects = append(frame.allocatedObjects[:i], frame.allocatedObjects[i+1:]...)
			break
		}
	}
}

// Global

// NewGlobalReference marks the current object as global
// This returns a new JObject referencing the same object in JNIState
// The object will not be deallocated until is deallocated both in local frame and global frame
func (j *JNIState) NewGlobalReference(obj JObject) JObject {
	jobj := j.getRawObject(obj)
	if jobj != nil {
		jobj.markedAsGlobal = true
	}
	return JObject(globalObjectOffset) + obj
}

// DelGlobalReference removes the global reference of the object
// This deallocates the object if it has been already deallocated in the local frame
func (j *JNIState) DelGlobalReference(obj JObject) {
	jobj := j.getRawObject(obj)
	if jobj != nil && jobj.markedAsGlobal {
		jobj.deletedAsGlobal = true
		if jobj.deletedAsLocal {
			j.DeallocObject(obj)
		}
	}
}

// Helpers

// GetClassName returns the class name of the specified java class ID
func (j *JNIState) GetClassName(id JClass) string {
	if name, ok := j.classRevMap[id]; ok {
		return name
	}
	return fmt.Sprintf("JClass(%d)", id)
}

// GetClassFieldName returns the field name for a field inside a class
func (j *JNIState) GetClassFieldName(id JClass, fieldId JFieldID) string {
	className, ok := j.classRevMap[id]
	if !ok {
		return fmt.Sprintf("JFieldID(%s, %d)", j.GetClassName(id), fieldId)
	}
	for i, v := range j.classMap[className].fields {
		if v == fieldId {
			return i
		}
	}
	return fmt.Sprintf("JFieldID(%s, %d)", j.GetClassName(id), fieldId)
}

// GetClassStaticFieldName returns the static field name for a static field inside a class
func (j *JNIState) GetClassStaticFieldName(id JClass, fieldId JFieldID) string {
	className, ok := j.classRevMap[id]
	if !ok {
		return ""
	}
	for i, v := range j.classMap[className].staticFields {
		if v == fieldId {
			return i
		}
	}
	return ""
}

// GetJNI returns a JNIState Singleton
func GetJNI() *JNIState {
	return fakeJniEnv
}

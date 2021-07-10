package gojni

type jniObj struct {
	id      JObject
	objType ArgType
	class   JClass
	content interface{}

	// GC Status
	markedAsGlobal  bool
	deletedAsLocal  bool
	deletedAsGlobal bool
}

type localFrame struct {
	length           int
	allocatedObjects []JObject
}

func (j *JNIState) getRawObject(id JObject) *jniObj {
	if id > globalObjectOffset {
		id -= globalObjectOffset
	}
	if id > nonClassObjectOffset {
		return j.allocatedObjects[id-nonClassObjectOffset]
	}

	return j.allocatedClassObjects[id]
}

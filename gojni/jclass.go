package gojni

import "fmt"

// JClassHolder holds the java class definition with handlers for fields, methods and functions
type JClassHolder struct {
	id            JClass
	name          string
	methods       map[string]JMethodID
	staticMethods map[string]JMethodID
	fields        map[string]JFieldID
	staticFields  map[string]JFieldID

	lastMethod       JMethodID
	lastStaticMethod JMethodID
	lastField        JFieldID
	lastStaticField  JFieldID

	methodArgs       map[JMethodID]JNIMethodArgs
	staticMethodArgs map[JMethodID]JNIMethodArgs

	methodHandle            map[JMethodID]JNICallHandle
	staticMethodHandle      map[JMethodID]JNICallHandle
	fieldSetterHandle       map[JFieldID]JNICallHandle
	fieldGetterHandle       map[JFieldID]JNICallHandle
	staticFieldSetterHandle map[JFieldID]JNICallHandle
	staticFieldGetterHandle map[JFieldID]JNICallHandle
}

func newJClassHolder(id JClass, name string) *JClassHolder {
	return &JClassHolder{
		id:                      id,
		name:                    name,
		methods:                 make(map[string]JMethodID),
		staticMethods:           make(map[string]JMethodID),
		fields:                  make(map[string]JFieldID),
		staticFields:            make(map[string]JFieldID),
		methodArgs:              make(map[JMethodID]JNIMethodArgs),
		staticMethodArgs:        make(map[JMethodID]JNIMethodArgs),
		methodHandle:            make(map[JMethodID]JNICallHandle),
		staticMethodHandle:      make(map[JMethodID]JNICallHandle),
		fieldSetterHandle:       make(map[JFieldID]JNICallHandle),
		fieldGetterHandle:       make(map[JFieldID]JNICallHandle),
		staticFieldSetterHandle: make(map[JFieldID]JNICallHandle),
		staticFieldGetterHandle: make(map[JFieldID]JNICallHandle),
	}
}

func (c *JClassHolder) nameArgsToName(name, args string) string {
	return name + "|" + args
}

// AddMethod adds a method handler to the class
func (c *JClassHolder) AddMethod(name, args string, handler JNICallHandle) *JClassHolder {
	name = c.nameArgsToName(name, args)
	currentMethodId := JMethodID(0)
	if _, ok := c.methods[name]; ok {
		fmt.Printf("warning: method %s already exists.\n", name)
		currentMethodId = c.methods[name]
	} else {
		currentMethodId = JMethodID(int(c.lastMethod) + 1)
		c.lastMethod = currentMethodId
	}

	c.methods[name] = currentMethodId
	c.methodHandle[currentMethodId] = handler
	c.methodArgs[currentMethodId] = createTypedMethodArgsFromString(args)

	return c
}

// AddStaticMethod adds a static method handler to the class
func (c *JClassHolder) AddStaticMethod(name, args string, handler JNICallHandle) *JClassHolder {
	name = c.nameArgsToName(name, args)
	currentMethodId := JMethodID(0)
	if _, ok := c.staticMethods[name]; ok {
		fmt.Printf("warning: static method %s already exists.\n", name)
		currentMethodId = c.staticMethods[name]
	} else {
		currentMethodId = JMethodID(int(c.lastStaticMethod) + 1)
		c.lastStaticMethod = currentMethodId
	}

	c.staticMethods[name] = currentMethodId
	c.staticMethodHandle[currentMethodId] = handler
	c.staticMethodArgs[currentMethodId] = createTypedMethodArgsFromString(args)

	return c
}

// AddFieldGetter adds a field getter handler to the class
func (c *JClassHolder) AddFieldGetter(name string, handler JNICallHandle) *JClassHolder {
	currentFieldId := JFieldID(0)
	if _, ok := c.fields[name]; ok {
		currentFieldId = c.fields[name]
	} else {
		currentFieldId = JFieldID(int(c.lastField) + 1)
		c.lastField = currentFieldId
	}

	c.fields[name] = currentFieldId
	c.fieldGetterHandle[currentFieldId] = handler

	return c
}

// AddFieldSetter adds a field setter handler to the class
func (c *JClassHolder) AddFieldSetter(name string, handler JNICallHandle) *JClassHolder {
	currentFieldId := JFieldID(0)
	if _, ok := c.fields[name]; ok {
		currentFieldId = c.fields[name]
	} else {
		currentFieldId = JFieldID(int(c.lastField) + 1)
		c.lastField = currentFieldId
	}

	c.fields[name] = currentFieldId
	c.fieldSetterHandle[currentFieldId] = handler

	return c
}

// AddStaticFieldGetter adds a static field getter handler to the class
func (c *JClassHolder) AddStaticFieldGetter(name string, handler JNICallHandle) *JClassHolder {
	currentFieldId := JFieldID(0)
	if _, ok := c.staticFields[name]; ok {
		currentFieldId = c.staticFields[name]
	} else {
		currentFieldId = JFieldID(int(c.lastStaticField) + 1)
		c.lastStaticField = currentFieldId
	}

	c.staticFields[name] = currentFieldId
	c.staticFieldGetterHandle[currentFieldId] = handler

	return c
}

// AddStaticFieldSetter adds a static field setter handler to the class
func (c *JClassHolder) AddStaticFieldSetter(name string, handler JNICallHandle) *JClassHolder {
	currentFieldId := JFieldID(0)
	if _, ok := c.staticFields[name]; ok {
		currentFieldId = c.staticFields[name]
	} else {
		currentFieldId = JFieldID(int(c.lastStaticField) + 1)
		c.lastField = currentFieldId
	}

	c.staticFields[name] = currentFieldId
	c.staticFieldSetterHandle[currentFieldId] = handler

	return c
}

// GetMethod returns the java method id for the specified method
func (c *JClassHolder) GetMethod(name, args string) (JMethodID, error) {
	name = c.nameArgsToName(name, args)
	methodId, ok := c.methods[name]
	if ok {
		return methodId, nil
	}
	return InvalidMethod, fmt.Errorf("method %s.%s(%s) not found", c.name, name, args)
}

// GetStaticMethod returns the java method id for the specified static method
func (c *JClassHolder) GetStaticMethod(name, args string) (JMethodID, error) {
	name = c.nameArgsToName(name, args)
	methodId, ok := c.staticMethods[name]
	if ok {
		return methodId, nil
	}
	return InvalidMethod, fmt.Errorf("method %s::%s(%s) not found", c.name, name, args)
}

// GetStaticMethodArgs returns args needed for calling the specified static method
func (c *JClassHolder) GetStaticMethodArgs(id JMethodID) (JNIMethodArgs, error) {
	if args, ok := c.staticMethodArgs[id]; ok {
		return args, nil
	}
	return nil, fmt.Errorf("static method with id %d not found in class %s\n", id, c.name)
}

// GetMethodArgs returns args needed for calling the specified method
func (c *JClassHolder) GetMethodArgs(id JMethodID) (JNIMethodArgs, error) {
	if args, ok := c.methodArgs[id]; ok {
		return args, nil
	}
	return nil, fmt.Errorf("method with id %d not found in class %s\n", id, c.name)
}

// CallStaticMethod calls the static method with the specified arguments
func (c *JClassHolder) CallStaticMethod(id JMethodID, args JNIMethodArgs) (JObject, error) {
	if h, ok := c.staticMethodHandle[id]; ok {
		return h(fakeJniEnv, args), nil
	}
	return InvalidObject, fmt.Errorf("cannot find static method with id %d in class %s\n", id, c.name)
}

// CallMethod calls the method with the specified arguments
func (c *JClassHolder) CallMethod(id JMethodID, args JNIMethodArgs) (JObject, error) {
	if h, ok := c.methodHandle[id]; ok {
		return h(fakeJniEnv, args), nil
	}
	return InvalidObject, fmt.Errorf("cannot find method with id %d in class %s\n", id, c.name)
}

// FieldGetter calls the field getter
func (c *JClassHolder) FieldGetter(id JFieldID) (JObject, error) {
	if h, ok := c.fieldGetterHandle[id]; ok {
		return h(fakeJniEnv, nil), nil
	}
	return InvalidObject, fmt.Errorf("cannot find field with id %d in class %s\n", id, c.name)
}

// StaticFieldGetter calls the static field getter
func (c *JClassHolder) StaticFieldGetter(id JFieldID) (JObject, error) {
	if h, ok := c.staticFieldGetterHandle[id]; ok {
		return h(fakeJniEnv, nil), nil
	}
	return InvalidObject, fmt.Errorf("cannot find static field with id %d in class %s\n", id, c.name)
}

// GetStaticField returns the java field ID for the specified static field
func (c *JClassHolder) GetStaticField(name, args string) (JFieldID, error) {
	//name = c.nameArgsToName(name, args)
	fieldId, ok := c.staticFields[name]
	if ok {
		return fieldId, nil
	}
	return InvalidField, fmt.Errorf("field %s::%s(%s) not found", c.name, name, args)
}

// GetField returns the java field ID for the specified field
func (c *JClassHolder) GetField(name, args string) (JFieldID, error) {
	//name = c.nameArgsToName(name, args)
	fieldId, ok := c.fields[name]
	if ok {
		return fieldId, nil
	}
	return InvalidField, fmt.Errorf("field %s.%s(%s) not found", c.name, name, args)
}

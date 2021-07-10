#!/bin/bash

JAVA_INC=$JAVA_HOME/include

javac HelloJNI.java

g++ -std=c++11 -shared -fPIC -I"$JAVA_INC" -I"$JAVA_INC"/linux HelloJNI.cpp -o libhello.so

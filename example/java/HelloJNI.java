public class HelloJNI {
    static {
        System.loadLibrary("hello");
    }

    private native void sayHello(String name);

    public static void main(String[] args) {
        new HelloJNI().sayHello("Dave");
    }
}
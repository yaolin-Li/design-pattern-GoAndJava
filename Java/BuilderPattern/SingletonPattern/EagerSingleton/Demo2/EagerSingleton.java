package EagerSingleton.Demo2;

/**
 * 饿汉式
 * 静态变量创建类的对象
 */
public class EagerSingleton {

    private EagerSingleton() {}

    private static EagerSingleton instance;
    
    static {
        instance = new EagerSingleton();
    } 

    public static EagerSingleton getInstance() {
        return instance;
    }
}
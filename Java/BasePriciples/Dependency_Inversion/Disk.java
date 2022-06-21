
public class Disk implements AbstractDisk{
    
    public void save(String data) {
        System.out.println("使用硬盘存储数据为："+data);
    }

    public String get() {
        System.out.println("使用硬盘读取数据");
        return "数据";
    }
}
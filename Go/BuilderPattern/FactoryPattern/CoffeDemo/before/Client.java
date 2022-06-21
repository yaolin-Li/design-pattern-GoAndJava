package before;
public class Client {
    public static void main(String[] args) {
        CoffeStore cs = new CoffeStore();
        Coffe c = cs.orderCoffee("latte");
        System.out.println(c.getName());
    }
}

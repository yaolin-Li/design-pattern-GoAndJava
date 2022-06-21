package before;
import javax.management.RuntimeErrorException;

public class CoffeStore {
    public Coffe orderCoffee(String type) {
        Coffe c = null;
        if ("american".equals(type)) {
            c = new AmericanCoffe();
        } else if ("latte".equals(type)) {
            c = new LatteCoffe();
        } else {
            throw new RuntimeErrorException("没有该咖啡");
        }
        c.addMike();
        c.addsugar();
        return c;
    }
}


public class CoffeStore {
    public Coffe orderCoffee(String type) {
        Coffe c = null;
        if ("american".equals(type)) {
            c = new AmericanCoffe();
        } else if ("latte".equals(type)) {
            c = new LatteCoffe();
        }
        c.addMike();
        c.addsugar();
        return c;
    }
}

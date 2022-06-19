package JavaVersion;

public class Client {
    public static void main(String[] args) {
        SougouInput input = new SougouInput();

        //DefaultSkin skin = new DefaultSkin();

        HeimaSkin skin = new HeimaSkin(); 

        input.setSkin(skin);

        input.display();
    }
}

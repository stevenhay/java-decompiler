public abstract class Main {

    public static int sum(int x, int y) {
        return (x + y) + 2;
    }

    public static void main(String[] args) {
        int x = sum(12, 23);
        System.out.println(x);
    }
}

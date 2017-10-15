class Booleans{
    public boolean trueOrFalse(){
        return false;
    }
    public static void main(String[] args) {
        Booleans object = new Booleans();
        boolean result = object.trueOrFalse();
        if (result == true || result == false){
            System.out.println("Correct!");
        } else {
            System.out.println("Incorrect!");
        }
    }
}
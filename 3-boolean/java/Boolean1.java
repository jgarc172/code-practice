class Boolean1 {
    public static boolean isPositive(int num) {
        return num > 0;
    }


    
    public static void main(String[] args) {
        int[] ints = {-4, 0, 10};

        if (isPositive(ints[0])){
            System.out.println("incorrect");
        }
        if (isPositive(ints[1])){
            System.out.println("incorrect");
        }
        if (!isPositive(ints[2])){
            System.out.println("incorrect");
        }

        System.out.println("Correct!");
        
    }
}
class Problem1 {
    public static boolean isPositive(int num) {
        return num > 0;
    }

    public static void main(String[] args) {

        Test[] tests = {
            new Test(-4, false), 
            new Test(0, false), 
            new Test(10, true)
            };
        
        runTests(tests);
    }

    public static void runTests(Test[] tests){
        System.out.println();
        for (Test test : tests) {
            boolean result = isPositive(test.input);
            String output = String.format("isPositive(%d) \t ->  %b ", test.input, result);

            if (result == test.expected){
                output += "\t (OK)";
            } else {
                output += String.format("\t(X) expected %b %n", test.expected);
            }

            System.out.println(output);
        }
    }

        static class Test {
            int input;
            boolean expected;

            Test(int input, boolean expected) {
                this.input = input;
                this.expected = expected;
            }
        }

}
## Conversion from integer to boolean
One can create a function that converts the set of integer values to the set of boolean values.  A 
generic description of such a function is of this form:

    integer -> boolean

But how does one decide which set of integers is true and which is false?  It depends on the 
problem.  For example, we can create a function, isPositive, that maps those integer values to
their corresponding boolean value.

## Problem/Solution

create a function, isPositive, that takes an integer and returns a boolean
to indicate if the integer is positive or not

    isPositive: integer -> boolean

For example:

    isPositive(-4) 	 ->  false 	 (OK)
    isPositive(0) 	 ->  false 	 (OK)
    isPositive(10) 	 ->  true 	 (OK)

# Solution
By definition, an integer is positive if its value is greater than zero.  Therefore the set of integer 
values can be split into 
    values that are greater than zero -> true
    and 
    values that are not -> false

# Java
complete the function named isPositive that takes an integer num and returns true
if num is positive; otherwise it returns false:


```Java
class Boolean1 {

    public static boolean isPositive(int num) {

        // implement this function
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
```

# go
Create a function named invoke that can be invoked from within 
the main method:

    invoke()

```go
package main

// function invoke




func main(){
    invoke()
}
```

# Python
Create a function named invoke that can be invoked from within 
the main code:

    invoke()

```Python
# function invoke
   

invoke()


```

# JavaScript
Create a function named invoke that can be invoked from within 
the main code:

    invoke()

```JavaScript

// function invoke



invoke();
```
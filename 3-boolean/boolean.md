## Problem/Solution

create a function, isPositive, that takes an integer and returns a boolean
to indicate if the integer is positive or negative

    isPositive: integer -> boolean

# Solution
There are two boolean values: positive and negative.
The set of integer values can be split into two subsets: those values that are positive and those
that are not.  An integer is positive if its value is greater than zero.

# Java
complete the function named isPositive that takes an integer num and returns true
if num is positive:


```Java
class Boolean1 {

    public static boolean isPositive(int num) {


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
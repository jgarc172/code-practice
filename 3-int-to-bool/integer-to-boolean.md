# Conversion from integer to boolean
One can create a function that converts the set of integer values to the set of boolean values.  A 
generic description of such a function is of this form:

    integer -> boolean

But how does one decide which set of integers is true and which is false?  It depends on the 
problem.  For example, we can create a function, isPositive, that maps those integer values to
their corresponding boolean value.

## Problem 1, isPositive
create a function, isPositive, that takes an integer and returns a boolean
to indicate if the integer is positive or not

    isPositive: integer -> boolean

For example:

    isPositive(-4) 	 ->  false 	 (OK)
    isPositive(0) 	 ->  false 	 (OK)
    isPositive(10) 	 ->  true 	 (OK)

## Solution 1, isPositive
By definition, an integer is positive if its value is greater than zero.  Therefore the set of integer values can be split into 

    values that are greater than zero -> true
    and 
    values that are not -> false

## Java Implementation
The following implementation returns true
if num is greater than zero; otherwise it returns false:

```Java
public static boolean isPositive(int num) {

    return num > 0;
}
```

# Expressions and Value Composition 

Now that we've seen how literals, operators, and control structures produce a value of a given type, we can combine them as needed to produce values of a desired type.  This approach of composing values is basically creating expressions but can be tought of as **value composition**.  The basic level practice exercises will aim to illustrate the use of common expressions as value composition.  These in turn will become useful when dealing with more advanced exercises.

## Examples of Expressions
### Arithmetic Expressions

Let's combine the application of the `+` and the `*` operators in the following way:

```
(3 + 5) * 10 
   8    * 10
      80
```

So, the result of the `+` expression is an integer and is also the input to the `*` expression to eventually produce an integer.

```
( (+) -> integer ) -> (*) -> integer
```

### Comparison and If Control Expressions

Now, let's express the following problem using the `if` control construct as part of the solution:

#### Problem:
> Determine if the water is frozen.  The temperatue is 25 degrees Fahrenheit.  Water becomes ice (frozen) when the temperature is 32 degrees Fahrenheit or less.

#### Solution:
- **Result**: The final value should be whether or not the water is frozen.  This is naturally a boolean value: true if the water is frozen; false otherwise.  We can use the boolean variable `isFrozen` to assign the final boolean value.
- **Given**: The current temperature.  Assign the integer value `25` to a variable `temperature`.  But we need a boolean.  See below on how the `temperature` integer is converted to a boolean using a comparison operator.
- **Given**: Definition of ice (frozen water).  Water becomes ice when the temperature is 32 degrees F. or less.  
- **Given**: We know that the comparison operators produce a boolean, and the definition is a comparison of two values:  `temperature <= 32`. 
- **solution 1**: just assign the result of the comparison operation to the variable `isFrozen`
- **solution 2**: use the `if` control structure to explicitly assign the boolean value to `isFrozen`, using the `<=` operation to obtain the boolean condition for the `if` control.

### Implementation 1 in Java (no `if/else` control)
```java
boolean isFrozen; // final value to be determined. In Java it defaults to false.
int temperature = 25;
isFrozen = temperature <= 32;
. . .
```

The application of the `<=` operator produces a boolean value that determines the value to be assigned to `isFrozen`: 

```java
   (<=) -> boolean
   
temperature <= 32  
    25      <= 32
       true

isFrozen = true;
```

### Implementation 2 in Java (`if/else` control)
```java
boolean isFrozen; // final value to be determined. In Java it defaults to false.
int temperature = 25;

if (temperature <= 32){
   isFrozen = true; 
} else {
   isFrozen = false;
}
. . .
```

The application of the `<=` operator produces the boolean value needed by the `if/else` control structure.  Since the boolean expression is `true` the `if` block statement is executed.  The `else` block statement is not executed.

```java
   ( (<=) -> boolean ) -> (if/else) -> execute one of two blocks of statements

if (temperature <= 32){
   isFrozen = true; 
} else {
   isFrozen = false;
}
. . .
if (true){
   isFrozen = true; 
} else {
   isFrozen = false;
}
. . .
   isFrozen = true; 
```

### Comparing Implementation 1 and Implementation 2

The two implementations are equivalent, as they both produce the boolean value needed to determine the value of `isFrozen`.

```java
Implementation 1:
   
   (<=) -> boolean
   
   isFrozen = temperature <= 32;
   
Implentation 2:

   ( (<=) -> boolean ) -> (if/else) -> execute one of two blocks of statements
   
   if (temperature <= 32){
      isFrozen = true; 
   } else {
      isFrozen = false;
   }
```

In this case, the `if/else` construct was **not** necessary as the final value is a boolean.  For other cases where the final result is not a boolean, the `if` or the `if/else` construct is very important.

### Expressions in Functions

If you find that an expression can be reused, you can define it as a function (method).  For example, the solution in the previous exercise was captured as a boolean variable `isFrozen`.  We can also capture the above expression as a named function and apply it to more than just the temperature `25`.

#### Problem
> Implement a function named `isFrozen` to determine if water is frozen for a given temperature, knowing that water becomes ice (frozen) when the temperature is 32 degrees Fahrenheit or less.

Implement the function using Implementation 1 and then using Implementation 2.

```java
boolean isFrozen(int temperature){

}
```
<details>
<summary>Solution Using Implementation 1:</summary>
<pre><code>

boolean isFrozen(int temperature){
   return temperature <= 32;
}

</code><pre>
</details>

<details>
<summary>Solution Using Implementation 2:</summary>
<pre><code>

boolean isFrozen(int temperature){
   if (temperature <= 32){
      return true;
   } else {
      return false;
   }
}

</code></pre>
</details>

## Review
Solutions to a problem is a combination of applying smaller expressions.  It is preferrable to use small expression when the clarity of the solution is obvious.  If it is more clear to use a longer expressions, even if it takes extra steps, then use it. 
 

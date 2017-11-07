
# Expressions and Value Composition as Problem Solving

Now that we've seen how literals, operators, and control structures produce a value of a given type, we can combine them as needed to produce values of a desired type.  This approach of composing values is basically creating expressions but can be tought of as **value composition**.  The basic level practice exercises will aim to illustrate the use of common expressions as value composition.  These in turn will become useful when dealing with more advanced exercises.

## Examples of Expressions

Let's combine the application of the `+` and the `*` operators in the following way:

```
(3 + 5) * 10 
   8    * 10
      80
```

So, the result of the `+` expression is an input to the `*` expression.

Now, let's express the following problem using the `if` control construct:

> Determine if the water is frozen.  The temperatue is 25 degrees Fahrenheit.  Water becomes ice when the temperature is 32 degrees Fahrenheit or less.

```java
boolean isFrozen = false;
int temperature = 25;

if (temperature <= 32){
   isFrozen = true; 
}
. . .
```

The application of the `<=` operator produces a boolean value that determines if the next block of code is exectuted.

```java
temperature <= 32  -> true
```

Since the boolean expression is `true` the next statement is executed:

```java
isFrozen = true;
```
 
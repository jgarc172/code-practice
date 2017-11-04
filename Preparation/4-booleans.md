# Operators that Create Boolean Values

Programming languages have control structures (if/else and loops) that control the flow of execution depending on a given 
condition of `true` or `false`.  That is, these control structures depend on a boolean value to execute the next statement.

One way to create boolean values is to just use the corresponding literals: `true` and `false`.  But this is not dynamic.  
Programming languages use comparison operators to create boolean values from other data types.

## Comparison Operators

Numeric data types like integers have a set of comparison operators that produce boolean values:

```
== (equal)
!= (not equal)
>  (greater than)
<  (less than)
>= (greater than or equal)
<= (leass than or equal)
```

All these operators can be thought of being functions that take two integers and return a boolean.  For example, the 
`==` operator can be described in the following function notation:

```
equal : integer, integer -> boolean
```

In Java, this function can be implemented as follows:

```java
boolean equal(int num1, int num2){
  return num1 == num2;
}
```

## Logical Operators

boolean data types have logical operators that produce boolean values:

```
&& (and)
|| (or)
!  (not)
```
As we saw in an earlier section, the `not` logical operator is a unary operator that takes one boolean and returns a boolean.
The `and` and `or` logical operators can be thought of functions that take two boolean values and return a boolean.  For 
example the `or` operator can be described in the following function notation:

```
or : boolean, boolean -> boolean
```

In Java, this function can be implemented as follows:

```java
boolean or(boolean val1, boolean val2){
  return val1 && val2;
}
```

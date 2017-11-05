
# Value Composition as Problem Solving

Now that we've seen how literals, operators, and control structures produce a value of a given type, we can combine them as 
needed to produce values of a desired type.  This approach of composing values can be tought of **value** composition as 
opposed to **function** composition.  The basic level practice exercises will make use of value composition.  Eventually, for 
advanced practice exercises we can make use of function composition to solve problems.

## Logical NOT Function as an Example of Composition
This is a trivial example, but it is used to demonstrate the key approaches to solving problems.

Here is the definition and function signature:

```
The not function takes a boolean value and it returns
the reverse of its value:

  not : boolean -> boolean
```

A simple solution would be to ask the question: "is there an existing function or operator that has the same function signature 
and the same definition?  One operator that comes to mind is the `!` operator in Java, Go, and JavaScript.  In Python there is
the `Not` operator.  We'll use the Go operator for the demonstration.

```go
! reverses the logical value of its operand

    ! : boolean -> boolean

!true  -> false
!false -> true
```

Then we can implement it using the Go logical operator:

```go
func not(value bool) bool {
    return !value
}
```

We can claim that the solution is made up of the composition of applying the operator `!` to the given parameter `value`.  It 
aids to know that the operator `!` is itself a function that takes a boolean value and returns a boolean value as:

```
! : boolean -> boolen
```

That is, the function signature `not(value bool)` is equivalent to its implementation `!value`.  At runtime, when we use 
actual values, we can claim that

```
not(false) -> !false -> true
```

Therefore we have shown that

> applying the function **not** is equivalent to applying the operator **!**

## Steps to Solving Problems with Composition

* Use the input parameters as they should always be part of the solution
* Consider the function signature and consider any built-in operators or functions that have the same function signature
* If the built-in operator or function doesn't have the exact needed signature, **convert** the operator using another 
function
* Use composition to achieve the solution

In the exercises, we will try to provide hints that will allow you to apply the above steps.

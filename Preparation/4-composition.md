
# Composition as Problem Solving

To solve any problem, small or large, it is important that we understand the smallest components of a problem and its 
respective solution.

Programming languages already give us those small and fundamental (**primitive**) elements.  We already introduced the *three
basic value types* and some of the *built-in* operators.  In addition to those, we have to include the use of **literals** to 
produce such element values.

## Logical NOT Function as an Example of Composition
The function `not` that was implemented in a previous section is an example of composition.  Here is the definition and 
function signature:

```
The not function takes a boolean value and it returns
the reverse of its value:

  not : boolean -> boolean
```
And here it is its implementation in Go:

```go
func not(value bool) bool {
    return !value
}
```

We can claim that the solution is made up of the composition of the operator `!` and the given parameter `value`.  We can also
claim that the operator `!` is itself a function that takes a boolean value and returns a boolean value as:

```
! : boolean -> boolen
```

That is, the function definition `not(value bool)` is equal to `!value`.  At runtime, when we use actual values, we can claim 
that

```
not(false) -> !false -> true
```

Therefore we can claim that 

```
not is equal to !
```

## Steps to Solving Problems with Composition

* Use the input parameters as they should always be part of the solution
* Consider the function signature and consider any built-in operators or functions that have the same function signature
* If the built-in operator or function doesn't have the exact needed signature, **convert** the operator using another function
* Use composition to achieve the solution

In the exercises, we will try to provide hints that will allow you to apply the above steps.

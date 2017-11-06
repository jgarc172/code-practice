# Function Notation

## Function (Method) Notation
Exercises will be based on implementing a given function (method) signature with its corresponding definition.

The following function signature notation will be used for all the exercises and will be translated to their respective programming language:

```
name of function or method : type1, type2, ... -> returned type
```

To demonstrate the use of the notation, we can define and implement a function that takes two integer values and returns an integer value.

## plus function

The plus function is a simple example of using a function instead of the operator `+` to sum two integer values.

Here is the definition and function signature:

```
The plus function takes two integer values and it returns their sum as an integer:

  plus : integer, integer -> integer
```

Example invocation (application) of `and` and its result:
```
plus(2, 3)  -> 5
plus(5, -5) -> 0
```

### Implementation in Java

```java
int plus(int num1, int num2){
    return num1 + num2;
}
```

which has a signature:

```java
int plus(int num1, int num2)
```

and an implementation:

```java
return num1 + num2;
```

invocation:
```
object.plus(2, 3)  -> 5
object.plus(5, -5) -> 0
```

### Implementation in Go 

```go
func plus(num1 int, num2 int) int {
    return num1 + num2
}
```

which has a signature:

```go
plus(num1 int, num2 int) int
```

and an implementation:

```go
return num1 + num2
```

invocation:
```go
plus(2, 3)  -> 5
plus(5, -5) -> 0
```

### Implementation in Python 

In Python there is already a logical operator `not`, so for the emplementation we will use `Not` as the name of the function.

```python
def Not(value):
    # boolean -> boolean
    return not value
```

which has a signature (and a comment):

```python
Not(value):
  # boolean -> boolean
```

and an implementation:

```python
return not value
```

invocation:
```
Not(True)  -> False
Not(False) -> True
```

### Implementation in JavaScript 

```javascript
function not(value){
    // boolean -> boolean
    return !value;
}
```

which has a signature (and a comment):

```javascript
not(value)
  // boolean -> boolean
```

and an implementation:

```javascript
return !value;
```

invocation:
```
not(true)  -> false
not(false) -> true
```

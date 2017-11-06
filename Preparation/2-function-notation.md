# Function Notation

## Function (Method) Notation
Exercises will be based on implementing a given function (method) signature with its corresponding definition.

The following function signature notation will be used for all the exercises and will be translated to their respective 
programming language:

```
name of function or method : type1, type2, ... -> returned type
```

To demonstrate the use of the notation, we can define and implement a function that takes one boolean value and 
returns a boolean value.

## Logical NOT Operator as a Conversion Function 
The logical NOT operator is a conversion operator; that is, it takes a boolean value and it converts the value to its logical reverse.

Here is the definition and function signature:

```
The not function takes a boolean value and it returns the reverse of its value:

  not : boolean -> boolean
```

Example invocation (application) of `not` and its result:
```
not(true)  -> false
not(false) -> true
```

### Implementation in Java

```java
boolean not(boolean value){
    return !value;
}
```

which has a signature:

```java
boolean not(boolean value)
```

and an implementation:

```java
return !value;
```

invocation:
```
object.not(true)  -> false
object.not(false) -> true
```

### Implementation in Go 

```go
func not(value bool) bool {
    return !value
}
```

which has a signature:

```go
not(value bool) bool
```

and an implementation:

```go
return !value
```

invocation:
```
not(true)  -> false
not(false) -> true
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

# Basic Types and Function Notation

## Basic Value Types 
The coding exercises will be based on the three most common fundamental value types in the programming languages used in the 
exercises: 

* boolean
* integer
* string

## Function (Method) Notation
The following function signature notation will be used for all the exercises and will be translated to their respective 
programming language:

```
name of function or method : type1, type2, ... -> returned type
```

To demonstrate the use of the notation, we can define and implement a function that takes one boolean boolean parameter and 
returns a boolean value.

## Logical NOT Operator as a Conversion Function 
The logical NOT operator is a conversion operator; that is, it takes a boolean value and it converts the value to its logical 
reverse.

Here is the definition and function signature:

```
The not function takes a boolean value and it returns
the reverse of its value:

  not : boolean -> boolean
```

Example invocation of `not` and its result:
```
not(true)  -> false
not(false) -> true
```

### Implementation in Java

```java
boolean not(boolean val){
    return !val;
}

invocation:
object.not(true)  -> false
object.not(false) -> true
```

### Implementation in Go 

```go
func not(val bool) bool {
    return !val
}

invocation:
not(true)  -> false
not(false) -> true
```

### Implementation in Python 

In Python there is already a logical operator `not`, so for the emplementation we will use `Not` as the name of the function.

```python
def Not(val):
    # boolean -> boolean
    return not val
    
Not(True)  -> False
Not(False) -> True
```

### Implementation in JavaScript 

```javascript
function not(val){
    // boolean -> boolean
    return !val;
}

invocation:
not(true)  -> false
not(false) -> true
```

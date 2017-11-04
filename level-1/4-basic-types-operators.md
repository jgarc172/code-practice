# Basic Types and Built-in Operators as Functions

## Basic Value Types 
The coding exercises will be based on the three most common fundamental value types: 

* boolean
* integer
* string

## Logical NOT Operator as a Conversion Function 
The logical NOT operator is a conversion operator; that is, it takes a boolean value and it converts the value to its logical reverse.

Let's define and implement the NOT operator as a function:

> The not function takes a boolean value and it returns
the reverse of its value:

>>not: boolean -> boolean

> Example:
>> not(true)  -> false<br>
>> not(false) -> true

### Implementation in Java

```java
boolean not(boolean val){
    return !val;
}
```

### Implementation in Go 

```go
func not(val bool) bool {
    return !val
}
```

### Implementation in Python 

```python
def not(val):
    # boolean -> boolean
    return not val
```

### Implementation in JavaScript 

```javascript
function not(val){
    // boolean -> boolean
    return !val;
}
```
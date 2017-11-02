# Problem

Demonstrate the use of a simple function (or a method).  

## Solution
The minimalist function or method that can be invoked is one defined as taking no input parameters and returning nothing.  The name
of the function is arbitrary in this case.

## Java
In Java, we can write the following `invoke` method and be able to call it:

```java runnable
class Main {

    void invoke() {        
    }
    
    public static void main(String[] args) {
        (new Main()).invoke();
        
    }
}
```


## Go
In Go, we can write the following `invoke` function and be able to call it:

```go runnable
package main

func invoke(){
}

func main(){
    invoke()
}
```

# Python
In Python, we can write the following `invoke` function and be able to call it:

```python runnable
def invoke():
   return

invoke()
```

# JavaScript
In JavaScript, we can write the following `invoke` function and be able to call it:

```javascript runnable
function invoke(){
}

invoke();
```

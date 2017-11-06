# Executing a Minimal program

Execute is the minimum program that can be executed by the respective programming language runtime.  The program should 
do "nothing".

## Java
In Java, we can write the following code that does nothing and be able to execute it:

```java runnable
class Main {
    public static void main(String[] args) {
        
    }
}
```
The above is typically executed with the following **Java** commands when the code is saved in the 'Execute.java` file:

```bash
$ javac Execute.java
$ java Execute
$ 
```
There is no output as the program does nothing.

## Go
In Go, we can write the following code and be able to execute it:

```go runnable
package main

func main(){
    
}
```
The above is typically executed with the following **Go** commands when the code is saved in the 'execute.go` file:

```bash
$ go run execute.go
$ 
```

# Python
In Python, we can write the following code and be able to execute it.  Since the file is empty of Python code, it will 
execute nothing:

```python runnable
# this is a comment
```
The above is typically executed with the following **Python** commands when the code is saved in the 'execute.py` file:

```bash
$ python execute.py
$
```

# JavaScript
In JavaScript, we can write the following code and be able to execute it.  Since the file is empty of JavaScripts code, it will 
execute nothing:

```javascript runnable
// this is a comment
```
The above is typically executed in a browser or with the following **node** commands when the code is saved in the 'execute.js` file:

```bash
$ node execute.js
$
```

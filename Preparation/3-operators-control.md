
# Function Signature of Built-in Operators and Control Structures

To take advantage and to make proper use of built-in operators and control structures, it helps to understand their inherent function signature.  By doing so, we can see that the applying of these operators and control structures requires the input of values of a given type, and that some of them will produce a value of a given type.

## Built-in Operators
Here is a list of some of the common built-in operators in the programming languages used in the exercises:

```
boolean operators and their function signature: 

!  : boolean -> boolean
example:    !true    -> false

&& : boolean, boolean -> boolean
example:    true && true -> true

|| : boolean, boolean -> boolean
example:    false || true -> true

integer operators and their function signature:

+ : integer, integer -> integer
example:    10 + 5 -> 15

- : integer, integer -> integer
example:    10 - 5 -> 5

* : integer, integer -> integer
example:    10 * 5 -> 50

string: 
(ToDo)

integer comparison operators and their function signature:

== : integer, integer -> boolean
example:    10 == 9 -> false

>  : integer, integer -> boolean
example:    11 > 13 -> false

<  : integer, integer -> boolean
>= : integer, integer -> boolean
<= : integer, integer -> boolean
```

## Control Structures
Here is a list of some of the common control structures.  An attempt to define their function signature is provided.  For example, the `if` control structure can be viewed as a function that takes a boolean value and exectues the next statements if the value is true:

```
if     : boolean -> execute the next block of statements if true
for    : boolean -> repeat execution of next block of statements if true
while  : boolean -> repeat execution of next block of statements if true
switch :  (ToDo)
```

Variables:
---------------------------------------------------------------
-> Need to declare the variables before we can use it for compile-time safety..!
-> Golang(GO) is statically typed (the type of variable is known in prior i.e., at compile time itself)
-> Golang is strongly typed (can't automatically convert types)

variable declaration syntax:
    var <name_of_the_variable> <type_of_the_variable>
    ```ex: var age int```

Path Separator:
----------------------------------------------------------------
-> path package provides utility functions for working with url path strings

import "path"
path.Split(string path) -> returns dir (string), file_name (string)
```ex: path.Split(Desktop/GO-LEARNING/Basics/README.md)```
```
output: dir = Desktop/GO-LEARNING/Basics/
        file_name = README.md
```

Short Declaration
----------------------------------------------------------------
:= is the short declaration 
It's used basically for declaring and assigning the value
-> It can be used only inside the functions
-> Cannot reuse for already declared variables
-> Type is fixed after inference (basically go looks at the value and assigns the type)

Short Declaration vs Normal Declaration
-----------------------------------------
-> Short declaration is the most widely used declaration type and it's suggested to use when the initial values are known and whereas normal declaration when the initial value is unknown!!
-> Short declaration can be used to keep the code concise..
-> Normal declaration can be used when the other developers want to know that some variables are related..! like for instance:
    
```
    var (
        height int
        age int
    
        name string
    )```

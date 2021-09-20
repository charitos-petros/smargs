# smargs ~ Smart Arguments

Smargs is a GO module that helps you add command line arguments in your programs.
For example, you can execute your programs like that:
```bash
./example -flag1 newValue1 -flag2 newValue2
```
```bash
./example ^flag2 newValue2 ^flag1 newValue1
```
```bash
./example -flag2 newValue2
```

## Installation
```bash
go get github.com/charitos-petros/smargs
```

## Functions and Methods
- Controller(): Creates a controller object with no arguments and - as prefix.
- Arg(flag string, variable *string, defaultValue string): Adds an argument to the controller.
- Execute(): Parses the values to the variables.
- Prefix(p string): [Optional] Sets p as prefix.

## Examples
Example that use the default prefix (-):
```go
package main

import (
	"fmt"
	"github.com/charitos-petros/smargs"
)

func main() {
	var var1 string
	var var2 string

	c := smargs.Controller()

	c.Arg("flag1", &var1, "Default1")
	c.Arg("flag2", &var2, "Default2")

	err := c.Execute()
	if err != nil {panic(err)}

	fmt.Println(var1)
	fmt.Println(var2)
}
```
You can assign values to both variables:
```bash
./example -flag1 newValue1 -flag2 newValue2
```
Or, a variable can take the default value if its flag is omitted:
```bash
./example -flag2 newValue2
```

Example with different prefix (^ in this example):
```go
package main

import (
	"fmt"
	"github.com/charitos-petros/smargs"
)

func main() {
	var var1 string
	var var2 string

	c := smargs.Controller()

	c.Arg("flag1", &var1, "Default1")
	c.Arg("flag2", &var2, "Default2")
    
    c.Prefix("^")
	err := c.Execute()
	if err != nil {panic(err)}

	fmt.Println(var1)
	fmt.Println(var2)
}
```
Now, you can execute the program like that:
```bash
./example ^flag2 newValue2 ^flag1 newValue1
```
Or, use different order of flags, or less number of them like before.

## Errors
Execute() function can return errors:
- odd number of arguments: Odd number of arguments given
    For example:
    ```bash
    ./example -flag1
    ```
    ```bash
    ./example -flag2 -flag1 newValue2
    ```
- unknown argument: Unknown flag of argument given
    For example if you execute one the example scripts like that:
    ```bash
    ./example -fLaG newValue2 -flag1 newValue1
    ```
- nil variable given: Second parameter of Arg() function is nil
    For example:
    ```go
    c.Arg("flag1", nil, "Default1")
    ```

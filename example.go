package main

import (
	"fmt"
	"smargs"
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
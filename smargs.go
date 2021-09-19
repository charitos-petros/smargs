// ================= //
//       Smargs      //
//  Smart Arguments  //
//       Author:     //
//  Petros Charitos  //
// ================= //

package smargs

import (
	"os"
	"strings"
	"errors"
)

type controller struct {
	prefix string
	args map[string]data
}

type data struct {
	variable_ *string
	default_ string
	edited bool
}

// Function that creates a Controller object
func Controller() *controller {
	c := controller{"-", make(map[string]data)}
	return &c
}

// Function that adds arguments to Controller
func (c *controller) Arg(flag string, variable_ *string, default_ string) {
	c.args[flag] = data{variable_, default_, false}
}

// Function that changes the prefix of flags
func (c *controller) Prefix(prefix string) {
	c.prefix = prefix
}

// Function that assignes values to variables
func (c *controller) Execute() error{
	l := len(os.Args)
	if l % 2 == 0 {return errors.New("odd number of arguments")}

	// Assign new values to the called arguments
	for i := 1; i < len(os.Args)-1; i = i+2 {
		if strings.HasPrefix(os.Args[i], c.prefix) {
			flag := strings.TrimPrefix(os.Args[i], c.prefix)

			_, ok := c.args[flag]
			if !ok {return errors.New("unknown argument")} // Check if flag is valid

			if c.args[flag].variable_ == nil {return errors.New("nil variable given")}
			*c.args[flag].variable_ = os.Args[i+1]

			c.args[flag] = c.args[flag].get_edited() // Edited status <- True
		}
	}

	// Default values for the other arguments
	for _, d := range c.args {
		if d.edited == false {
			if d.variable_ == nil {return errors.New("nil variable given")}
			*d.variable_ = d.default_
		}
	}

	return nil
}

// Function that changes the Edited status to True
func (d data) get_edited() data{
	return data{d.variable_, d.default_, true}
}

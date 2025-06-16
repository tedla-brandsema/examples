package main

import (
	"fmt"
	"github.com/tedla-brandsema/tagex"
)

// RangeDirective implements the "tagex.Directive[T any]" interface by defining
// both the "Name() string", "Mode() tagex.DirectiveMode" and "Handle(val T) (T, error)" methods.
//
// It also marks two fields (Min and Max) as parameters.
type RangeDirective struct {
	Min int `param:"min"`
	Max int `param:"max"`
}

// Name returns the name of the directive to be used as the directive identifier.
func (d *RangeDirective) Name() string {
	return "range"
}

// Mode returns either `tagex.EvalMode` or `tagex.MutMode`, which indicates whether the directive
// only evaluates the field value or mutates its contents.
func (d *RangeDirective) Mode() tagex.DirectiveMode {
	return tagex.EvalMode
}

// Handle is where the actual work of the directive is performed. Depending on the `tagex.DirectiveMode`
// returned by the Mode() method, it either sets the return value as the field value (i.e., tagex.MutMode)
// or ignores the return value (i.e., tagex.EvalMode).
//
// Even though tagex.Directive[T any] is generic, your implementation of it can be explicit.
// Here Handle takes a val of type "int", therefore "RangeDirective" is of type "int".
// This means we can only apply our RangeDirective to fields of type "int".
func (d *RangeDirective) Handle(val int) (int, error) {
	if val < d.Min || val > d.Max {
		return val, fmt.Errorf("value %d out of range [%d, %d]", val, d.Min, d.Max)
	}
	return val, nil
}

func main() {
	// Create our "check" tag
	checkTag := tagex.NewTag("check")

	// Register our "range" directive with our check tag
	tagex.RegisterDirective(&checkTag, &RangeDirective{})

	// Now we can use our "range" directive on "int" fields of our Car struct
	type Car struct {
		Name   string
		Doors  int `check:"range, min=2, max=4"`
		Wheels int `check:"range, min=3, max=4"`
	}

	// Create an array of "Car" instances
	cars := [...]Car{
		{
			Name:   "Citroën Deux Chevaux",
			Doors:  4,
			Wheels: 4,
		},
		{
			Name:   "Reliant Robin",
			Doors:  3,
			Wheels: 3,
		},
		{
			Name:   "VW Golf",
			Doors:  5,
			Wheels: 4,
		},
	}

	// Invoke the range directive on each car by calling "ProcessStruct" on "checkTag"
	for _, car := range cars {
		if ok, err := checkTag.ProcessStruct(&car); !ok {
			fmt.Printf("The %s did not pass our checks: %v\n", car.Name, err)
			continue
		}
		fmt.Printf("The %s passed our checks!\n", car.Name)
	}
}

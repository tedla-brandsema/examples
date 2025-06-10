package main

import (
	"fmt"
	"github.com/tedla-brandsema/tagex"
)

// RangeDirective implements the "tagex.Directive[T any]" interface by defining
// both the "Name() string" and "Handle(val T) error" methods.
//
// It also marks two fields (Min and Max) as parameters.
type RangeDirective struct {
	Min int `param:"min"`
	Max int `param:"max"`
}

func (d *RangeDirective) Name() string {
	return "range"
}

func (d *RangeDirective) Mode() tagex.DirectiveMode {
	return tagex.EvalMode
}

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
			Name:   "Deux Chevaux",
			Doors:  4,
			Wheels: 4,
		},
		{
			Name:   "Reliant Robin",
			Doors:  3,
			Wheels: 3,
		},
		{
			Name:   "Eliica",
			Doors:  4,
			Wheels: 8,
		},
	}

	// Invoke the range directive on each car by calling "ProcessStruct" on "checkTag"
	for _, car := range cars {
		if ok, err := checkTag.ProcessStruct(car); !ok {
			fmt.Printf("The %s did not pass our checks: %v\n", car.Name, err)
			continue
		}
		fmt.Printf("The %s passed our checks!\n", car.Name)
	}
}

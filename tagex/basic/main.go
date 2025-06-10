package main

import (
	"fmt"
	"github.com/tedla-brandsema/tagex"
	"net/mail"
)

// EmailDirective implements the tagex.Directive[T any] interface by defining
// both the "Name() string" and "Handle(val T) error" methods.
type EmailDirective struct{}

func (d *EmailDirective) Name() string {
	return "email"
}

func (d *EmailDirective) Mode() tagex.DirectiveMode {
	return tagex.EvalMode
}

// Even though tagex.Directive[T any] is generic, your implementation of it can be explicit.
// Here Handle() explicitly is of type "int", which makes our "EmailDirective" explicitly of type "string".
// This means we can use our EmailDirective only on fields of type "string".
func (v *EmailDirective) Handle(val string) (string, error) {
	_, err := mail.ParseAddress(val)
	return val, err
}

func main() {
	// Create our "check" tag
	checkTag := tagex.NewTag("check")

	// Register our "email" directive with our check tag
	tagex.RegisterDirective(&checkTag, &EmailDirective{})

	// Now we can use our "email" directive on the Email field of our "Contact" struct, which works because the field is of type string
	type Contact struct {
		Firstname string
		Lastname  string
		Email     string `check:"email"`
	}

	// Create instances of our "Contact" struct
	contacts := []Contact{
		{
			Email: "john.doe@example.com",
		},
		{
			Email: "user@.com",
		},
		{
			Email: "user123@gmail.com",
		},
		{
			Email: "user@@example.com",
		},
		{
			Email: "info@company.co.uk",
		},
		{
			Email: "user.@example.com",
		},
		{
			Email: "support@web-services.org",
		},
	}

	// Check our contacts by calling "ProcessStruct" on our tag
	for _, contact := range contacts {
		if ok, err := checkTag.ProcessStruct(contact); !ok {
			fmt.Printf("Invalid email %q: %v\n", contact.Email, err)
			continue
		}
		fmt.Printf("Valid email %q\n", contact.Email)
	}
}

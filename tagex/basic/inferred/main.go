package main

import (
	"fmt"
	"github.com/tedla-brandsema/tagex"
	"net/mail"
)

// START inferred-directive

// EmailDirective implements the tagex.Directive interface.
type EmailDirective struct{}

// Name returns the directive name. This string is used in when applying the directive to a field via a tag.
func (d *EmailDirective) Name() string {
	return "email"
}

// Mode tells the tagex library whether this directive simply evaluates the field value or mutates it.
func (d *EmailDirective) Mode() tagex.DirectiveMode {
	return tagex.EvalMode
}

// Handle is where the actual work of the directive takes place. Here you can evaluate the field value or mutate it,
// depending on what value the Mode method returns
func (d *EmailDirective) Handle(val string) (string, error) {
	_, err := mail.ParseAddress(val)
	return val, err
}

// END inferred-directive

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
		if ok, err := checkTag.ProcessStruct(&contact); !ok {
			fmt.Printf("Invalid email %q: %v\n", contact.Email, err)
			continue
		}
		fmt.Printf("Valid email %q\n", contact.Email)
	}
}

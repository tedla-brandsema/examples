package main

import (
	"crypto/sha256"
	"fmt"
	"github.com/tedla-brandsema/tagex"
	"net/mail"
	"os"
	"strings"
	"text/tabwriter"
	"time"
	"unicode"
)

// User implements both the tagex.PreProcessor and tagex.PostProcessor interfaces.
type User struct {
	Username string
	Email    string `validate:"email"`
	Password string `validate:"password, min=8"`
	Updated  time.Time
}

// Before implements the PreProcessor interface:
// Here we can enforce prerequisites, in this case required fields.
func (u *User) Before() error {
	var err error

	if u.Username, err = requiredField("Username", u.Username); err != nil {
		return err
	}

	if u.Email, err = requiredField("Email", u.Email); err != nil {
		return err
	}

	if u.Password, err = requiredField("Password", u.Password); err != nil {
		return err
	}

	return nil
}

// PostProcessor implementation
func (u *User) After() error {
	// Hash the password
	// WARNING: Oversimplified hashing of password; do not use in code meant for real world use.
	h := sha256.New()
	h.Write([]byte(u.Password))
	u.Password = string(h.Sum(nil))

	u.Updated = time.Now().UTC()

	return nil
}

func requiredField(fieldName, fieldValue string) (string, error) {
	trimmed := strings.TrimSpace(fieldValue)
	if trimmed == "" {
		return "", fmt.Errorf("%q is a required field", fieldName)
	}
	return trimmed, nil
}

// EmailDirective validates email strings
type EmailDirective struct{}

func (d *EmailDirective) Name() string {
	return "email"
}

func (d *EmailDirective) Mode() tagex.DirectiveMode {
	return tagex.EvalMode
}

func (d *EmailDirective) Handle(val string) (string, error) {
	var err error
	_, err = mail.ParseAddress(val)
	return val, err
}

// PasswordDirective enforces required characters for the given password
type PasswordDirective struct {
	Min int `param:"min"`
}

func (d *PasswordDirective) Name() string {
	return "password"
}

func (d *PasswordDirective) Mode() tagex.DirectiveMode {
	return tagex.EvalMode
}

func (d *PasswordDirective) Handle(val string) (string, error) {
	if len(val) < d.Min {
		return val, fmt.Errorf("%s needs to have a minimum length of %d characters: current length is %d",
			d.Name(),
			d.Min,
			len(val))
	}

	var hasUpper, hasDigit bool
	for _, r := range val {
		switch {
		case unicode.IsUpper(r):
			hasUpper = true
		case unicode.IsDigit(r):
			hasDigit = true
		}
		if hasUpper && hasDigit {
			return val, nil // password is valid
		}
	}
	if !hasUpper {
		return val, fmt.Errorf("%s needs to contain at least one uppercase character", d.Name())
	}
	if !hasDigit {
		return val, fmt.Errorf("%s needs to contain at least one digit", d.Name())
	}

	return val, nil
}

func printUser(user User) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	_, _ = fmt.Fprintf(w, "Username:\t%s\n", user.Username)
	_, _ = fmt.Fprintf(w, "Email:\t%s\n", user.Email)
	_, _ = fmt.Fprintf(w, "Password:\t%s\n", user.Password)
	_, _ = fmt.Fprintf(w, "Updated:\t%s\n", user.Updated.Format(time.RFC822))
	_ = w.Flush()
}

func main() {
	// Create a User
	user := User{
		Username: "TagexUser",
		Email:    "me@example.com",
		Password: "SuperSecret1",
	}

	// Create a tag "validate"
	validateTag := tagex.NewTag("validate")
	// Register the "EmailDirective" with our "validateTag"
	tagex.RegisterDirective(&validateTag, &EmailDirective{})
	// Register the "PasswordDirective" with our "validateTag"
	tagex.RegisterDirective(&validateTag, &PasswordDirective{})

	// Process the "User" struct we created. ProcessStruct will invoke the "Before()", "Handle()" and "After()" methods
	if ok, err := validateTag.ProcessStruct(&user); !ok {
		fmt.Println(err)
		return
	}

	printUser(user)
}

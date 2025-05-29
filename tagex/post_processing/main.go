package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/tedla-brandsema/tagex"
)

type Timekeeper struct {
	UTC   time.Time
	Milli int64
}

func (t *Timekeeper) After() error {
	if t.UTC.IsZero() {
		return errors.New("field London is not set")
	}

	t.Milli = t.UTC.UnixMilli()
	return nil
}

func main() {
	t := &Timekeeper{
		UTC: time.Now().UTC(),
	}

	if err := tagex.InvokePostProcessor(t); err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", t)
}

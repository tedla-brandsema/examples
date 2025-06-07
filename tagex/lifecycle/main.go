package main

import (
	"fmt"
	"github.com/tedla-brandsema/tagex"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"os"
	"path"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"
)

type Article struct {
	Title   string `edit:"title"`
	created time.Time
	path    string
}

// Implementing the PreProcessor interface
func (t *Article) Before() error {
	if t.created.IsZero() {
		t.created = time.Now().UTC()
	}

	return nil
}

// Implementing the PostProcessor interface
func (t *Article) After() error {
	// Oversimplified slugifying of article.Title (does not handle diacritics or other separators)
	slug := strings.ToLower(strings.ReplaceAll(t.Title, " ", "-"))

	year, month, _ := t.created.Date()
	t.path = path.Join("articles", strings.ToLower(month.String()), strconv.Itoa(year), slug)

	return nil
}

type TitleDirective struct{}

func (d *TitleDirective) Name() string {
	return "title"
}
func (d *TitleDirective) Handle(val string) (string, error) {
	return cases.Title(language.English, cases.Compact).String(val), nil
}

func main() {
	editTag := tagex.NewTag("edit")
	tagex.RegisterDirective(&editTag, &TitleDirective{})

	article := &Article{
		Title: "title of my article",
	}

	if ok, err := editTag.ProcessStruct(article); !ok {
		panic(err)
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintf(w, "Title:\t%s\n", article.Title)
	fmt.Fprintf(w, "Created:\t%s\n", article.created.Format(time.RFC822))
	fmt.Fprintf(w, "Path:\t%s\n", article.path)
	w.Flush()
}

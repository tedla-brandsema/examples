package main

import (
	"errors"
	"fmt"
	"github.com/tedla-brandsema/tagex"
	"os"
	"path"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"
)

// Article implements both the tagex.PreProcessor and tagex.PostProcessor interfaces.
type Article struct {
	Title   string
	Body    string
	created time.Time
	path    string
}

// PreProcessor implementation
func (t *Article) Before() error {
	if t.created.IsZero() {
		t.created = time.Now().UTC()
	}

	return nil
}

// PostProcessor implementation
func (t *Article) After() error {
	if t.Title == "" {
		return errors.New("field Title is empty")
	}
	// Oversimplified slugifying of article.Title (does not handle diacritics nor separators other than space)
	slug := strings.ToLower(strings.ReplaceAll(t.Title, " ", "-"))

	if t.created.IsZero() {
		return errors.New("field created is empty")
	}
	year, month, _ := t.created.Date()
	t.path = path.Join("articles", strings.ToLower(month.String()), strconv.Itoa(year), slug)

	return nil
}

func printArticle(article Article) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	_, _ = fmt.Fprintf(w, "Title:\t%s\n", article.Title)
	_, _ = fmt.Fprintf(w, "Created:\t%s\n", article.created.Format(time.RFC822))
	_, _ = fmt.Fprintf(w, "Path:\t%s\n", article.path)
	_, _ = fmt.Fprintf(w, "Body:\t%s\n", article.Body)
	_ = w.Flush()
}

func main() {
	// Create an Article instance
	article := Article{
		Title: "Article Title",
		Body:  "Article body.",
	}

	// Invoking the PreProcessor
	if err := tagex.InvokePreProcessor(&article); err != nil {
		panic(err)
	}

	// Invoking the PostProcessor
	if err := tagex.InvokePostProcessor(&article); err != nil {
		panic(err)
	}

	printArticle(article)
}

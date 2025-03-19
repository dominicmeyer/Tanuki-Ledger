package components_test

import (
	"context"
	"io"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/a-h/templ"
)

func RenderComponent(t *testing.T, component templ.Component) *goquery.Document {
	t.Helper()

	reader, writer := io.Pipe()
	go func() {
		_ = component.Render(context.Background(), writer)
		_ = writer.Close()
	}()

	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		t.Fatalf("failed to read template: %v", err)
	}

	return doc
}

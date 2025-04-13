package components_test

import (
	"fmt"
	"testing"

	"github.com/a-h/templ"
	"github.com/dominicmeyer/Tanuki-Ledger/internal/components"
	"github.com/stretchr/testify/assert"
)

func TestBasePage_DefinesDoctype(t *testing.T) {
	t.Parallel()

	doc := RenderComponent(t, components.BasePage())
	html, _ := doc.Html()

	assert.Contains(t, html, "<!DOCTYPE html>", "BasePage should include the doctype")
}

func TestBasePage_DefinesCharset(t *testing.T) {
	t.Parallel()

	doc := RenderComponent(t, components.BasePage())
	headHTML, _ := doc.Find("head").Html()

	assert.Contains(t, headHTML, "<meta charset=\"UTF-8\"/>", "BasePage should include a charset")
}

func TestBasePage_DefinesViewport(t *testing.T) {
	t.Parallel()

	doc := RenderComponent(t, components.BasePage())
	headHTML, _ := doc.Find("head").Html()

	assert.Contains(
		t,
		headHTML,
		"<meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"/>",
		"BasePage should include a viewport",
	)
}

func TestBasePage_LinksTailwindOutput(t *testing.T) {
	t.Parallel()

	doc := RenderComponent(t, components.BasePage())
	headHTML, _ := doc.Find("head").Html()

	assert.Contains(
		t,
		headHTML,
		"<link href=\"/static/css/tailwind.css\" rel=\"stylesheet\"/>",
		"BasePage should link to tailwind output",
	)
}

func TestBasePage_LinksHtmxCdn(t *testing.T) {
	t.Parallel()

	doc := RenderComponent(t, components.BasePage())
	headHTML, _ := doc.Find("head").Html()

	htmxScriptSrc := "src=\"https://unpkg.com/htmx.org@[0-9]*.[0-9]*.[0-9]*/dist/htmx.js\""

	assert.Regexp(
		t,
		fmt.Sprintf("<script %s integrity=\"[^\"]*\" crossorigin=\"anonymous\"></script>", htmxScriptSrc),
		headHTML,
		"BasePage should link to htmx cdn",
	)
}

func TestBasePage_IncludesInput(t *testing.T) {
	t.Parallel()

	doc := RenderComponent(t, components.BasePage(
		templ.Raw("<div data-testid=\"test-div-1\">Testing Input 1</div>"),
		templ.Raw("<div data-testid=\"test-div-2\">Testing Input 2</div>"),
	))

	assert.Equal(t, 1, doc.Find(`[data-testid="test-div-1"]`).Length())
	assert.Equal(t, "Testing Input 1", doc.Find(`[data-testid="test-div-1"]`).Text())
	assert.Equal(t, 1, doc.Find(`[data-testid="test-div-2"]`).Length())
	assert.Equal(t, "Testing Input 2", doc.Find(`[data-testid="test-div-2"]`).Text())
}

func TestBasePage_DefinesMain(t *testing.T) {
	t.Parallel()

	doc := RenderComponent(t, components.BasePage())
	assert.Equal(t, 1, doc.Find("main").Length())
}

func TestBasePage_DefinesLanguage(t *testing.T) {
	t.Parallel()

	doc := RenderComponent(t, components.BasePage())
	attributes := make(map[string]string)

	for _, attribute := range doc.Find("html").Get(0).Attr {
		attributes[attribute.Key] = attribute.Val
	}

	t.Log(doc.Html())

	assert.NotNil(t, attributes["lang"], "BasePage should define language")
	assert.Equal(t, "en", attributes["lang"], "BasePage should define english as language")
}

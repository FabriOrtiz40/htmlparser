package link

import (
	"strings"
	"testing"
)

func TestParseSingleLink(t *testing.T) {
	html := `<a href="/dog">Dog</a>`

	links, err := Parse(strings.NewReader(html))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(links) != 1 {
		t.Fatalf("expected 1 link, got %d", len(links))
	}

	if links[0].Href != "/dog" {
		t.Errorf("expected href /dog, got %s", links[0].Href)
	}

	if links[0].Text != "Dog" {
		t.Errorf("expected text Dog, got %q", links[0].Text)
	}
}


func TestParseNestedHTML(t *testing.T) {
	html := `
	<a href="/dog">
		<span>Dog</span>
		and
		<b>Cat</b>
	</a>
	`

	links, err := Parse(strings.NewReader(html))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if links[0].Text != "Dog and Cat" {
		t.Errorf("expected 'Dog and Cat', got %q", links[0].Text)
	}
}

func TestIgnoreNestedLinks(t *testing.T) {
	html := `
	<a href="#">
		Outer
		<a href="/dog">Inner</a>
	</a>
	`

	links, err := Parse(strings.NewReader(html))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(links) != 1 {
		t.Fatalf("expected 1 link, got %d", len(links))
	}

	if links[0].Text != "Outer Inner" {
		t.Errorf("unexpected text: %q", links[0].Text)
	}
}

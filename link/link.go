package link

import (
	"io"
	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

// Recibe HTML desde cualquier fuente y parsea. Devuelve todos los links encontrados
func Parse(r io.Reader) ([]Link, error) {
		// 1. Parsear el HTML
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	// Por ahora no hacemos nada con doc
	_ = doc

	return nil, nil
}

// El Parse usa esta funcion para extraer todos los Links
func extractLinks(n *html.Node, links *[]Link)

// Cuando encuentra un <a> llama a esta función
func buildLink(n *html.Node) Link

// Y a esta
func extractText(n *html.Node) string

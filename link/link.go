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
	// 2. Recorrer el árbol
	// 3. Encontrar <a>
	// 4. Construir []Link
	return nil, nil
}

// El Parse usa esta funcion para extraer todos los Links
func extractLinks(n *html.Node, links *[]Link)

// Cuando encuentra un <a> llama a esta función
func buildLink(n *html.Node) Link

// Y a esta
func extractText(n *html.Node) string

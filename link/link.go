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

	var links []Link
	extractLinks(doc, &links)

	return links, nil
}

// El Parse usa esta funcion para extraer todos los Links
func extractLinks(n *html.Node, links *[]Link) {
	if n == nil {
		return
	}

	// Acá después vamos a hacer cosas con el nodo n
	if n.Type == html.ElementNode && n.Data == "a" {
		var href string

		for _, attr := range n.Attr {
			if attr.Key == "href" {
				href = attr.Val
				break
			}
		}

		link := Link{
			Href: href,
		}

		*links = append(*links, link)
	}

	// DFS: primer hijo
	extractLinks(n.FirstChild, links)

	// DFS: siguiente hermano
	extractLinks(n.NextSibling, links)
}

// Cuando encuentra un <a> llama a esta función
func buildLink(n *html.Node) Link

// Y a esta
func extractText(n *html.Node) string

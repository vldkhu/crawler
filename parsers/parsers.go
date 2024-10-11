package parsers

import (
	"crawler/models"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func ParsePage(url string) ([]models.Stajirovka, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	var stajirovki []models.Stajirovka
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			// Проверяем, есть ли у узла дочерние элементы
			if n.FirstChild != nil && n.FirstChild.Type == html.TextNode {
				title := strings.ToLower(n.FirstChild.Data)
				keywords := []string{"стажировка", "интернатура", "практика", "стажер"}
				for _, keyword := range keywords {
					if strings.Contains(title, keyword) {
						for _, attr := range n.Attr {
							if attr.Key == "href" {
								stajirovki = append(stajirovki, models.Stajirovka{
									Title: n.FirstChild.Data,
									Link:  attr.Val,
								})
							}
						}
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	return stajirovki, nil
}

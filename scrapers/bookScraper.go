package scrapers
import (
	"io"
	"log"
	"net/http"
	"strings"
	"golang.org/x/net/html"
)

func FetchBooks() string {
	response, err := http.Get("http://books.toscrape.com/")
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	html.Parse(strings.NewReader(string(body)))
	if err != nil {
		log.Fatal(err)
	}

	var html = string(body)

	return html
}

func ParseData(data string) []string {
	doc, err := html.Parse(strings.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}

	var f func(*html.Node)
	var items []string
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			// fmt.Println(n)
			for _, a := range n.Attr {
				if a.Key == "title" {
					items = append(items, a.Val)
				}
			}
		}
		// loops over html document, and above the data is processed for example we retrieve a in the markup
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
  
  return items
}

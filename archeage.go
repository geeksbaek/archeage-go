package archeage

import (
	"io"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type ArcheAge struct {
	client *http.Client
}

func New(c *http.Client) *ArcheAge {
	return &ArcheAge{c}
}

func (a *ArcheAge) post(url string, form io.Reader) (*goquery.Document, error) {
	return a.do(url, "POST", form)
}

func (a *ArcheAge) get(url string) (*goquery.Document, error) {
	return a.do(url, "GET", nil)
}

func (a *ArcheAge) do(url, method string, form io.Reader) (*goquery.Document, error) {
	req, err := http.NewRequest(method, url, form)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	resp, err := a.client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		log.Fatal(err)
	}
	return doc, nil
}

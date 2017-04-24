package archeage

import (
	"io"
	"log"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const (
	timeout = time.Second * 2
)

type archeAge struct {
	client *http.Client
}

func ArcheAge(c *http.Client) *archeAge {
	c.Timeout = timeout
	return &archeAge{c}
}

func (a *archeAge) post(url string, form io.Reader) (*goquery.Document, error) {
	return a.do(url, "POST", form)
}

func (a *archeAge) get(url string) (*goquery.Document, error) {
	return a.do(url, "GET", nil)
}

func (a *archeAge) do(url, method string, form io.Reader) (*goquery.Document, error) {
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

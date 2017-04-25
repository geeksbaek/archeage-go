package archeage

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type Expedition struct {
	Name      string
	URL       string
	Server    string
	Number    string
	Nation    string
	Thumbnail string
	Level     int
	Master    *Character
	Members   []*Character
	People    int
}

type Expeditions []Expedition

const (
	searchExpeditionFormat = "https://archeage.xlgames.com/search?dt=expeditions&keyword=%s&server=%s"
	expeditionURLFormat    = "https://archeage.xlgames.com/expeditions/%s/%s"
)

const (
	expRowQuery = "ul.lst > li > div.exped_info"
)

var (
	expeditionRe = regexp.MustCompile(`expeditions\/(\w+)\/`)
)

func (a *ArcheAge) SearchExpedition(server, name string) (e *Expedition, err error) {
	return
}

func (a *ArcheAge) fetchExpeditionByNum(server, number string) (e *Expedition, err error) {
	url := fmt.Sprintf(expeditionURLFormat, server, number)
	return a.fetchExpeditionByURL(url)
}

func (a *ArcheAge) fetchExpeditionByURL(url string) (e *Expedition, err error) {
	doc, err := a.get(url)
	if err != nil {
		return
	}
	return a.parseExpedition(doc, url)
}

func (a *ArcheAge) parseExpedition(doc *goquery.Document, url string) (e *Expedition, err error) {
	e.Name = doc.Find("span.name").Text()
	e.URL = url
	if matched := expeditionRe.FindStringSubmatch(e.URL); len(matched) == 2 {
		e.Server = matched[1]
	} else {
		return nil, errors.New("Unknown Server")
	}
	e.Number = regexp.MustCompile(`/\w+/\w+/(\d+)`).FindStringSubmatch(e.URL)[1]
	e.Nation = doc.Find("span.nation > em, span.group > em").Text()
	e.Thumbnail, _ = doc.Find("span.exped_thumb > img").Attr("src")
	e.Level, _ = strconv.Atoi(doc.Find("div.sub > span.lv > span.count").Text())
	// masterUUID, _ := doc.Find("a.character-link").Attr("data-uuid")
	// expedition.Master = fetchCharacter(masterUUID)
	// expedition.Members
	// expedition.People
	return
}

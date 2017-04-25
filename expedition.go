package archeage

import (
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
)

const (
	expRowQuery = "ul.lst > li > div.exped_info"
)

var (
	expeditionRe = regexp.MustCompile(`expeditions\/(\w+)\/`)
)

func (a *ArcheAge) FetchAllExpeditions(server string) {
	expListURL := fmt.Sprintf("https://archeage.xlgames.com/exps/all?gameServer=%s", server)

	doc, err := a.get(expListURL)
	if err != nil {
		return
	}

	expeditions := Expeditions{}

	doc.Find(expRowQuery).Each(func(i int, row *goquery.Selection) {
		URL, _ := row.Find("a").First().Attr("href")
		expeditions = append(expeditions, a.fetchSingleExpeditionByURL(URL))
	})
}

func (a *ArcheAge) SearchExpedition(server, name string) {

}

func (a *ArcheAge) fetchSingleExpeditionByParam(server, number string) Expedition {
	URL := fmt.Sprintf("https://archeage.xlgames.com/expeditions/%s/%s", server, number)
	return a.fetchSingleExpeditionByURL(URL)
}

func (a *ArcheAge) fetchSingleExpeditionByURL(URL string) (expedition Expedition) {
	doc, err := a.get(URL)
	if err != nil {
		return
	}

	expedition.Name = doc.Find("span.name").Text()
	expedition.URL = URL
	expedition.Server = func() string {
		if matched := expeditionRe.FindStringSubmatch(URL); len(matched) == 2 {
			return matched[1]
		}
		return "Unknown"
	}()
	expedition.Number = regexp.MustCompile(`/\w+/\w+/(\d+)`).FindStringSubmatch(expedition.URL)[1]
	expedition.Nation = doc.Find("span.nation > em, span.group > em").Text()
	expedition.Thumbnail, _ = doc.Find("span.exped_thumb > img").Attr("src")
	expedition.Level, _ = strconv.Atoi(doc.Find("div.sub > span.lv > span.count").Text())
	// masterUUID, _ := doc.Find("a.character-link").Attr("data-uuid")
	// expedition.Master = fetchCharacter(masterUUID)
	// expedition.Members
	// expedition.People

	return
}

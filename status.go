package archeage

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const serverStatusURL = "https://archeage.xlgames.com/serverstatus"

const (
	serverStatusRowQuery = `table tr`
)

type ServerStatus map[string]bool

func (old ServerStatus) DiffString(new ServerStatus) (formattedString string) {
	for k := range new {
		if old[k] != new[k] {
			var line string
			if old[k] == false && new[k] == true {
				line = fmt.Sprintf("[%v] 서버 열림\n", k)
			} else {
				line = fmt.Sprintf("[%v] 서버 닫힘\n", k)
			}
			formattedString += line
		}
	}
	return
}

func (a *ArcheAge) FetchServerStatus() (serverStatus ServerStatus, err error) {
	doc, err := a.get(serverStatusURL)
	if err != nil {
		return
	}
	serverStatus = ServerStatus{}
	doc.Find(serverStatusRowQuery).Each(func(i int, s *goquery.Selection) {
		name := strings.TrimSpace(s.Find(".server").Text())
		if name == "" {
			return
		}
		status := s.Find(".stats span").HasClass("on")
		serverStatus[name] = status
	})
	return
}

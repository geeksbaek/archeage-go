package archeage

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const ServerStatusURL = "https://archeage.xlgames.com/serverstatus"

const (
	serverStatusRowQuery = `table tr`
)

type ServerStatus map[string]bool

func ParseServerStatus(doc *goquery.Document) (serverStatus ServerStatus, err error) {
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

package archeage

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Notice struct {
	Category    string
	Title       string
	Description string
	URL         string
	Date        string
}

type Notices []Notice

func (old Notices) diff(new Notices) (diff Notices) {
	// newnotice, oldnotice are sorted by recently.
	for _, newnotice := range new {
		for _, oldnotice := range old {
			if newnotice.URL == oldnotice.URL {
				return
			}
		}
		// If element is not found in oldnotice, add it to diff.
		diff = append(diff, newnotice)
	}
	return
}

type noticeCategory struct {
	Name   string
	URL    string
	Parser NoticeParser
}

type NoticeParser func(*goquery.Document) Notices

// query
const (
	noticeCategoryQuery = `.cont_head h2`
	basicNoticeRowQuery = `.news tbody tr`
	eventNoticeRowQuery = `ul.list-event li a`
)

const urlPrefix = "https://archeage.xlgames.com"

func BasicNoticeParser(doc *goquery.Document) (notices Notices) {
	categoryName := strings.TrimSpace(doc.Find(noticeCategoryQuery).Text())
	doc.Find(basicNoticeRowQuery).Each(func(i int, row *goquery.Selection) {
		var notice Notice

		notice.Category = categoryName
		if row.Find("a.pjax .tit, a.pjax strong.thumb-tit").Length() > 0 {
			notice.Title = strings.TrimSpace(row.Find("a.pjax .tit, a.pjax strong.thumb-tit").Text())
		} else {
			notice.Title = strings.TrimSpace(row.Find("a.pjax").Text())
		}
		notice.Description = strings.TrimSpace(row.Find("a.pjax .txt, a.pjax span.thumb-txt").Text())
		notice.URL, _ = row.Find("a.pjax").Attr("href")
		notice.URL = urlPrefix + strings.Split(notice.URL, "?")[0]
		notice.Date = strings.TrimSpace(row.Find("td.time").Text())

		notices = append(notices, notice)
	})
	return
}

func EventNoticeParser(doc *goquery.Document) (notices Notices) {
	categoryName := strings.TrimSpace(doc.Find(noticeCategoryQuery).Text())
	doc.Find(eventNoticeRowQuery).Each(func(i int, row *goquery.Selection) {
		var notice Notice

		notice.Category = categoryName
		notice.Title = strings.TrimSpace(row.Find("dl dt").Text())
		notice.Description = strings.TrimSpace(row.Find("dl dd:not(.img)").Text())
		notice.URL, _ = row.Attr("href")
		notice.URL = urlPrefix + strings.Split(notice.URL, "?")[0]

		notices = append(notices, notice)
	})
	return
}

type DocWithParser struct {
	doc    *goquery.Document
	parser NoticeParser
}

func ParseNotices(dps ...DocWithParser) (notices Notices) {
	for _, dp := range dps {
		notices = append(notices, dp.parser(dp.doc)...)
	}
	return
}

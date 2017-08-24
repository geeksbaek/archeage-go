package archeage

import (
	"fmt"
	"strings"

	"regexp"

	"github.com/PuerkitoBio/goquery"
)

type Notice struct {
	Category    string
	Title       string
	Description string
	URL         string
	Date        string
}

func (n Notice) String() string {
	return fmt.Sprintf("[%v] %v %v\n", n.Category, n.Title, n.URL)
}

type Notices []Notice

func (ns Notices) contains(n Notice) bool {
	for _, v := range ns {
		if v.URL == n.URL && v.Title == n.Title {
			return true
		}
	}
	return false
}

func (ns Notices) String() (ret string) {
	for _, v := range ns {
		ret += v.String()
	}
	return
}

func (old Notices) Diff(new Notices) (diff Notices) {
	for _, n := range new {
		if !old.contains(n) {
			diff = append(diff, n)
		}
	}
	return
}

func (old Notices) Merge(new Notices) (merged Notices) {
	merged = old
	for _, v := range new {
		if !old.contains(v) {
			merged = append(merged, v)
		}
	}
	return
}

type noticeCategory struct {
	URL    string
	Parser noticeParser
}

type noticeParser func(*goquery.Document) Notices

// query
const (
	noticeCategoryQuery       = `.cont_head h2`
	basicNoticeRowQuery       = `.news tbody tr`
	eventNoticeRowQuery       = `ul.list-event li a`
	eventWinnerNoticeRowQuery = `.notice`
)

const urlPrefix = "https://archeage.xlgames.com"

var allNoticesCategory = []noticeCategory{
	{"https://archeage.xlgames.com/mboards/notice", basicNoticeParser},
	{"https://archeage.xlgames.com/mboards/patchnote", basicNoticeParser},
	{"https://archeage.xlgames.com/mboards/inside", basicNoticeParser},
	{"https://archeage.xlgames.com/mboards/amigo", basicNoticeParser},
	{"https://archeage.xlgames.com/events", eventNoticeParser},
	{"https://archeage.xlgames.com/events/winner", eventWinnerNoticeParser},
}

var multiLineReg = regexp.MustCompile(`\n+`)

func basicNoticeParser(doc *goquery.Document) (notices Notices) {
	categoryName := strings.TrimSpace(doc.Find(noticeCategoryQuery).Text())
	doc.Find(basicNoticeRowQuery).Each(func(i int, row *goquery.Selection) {
		var notice Notice

		notice.Category = categoryName
		notice.Category = regexp.MustCompile(`[\n\t]+`).ReplaceAllString(notice.Category, ": ") // for 아미고
		if row.Find("a.pjax .tit, a.pjax strong.thumb-tit").Length() > 0 {
			notice.Title = strings.TrimSpace(row.Find("a.pjax .tit, a.pjax strong.thumb-tit").Text())
		} else {
			notice.Title = strings.TrimSpace(row.Find("a.pjax").Text())
		}
		notice.Title = multiLineReg.ReplaceAllString(notice.Title, " ")
		notice.Description = strings.TrimSpace(row.Find("a.pjax .txt, a.pjax span.thumb-txt").Text())
		notice.URL, _ = row.Find("a.pjax").Attr("href")
		notice.URL = urlPrefix + strings.Split(notice.URL, "?")[0]
		notice.Date = strings.TrimSpace(row.Find("td.time").Text())

		notices = append(notices, notice)
	})
	return
}

func eventNoticeParser(doc *goquery.Document) (notices Notices) {
	categoryName := strings.TrimSpace(doc.Find(noticeCategoryQuery).Text())
	doc.Find(eventNoticeRowQuery).Each(func(i int, row *goquery.Selection) {
		var notice Notice

		notice.Category = categoryName
		notice.Title = strings.TrimSpace(row.Find(".cont").Text())
		notice.Title = multiLineReg.ReplaceAllString(notice.Title, " ")
		notice.Description = strings.TrimSpace(row.Find(".time").Text())
		notice.URL, _ = row.Attr("href")
		notice.URL = urlPrefix + strings.Split(notice.URL, "?")[0]

		notices = append(notices, notice)
	})
	return
}

func eventWinnerNoticeParser(doc *goquery.Document) (notices Notices) {
	categoryName := strings.TrimSpace(doc.Find(noticeCategoryQuery).Text())
	doc.Find(eventWinnerNoticeRowQuery).Each(func(i int, row *goquery.Selection) {
		var notice Notice

		notice.Category = categoryName
		notice.Title = strings.TrimSpace(row.Find(".cont").Text())
		notice.Title = strings.TrimLeft(notice.Title, "[이벤트] ")
		notice.Title = multiLineReg.ReplaceAllString(notice.Title, " ")
		notice.Description = strings.TrimSpace(row.Find(".time").Text())
		notice.URL, _ = row.Find("a").Attr("href")
		notice.URL = urlPrefix + strings.Split(notice.URL, "?")[0]

		notices = append(notices, notice)
	})
	return
}

func (a *ArcheAge) FetchNotice() (notices Notices, err error) {
	for _, nc := range allNoticesCategory {
		if doc, err := a.get(nc.URL); err == nil {
			notices = append(notices, nc.Parser(doc)...)
		}
	}
	return
}

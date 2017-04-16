package archeage

import (
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func MustDoc(url string) *goquery.Document {
	doc, _ := goquery.NewDocument(url)
	return doc
}

var cases = []DocWithParser{
	{MustDoc("https://archeage.xlgames.com/mboards/notice"), BasicNoticeParser},
	{MustDoc("https://archeage.xlgames.com/mboards/patchnote"), BasicNoticeParser},
	{MustDoc("https://archeage.xlgames.com/events"), EventNoticeParser},
	{MustDoc("https://archeage.xlgames.com/mboards/inside"), BasicNoticeParser},
	{MustDoc("https://archeage.xlgames.com/mboards/amigo"), BasicNoticeParser},
}

func Test(t *testing.T) {
	for _, c := range cases {
		if notices := ParseNotices(c); len(notices) == 0 {
			t.Error("Empty Notice")
		}
	}
}

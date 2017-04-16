package archeage

import (
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func MustNoticeDoc(url string) *goquery.Document {
	doc, _ := goquery.NewDocument(url)
	return doc
}

var cases = []DocWithParser{
	{MustNoticeDoc("https://archeage.xlgames.com/mboards/notice"), BasicNoticeParser},
	{MustNoticeDoc("https://archeage.xlgames.com/mboards/patchnote"), BasicNoticeParser},
	{MustNoticeDoc("https://archeage.xlgames.com/events"), EventNoticeParser},
	{MustNoticeDoc("https://archeage.xlgames.com/mboards/inside"), BasicNoticeParser},
	{MustNoticeDoc("https://archeage.xlgames.com/mboards/amigo"), BasicNoticeParser},
}

func TestParseNotices(t *testing.T) {
	for _, c := range cases {
		if notices := ParseNotices(c); len(notices) == 0 {
			t.Error("Empty Notice")
		}
	}
}

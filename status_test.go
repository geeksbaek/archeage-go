package archeage

import (
	"testing"

	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func TestParseServerStatus(t *testing.T) {
	doc, err := goquery.NewDocument(ServerStatusURL)
	if err != nil {
		t.Error(err)
	}
	ss, err := ParseServerStatus(doc)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ss)
}

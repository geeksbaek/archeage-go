package archeage

import (
	"net/http"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func MustAuctionDoc(ap AuctionSearchParam) *goquery.Document {
	resp, _ := http.Post(AuctionURL, "", MakeAuctionSearchForm(ap))
	doc, _ := goquery.NewDocumentFromResponse(resp)
	return doc
}

func TestParseAction(t *testing.T) {
	doc := MustAuctionDoc(AuctionSearchParam{"KYPROSA", "", "목재"})
	if auctionResult := ParseAuction(doc); len(auctionResult) == 0 {
		t.Error("Empty Auction Result")
	}
}

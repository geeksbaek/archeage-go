package archeage

import (
	"errors"
	"strconv"
	"strings"

	"fmt"

	"github.com/PuerkitoBio/goquery"
)

// AuctionSearchResult 구조체는 경매장 검색 결과를 표현합니다.
type AuctionSearchResult struct {
	Name        string
	Quantity    int
	Image       string
	TotalPrice  Price
	SinglePrice Price
}

// Price 구조체는 금, 은, 동으로 이루어진 가격 정보를 표현합니다.
type Price struct {
	Gold   int
	Silver int
	Bronze int
}

// Int 메소드는 Price 구조체를 정수로 변환합니다.
func (p Price) Int() int {
	return p.Bronze + p.Silver*100 + p.Gold*10000
}

func (p Price) String() (ret string) {
	if p.Gold != 0 {
		ret += fmt.Sprintf("%v금 ", p.Gold)
	}
	if p.Silver != 0 {
		ret += fmt.Sprintf("%v은 ", p.Silver)
	}
	if p.Bronze != 0 {
		ret += fmt.Sprintf("%v동 ", p.Bronze)
	}
	return
}

// IntPrice 타입은 정수에서 Price 타입으로 변환하기 위한 메소드를 붙이기 위한 타입입니다.
type IntPrice int

// Price 메소드는 정수를 Price 타입으로 변환합니다.
func (i IntPrice) Price() Price {
	return Price{
		Gold:   int(i) / 10000,
		Silver: (int(i) % 10000) / 100,
		Bronze: (int(i) % 10000) % 100,
	}
}

// func (p Price) Add(p2 Price) (ret Price) {}
// func (p Price) Sub(p2 Price) (ret Price) {}

func (p Price) Mul(n int) (ret Price) {
	return IntPrice(p.Int() * n).Price()
}

// Div 메소드는 Price의 값을 주어진 정수로 나눕니다.
func (p Price) Div(n int) (ret Price) {
	return IntPrice(p.Int() / n).Price()
}

// url
const (
	auctionURL = "https://archeage.xlgames.com/auctions/list/ajax"
)

// query
const (
	auctionRowQuery = `.tlist`
	nameQuery       = `.name`
	priceQuery      = `.auction-bidmoney > .buybid em.gol_num`
	quantityQuery   = `.item-num`
	imageQuery      = `.eq_img img`
)

func (a *ArcheAge) Auction(serverGroup, itemName string) ([]AuctionSearchResult, error) {
	searchForm := form(map[string]string{
		"sortType":     "BUYOUT_PRICE_ASC",
		"searchType":   "NAME",
		"serverCode":   serverGroup,
		"keyword":      itemName,
		"equalKeyword": "false",
	})

	doc, err := a.post(auctionURL, searchForm)
	if err != nil {
		return nil, err
	}

	searchResults := []AuctionSearchResult{}

	doc.Find(auctionRowQuery).Each(func(i int, row *goquery.Selection) {
		var searchResult AuctionSearchResult
		var err error

		// get price
		sumIntPrice := 0
		row.Find(priceQuery).Each(func(i int, moneyCell *goquery.Selection) {
			n, _ := strconv.Atoi(strings.Replace(moneyCell.Text(), ",", "", -1))
			sumIntPrice = (sumIntPrice * 100) + n
		})
		searchResult.Name = row.Find(nameQuery).Text()
		searchResult.TotalPrice = IntPrice(sumIntPrice).Price()
		if searchResult.Quantity, err = strconv.Atoi(row.Find(quantityQuery).Text()); err != nil {
			searchResult.Quantity = 1
		}
		searchResult.Image, _ = row.Find(imageQuery).Attr("src")
		searchResult.SinglePrice = searchResult.TotalPrice.Div(searchResult.Quantity)

		if searchResult.Image == "" {
			return
		}

		searchResults = append(searchResults, searchResult)
	})

	if len(searchResults) == 0 {
		return nil, errors.New("Empty Result")
	}
	return searchResults, nil
}

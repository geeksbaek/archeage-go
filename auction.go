package archeage

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

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

// AuctionSearchResults 타입은 AuctionSearchResult 타입의 슬라이스입니다.
type AuctionSearchResults []*AuctionSearchResult

// MinimumPrice 메소드는 경매장 검색 결과에서 해당 수량만큼의 최소 가격을 반환합니다.
func (r AuctionSearchResults) MinimumPrice(quantity int) (p Price) {
	// TODO.
	// AuctionSearchResults 슬라이스는 개별 가격이 낮은 순으로 정렬되어 입력됩니다.
	// 각 슬라이스에는 가격와 수량이 각각 SinglePrice, Quantity 라는 멤버 변수에 저장되어 있습니다.
	// 각 슬라이스를 순회하여 아이템을 낱개로 구매할 수 있다는 가정 하에,
	// 입력받은 quantity 만큼의 아이템 가격을 반환하세요.

	return
}

// MaximumPrice 메소드는 경매장 검색 결과에서 해당 수량만큼의 최대 가격을 반환합니다.
func (r AuctionSearchResults) MaximumPrice(quantity int) (p Price, overQuantity int) {
	// TODO.
	// 각 슬라이스를 순회하여 아이템을 낱개로 구매할 수 없다는 가정 하에,
	// 입력받은 quantity 만큼의 아이템 가격과 함께 초과하여 구매한 수량을 반환하세요.

	return
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

func (a *ArcheAge) Auction(serverGroup, itemName string) (AuctionSearchResults, error) {
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

	searchResults := AuctionSearchResults{}

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
		searchResult.Image = func() string {
			src, exists := row.Find(imageQuery).Attr("src")
			if exists {
				src = "https:" + src
			}
			return src
		}()
		searchResult.SinglePrice = searchResult.TotalPrice.Div(searchResult.Quantity)

		if searchResult.Image == "" {
			return
		}

		searchResults = append(searchResults, &searchResult)
	})

	if len(searchResults) == 0 {
		return nil, errors.New("Empty Result")
	}
	return searchResults, nil
}

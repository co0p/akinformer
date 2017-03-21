package asck

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/context"
	"google.golang.org/appengine/urlfetch"
)

// Crawler is responsible for crawling
type Crawler struct {
	ctx           context.Context
	url, selector string
}

// Run does the actual crawling and returns a bunch of extracted offers
func (c *Crawler) Run() ([]Offer, error) {

	client := urlfetch.Client(c.ctx)
	resp, err := client.Get(c.url)

	if err != nil {
		return nil, fmt.Errorf("Failed connecting to the offer page: %v", err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed reading body from the response: %v", err.Error())
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("Failed parsing body into html document: %v", err.Error())
	}

	offers := []Offer{}
	doc.Find("h3 > a[name=\"Weiterbildungsstellenangebote\"]").Parent().SiblingsFiltered("table").Last().Find("tr").Each(func(i int, s *goquery.Selection) {
		// TODO: STILL BROKEN SELECTOR
		text := s.Find("td").Text()
		address := s.Find("td").Text()
		date := s.Find("td").Text()

		newOffer := Offer{text: text, address: address, offerDate: date, dateFound: time.Now()}
		offers = append(offers, newOffer)
	})

	return offers, nil
}

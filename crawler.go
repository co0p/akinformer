package asck

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/appengine/urlfetch"
)

// Crawler is responsible for crawling
type Crawler struct {
	ctx context.Context
	url string
}

// Run does the actual crawling and returns a bunch of extracted offers
func (c *Crawler) Run() ([]Offer, error) {

	// fetch url

	// convert data to xml

	// extract tokens and construct offers

	// return offers
	fmt.Printf("starting to crawl '%T' for entries", c.url)
	client := urlfetch.Client(c.ctx)
	resp, err := client.Get(c.url)
	defer resp.Body.Close()

	if err != nil {
		return nil, errors.New("Failed connecting to the offer page: " + err.Error())
	}
	log.Println("connected.")

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Failed reading body from the response: " + err.Error())
	}
	log.Println("read data.")

	html := string(body)
	if len(html) == 0 {
		return nil, errors.New("No characters were read from the response: " + err.Error())
	}

	log.Println("read data.")

	return []Offer{}, nil
}

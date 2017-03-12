package asck

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/net/context"
	"golang.org/x/net/html"
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

	z := html.NewTokenizer(bytes.NewReader(body))
	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			log.Println("EXIT TOKENIZER")
			break
		}

		if name, _ := z.TagName(); "h3" == string(name) {
			log.Println("Found h3")
		}
	}

	return []Offer{}, nil
}

// Package akinformer is a little tool that parses a website and extracts job offers.
// New offers (based on dated created) are being send via mail to me
package akinformer

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"golang.org/x/net/context"

	"github.com/PuerkitoBio/goquery"

	"google.golang.org/appengine"
	"google.golang.org/appengine/mail"
	"google.golang.org/appengine/urlfetch"
)

const (
	url            string = "https://www.aerztekammer-berlin.de/10arzt/15_Weiterbildung/17WB-Stellenboerse/index.html"
	email                 = "julian.godesa@googlemail.com"
	selector              = ".tAngebot tr.job"
	confirmMessage string = `
We have found a new job offer for you!

go to: %s !

Added: %v
============================
%s
----------------------------
%s
----------------------------
%s`
)

// Offer represents an offer found on the website
type Offer struct {
	Specialization, Description, Address string
	DateCreated                          time.Time
}

// String returns a string representation of the offer o
func (o Offer) String() string {
	max := 25
	if len(o.Description) < max {
		max = len(o.Description)
	}
	summary := o.Description[:max]
	return fmt.Sprintf("date:\t%s\ndesc:\t%s...\naddr:\t%s\n", o.DateCreated, summary, o.Address)
}

// init is the entry point for this app engine app; no main !
func init() {
	http.HandleFunc("/api/check", handler)
	http.HandleFunc("/api/subscribe", subscribeHandler)
	http.HandleFunc("/api/unsubscribe", unsubscribeHandler)
	http.Handle("/", http.NotFoundHandler())
}

// handler is triggereing the actual action and will be called every one in a while
func handler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	log := LoggerWithContext(c)

	beforeDate := time.Now().AddDate(0, 0, -1)

	offers, err := parseURL(c, url, selector)
	if err != nil {
		log.Errorf("failed fetching offers: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	nSend, err := sendNewOffers(c, offers, email, beforeDate)
	if err != nil {
		log.Errorf("failed sending new offers: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Found %d offers and send %d to '%s'.", len(offers), nSend, email)
}

func parseURL(c context.Context, url string, selector string) ([]Offer, error) {
	log := LoggerWithContext(c)
	client := urlfetch.Client(c)
	resp, err := client.Get(url)
	offers := []Offer{}

	if err != nil {
		return nil, fmt.Errorf("failed connecting to the offer page: %v", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading body from the response: %v", err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("failed parsing body into html document: %v", err)
	}

	doc.Find(selector).Each(func(i int, s *goquery.Selection) {

		address := s.Find("td:nth-child(1)").Text()
		specialization := s.Find("td:nth-child(2)").Text()
		description := s.Find("td:nth-child(3)").Text()
		dateCreated := s.Find("td:nth-child(4)").Text()
		t, err := time.Parse("02.01.2006", strings.TrimSpace(dateCreated))
		if err != nil {
			log.Errorf("Failed parsing date of item %d, with err: %s! continue ...", i, err)
			return
		}

		o := Offer{strings.TrimSpace(specialization), strings.TrimSpace(description), strings.TrimSpace(address), t}
		offers = append(offers, o)
	})

	return offers, nil
}

func sendNewOffers(c context.Context, offers []Offer, address string, timestamp time.Time) (int, error) {
	log := LoggerWithContext(c)

	nSend := 0
	for _, offer := range offers {
		if offer.DateCreated.After(timestamp) {
			if err := sendMail(c, offer, address); err != nil {
				log.Errorf("failed sending offer mail: %v", err)
			}
			log.Infof("Successful send email: \n%s", offer)
			nSend++
		}
	}
	return nSend, nil
}

func sendMail(c context.Context, offer Offer, address string) error {

	msg := &mail.Message{
		Sender:  "akinfomer <jobs@asck-158619.appspotmail.com>",
		To:      []string{address},
		Subject: "AK-Informer - New Job Offers (" + offer.Specialization + ")",
		Body:    fmt.Sprintf(confirmMessage, url, offer.Specialization, offer.DateCreated, offer.Address, offer.Description),
	}
	if err := mail.Send(c, msg); err != nil {
		return err
	}
	return nil
}

func subscribeHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented yet", http.StatusInternalServerError)
	return
}

func unsubscribeHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented yet", http.StatusInternalServerError)
	return
}

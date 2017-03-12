package asck

import (
	"net/http"
	"time"

	"fmt"

	"google.golang.org/appengine"
)

// Configuration contains the configuration for the application in one place
type Configuration struct {
	url, email string
}

// Offer represents an offer found online
type Offer struct {
	ID            string
	text, address string
	dateFound     time.Time
}

var config Configuration
var crawler *Crawler
var datastore *Datastore
var sender *Sender

func init() {

	config = Configuration{
		url:   "https://www.aerztekammer-berlin.de/10arzt/15_Weiterbildung/17WB-Stellenboerse/index.html",
		email: "julian.godesa@googlemail.com",
	}

	http.HandleFunc("/update", offersHandler)
	http.Handle("/", http.NotFoundHandler())
}

func offersHandler(w http.ResponseWriter, r *http.Request) {

	ctx := appengine.NewContext(r)

	crawler = &Crawler{ctx: ctx, url: config.url}
	datastore = &Datastore{ctx: ctx}
	sender = &Sender{ctx: ctx, address: config.email}

	offers, err := crawler.Run()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newOffers, err := datastore.Update(offers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := sender.Send(newOffers); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Send '%v' new offers.", len(newOffers))
}

package asck

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
)

// Offer is a individual offer extracted from the job offer page
type Offer struct {
	Text      string
	Contact   string
	StartDate time.Time
}

var offersURL = "https://www.aerztekammer-berlin.de/10arzt/15_Weiterbildung/17WB-Stellenboerse/index.html"

func init() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/offers", offersHandler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}

func offersHandler(w http.ResponseWriter, r *http.Request) {

	offers, err := extractJobOffers(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("extracted a few bits:%v", offers)

	w.WriteHeader(http.StatusNoContent)
}

func extractJobOffers(r *http.Request) ([]Offer, error) {
	ctx := appengine.NewContext(r)
	client := urlfetch.Client(ctx)
	resp, err := client.Get(offersURL)

	if err != nil {
		return nil, errors.New("Failed connecting to the offer page: " + err.Error())
	}
	log.Println("connected.")

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Failed reading body from the response: " + err.Error())
	}
	defer resp.Body.Close()
	log.Println("read data.")

	html := string(body)
	if len(html) != 0 {
		return nil, errors.New("No characters were read from the response: " + err.Error())
	}

	log.Println("read data.")

	return []Offer{}, nil
}

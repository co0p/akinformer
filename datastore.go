package asck

import (
	"fmt"
	"log"

	"golang.org/x/net/context"
)

// Datastore is our interface to the google app engine storage
type Datastore struct {
	ctx context.Context
}

// Update updates the storage and returns any new offers
func (s *Datastore) Update(offers []Offer) ([]Offer, error) {
	log.Printf("Saving potentially '%v' offers", len(offers))

	oldOffers, err := s.getOffers()
	if err != nil {
		return nil, fmt.Errorf("Failed fetching old offers: %v", err)
	}
	log.Printf("Found '%v' older offers", len(oldOffers))

	// remove existing offers
	for _, oldOffer := range oldOffers {
		for i, offer := range offers {
			if offer.compareTo(oldOffer) {
				offers = append(offers[:i], offers[i+1:]...)
			}
		}
	}

	log.Printf("%+v", offers[0])

	return offers, nil
}

func (s *Datastore) getOffers() ([]Offer, error) {
	return []Offer{}, nil
}

package asck

import (
	"log"

	"golang.org/x/net/context"
)

// Datastore is our interface to the google app engine storage
type Datastore struct {
	ctx context.Context
}

// Update updates the storage and returns any new offers
func (s *Datastore) Update(offers []Offer) ([]Offer, error) {
	log.Printf("Updating '%v' offers", len(offers))
	return []Offer{}, nil
}

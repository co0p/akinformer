package asck

import (
	"log"

	"golang.org/x/net/context"
)

// Sender is responsible for sending emails
type Sender struct {
	ctx     context.Context
	address string
}

// Send send the new offers
func (s *Sender) Send(newOffers []Offer) error {
	log.Printf("Sending '%v' to address '%v'", len(newOffers), s.address)
	return nil
}

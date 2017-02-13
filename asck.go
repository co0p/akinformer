package asck

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/user"
)

// const OFFERS_URL "https://www.aerztekammer-berlin.de/10arzt/15_Weiterbildung/17WB-Stellenboerse/index.html"

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	name := "Anonymous"
	if u := user.Current(c); u != nil {
		name = u.String()
	}

	fmt.Fprintf(w, "Hello, world!, %v", name)
}

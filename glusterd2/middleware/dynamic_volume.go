package middleware

import (
	"fmt"
	"net/http"
	"strconv"

	peer "github.com/gluster/glusterd2/glusterd2/peer"
)


// Dynamic_volume is a middleware which generates adds bricks to a volume
// request if it has a key asking for auto brick allocation. It modifies the
// HTTP request and adds bricks to it.
func Dynamic_volume(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if param := r.URL.Query().Get("dynamic_volume"); param != "" {
			peerDetails, err := peer.GetPeers()
			if err != nil {
			}
			for _, peerDetail := range peerDetails {
				fmt.Printf("Printing PEERDETAILS: %s", peerDetail)
				_, err := strconv.Atoi(peerDetail.MetaData["group"])
				if err != nil {
				}
			}
		}
		
		next.ServeHTTP(w, r)
	})
}


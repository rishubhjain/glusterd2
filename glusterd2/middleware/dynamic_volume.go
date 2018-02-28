package middleware

import (
	"fmt"
	"net/http"
)


// Dynamic_volume is a middleware which generates adds bricks to a volume
// request if it has a key asking for auto brick allocation. It modifies the
// HTTP request and adds bricks to it.
func Dynamic_volume(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		param := r.URL.Query().Get("dynamic_volume")
		if param != "" {
		}
		fmt.Printf("HOlaaaaaaa****************8")
		
		next.ServeHTTP(w, r)
	})
}


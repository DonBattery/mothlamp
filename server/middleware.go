package server

import (
	"net/http"

	"github.com/spf13/viper"
)

// authCheck middleware check the HTTP request if "MOTHLAMP_API_TOKEN" can be found
// in the Header, and if its value is equals to the one set by Viper.
func authCheck(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if req.Header.Get("MOTHLAMP_API_TOKEN") != viper.GetString("MOTHLAMP_API_TOKEN") {
			res.WriteHeader(http.StatusUnauthorized)
			res.Write([]byte("Unauthorized\n"))
		} else {
			next.ServeHTTP(res, req)
		}
	})
}

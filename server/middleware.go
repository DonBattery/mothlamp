package server

import (
	"fmt"
	"net/http"

	"github.com/spf13/viper"
)

// authCheck middleware check the HTTP request if "MOTHLAMP_API_TOKEN" can be found
// in the Header, and if its value is equals to the one set by Viper.
func authCheck(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		for key, value := range req.Header {
			fmt.Printf("Incoming HEADER %v : %v\n", key, value)
		}
		if req.Header.Get("apitoken") != viper.GetString("MOTHLAMP_API_TOKEN") {
			fmt.Printf("Token missmatch! Expected : %v Actual : %v\n", viper.GetString("MOTHLAMP_API_TOKEN"), req.Header.Get("apitoken"))
			res.WriteHeader(http.StatusUnauthorized)
			res.Write([]byte("Unauthorized\n"))
		} else {
			next.ServeHTTP(res, req)
		}
	})
}

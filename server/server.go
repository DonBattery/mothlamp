package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-http-utils/logger"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

// Run sets up the router and starts the webserver
func Run() {
	router := mux.NewRouter()

	router.NotFoundHandler = custom404()

	router.Use(authCheck)

	router.HandleFunc("/", welcome).Methods("GET")
	// File upload handler matches for /drive and /drive/ too
	router.HandleFunc("/{drive:drive\\/?}", uploadFile).Methods("POST")
	// This is a very important line, this make the server to actually serve the statis folder
	router.PathPrefix("/drive/").Handler(http.StripPrefix("/drive/", http.FileServer(http.Dir(viper.GetString("MOTHLAMP_DIR"))))).Methods("GET")

	fmt.Printf("\nMothlamp server listening on PORT %v\n", viper.GetInt("MOTHLAMP_PORT"))

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(viper.GetInt("MOTHLAMP_PORT")), logger.Handler(router, os.Stdout, logger.CommonLoggerType)))
}

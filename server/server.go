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
	// Serve the static folder
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(viper.GetString("MOTHLAMP_DIR"))))).Methods("GET")
	// Serve POST requests
	router.HandleFunc("/", uploadFile).Methods("POST")

	fmt.Printf("\nMothlamp server listening on PORT %v\n", viper.GetInt("MOTHLAMP_PORT"))

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(viper.GetInt("MOTHLAMP_PORT")), logger.Handler(router, os.Stdout, logger.CommonLoggerType)))
}

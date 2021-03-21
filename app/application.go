package app

import (
	"log"
	"net/http"
	"os"

	"github.com/Kungfucoding23/bookstore_items_api/client/elasticsearch"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	elasticsearch.Init()
	mapUrls()
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}

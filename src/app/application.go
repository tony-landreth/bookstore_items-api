package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/tony-landreth/bookstore_items-api/src/clients/elasticsearch"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	elasticsearch.Init()
	fmt.Println("API Successfully Started")
	mapUrls()
	srv := &http.Server{
		Handler: router,
		Addr:    ":8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}

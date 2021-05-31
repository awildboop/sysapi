package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/awildboop/sysapi/handlers"
	"github.com/gorilla/mux"
)

func main() {
	config, err := initConfiguration()
	if err != nil {
		log.Fatal("Unable to load configuration.")
	}

	r := mux.NewRouter()

	r.HandleFunc("/api/stats/cpu", handlers.CPUHandler).Methods("GET")
	r.HandleFunc("/api/stats/mem", handlers.MemHandler).Methods("GET")
	r.HandleFunc("/api/stats/sys", handlers.SystemStats).Methods("GET")
	r.HandleFunc("/api/stats/all", handlers.SystemStats).Methods("GET")

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		http.ListenAndServe(fmt.Sprintf("%s:%s", config.Connection.Host, config.Connection.Port), r)
	}()

	fmt.Printf("Now listening at %s on port %s", config.Connection.Host, config.Connection.Port)
	wg.Wait()
}

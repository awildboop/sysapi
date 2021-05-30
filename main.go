package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/awildboop/sysapi/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/stats/cpu", handlers.CPUUsage).Methods("GET")
	r.HandleFunc("/api/stats/mem", handlers.MemUsage).Methods("GET")
	r.HandleFunc("/api/stats/sys", handlers.SystemStats).Methods("GET")
	r.HandleFunc("/api/stats/all", handlers.SystemStats).Methods("GET")

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		http.ListenAndServe(":7171", r)
	}()

	fmt.Println("listening")
	wg.Wait()
}

package cmd

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jayisaac0/auth-service/src/interfaces"
)

// HandleRequets function
func HandleRequets() {

	routes := interfaces.Routing()

	app := &http.Server{
		Addr:              "http://localhost:8080",
		Handler:           routes,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		MaxHeaderBytes:    1 << 20,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 20 * time.Second,
	}

	fmt.Println("Server listening to port 8000")

	if err := app.ListenAndServe(); err != nil {
		log.Fatal(http.ListenAndServe(":8000", routes))
	}
}

package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cruzj6/gosomewhere/random"
	"github.com/cruzj6/gosomewhere/routes"
)

func main() {
	fmt.Println("Starting up GoSomewhere")
	initializeServer()
}

func initializeServer() {
	// Register routes
	random.RegisterRoute()

	// Start server with routes
	r := routes.SetupRoutes()

	srv := &http.Server{
		Handler: r,
		Addr: ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	 }

	srv.ListenAndServe()
}

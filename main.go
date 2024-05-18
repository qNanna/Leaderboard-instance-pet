package main

import (
	// INTERNAL IMPORTS
	"fmt"
	"net/http"

	// PACKAGES IMPORT
	config "main/app/common"
	database "main/app/database"
	router "main/app/src"

	// EXTERNAL IMPORTS
	"github.com/gofor-little/env"
)

func main() {
	// INIT CONFIG/DATABASE/ROUTER
	config.Init()
	database.Init()
	mux := router.Init()

	// GET PARAMS FROM ENV
	port := env.Get("PORT", ":3000")

	// LOG BEFORE START
	fmt.Printf("Application starting at port: %s\n", port)

	// START API
	if err := http.ListenAndServe(port, mux); err != nil {
		fmt.Println("Failed to start server")
	}
}

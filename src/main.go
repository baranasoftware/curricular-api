package main

import (
	"curricular-api/api"
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := 8080

	fmt.Println("Curricular API server is listing on port:", port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), api.Server)
	if err != nil {
		fmt.Println("error starting the sever", err)
		os.Exit(1)
	}
}

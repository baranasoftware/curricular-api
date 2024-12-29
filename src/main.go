package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := 8080

	fmt.Println("Curricular API server is listing on port:", port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), Server)
	if err != nil {
		fmt.Println("error starting the sever", err)
		os.Exit(1)
	}
}

package main

import (
	"curricular-api/api"
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {

	// local deployment will set up OAuth2/JWT setup for demonstration,
	// AWS/Lambda deployment will use AWS APIGateway for API authorization.

	local := flag.Bool("local", false, "local deployment")
	port := flag.Int("port", 8080, "port")

	flag.Parse()

	if *local {
		fmt.Println("Curricular API server is listing on port:", *port)

		err := http.ListenAndServe(fmt.Sprintf(":%d", *port), api.Server)
		if err != nil {
			fmt.Println("error starting the sever", err)
			os.Exit(1)
		}
	} else {
		// AWS API Gateway
	}
}

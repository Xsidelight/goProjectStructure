package main

import (
	"fmt"
	"net/http"

	"github.com/Xsidelight/goProjectStructure/internal/routes"
)

func main() {
	router := routes.NewRouter()

	port := 8080
	addrs := fmt.Sprintf("%d", port)

	fmt.Printf("Listening on port %s\n", addrs)
	err := http.ListenAndServe(":"+addrs, router)
	if err != nil {
		panic(err)
	}

}

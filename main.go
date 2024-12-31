package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jqwez/wording/server"
)

func main() {
	s := server.NewServer()
	fmt.Println("Listening on port :8080")
	log.Fatal(http.ListenAndServe(":8080", s.Mux))
}

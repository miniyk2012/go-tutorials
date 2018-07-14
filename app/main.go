package main

import (
	"fmt"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"net/http"
	"log"
)

func main() {
	fmt.Println("Hello 世界！I‘m fucking using Docker！", errors.New("xxx"))

	router := mux.NewRouter()

	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(200)
		writer.Write([]byte("Hello 世界"))
	})

	log.Fatal(http.ListenAndServe(":8000", router))
}

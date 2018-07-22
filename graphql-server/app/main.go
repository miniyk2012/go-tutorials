package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"io/ioutil"
	g "github.com/CreatCodeBuild/go-tutorials/graphql-server/app/graphql"
)



func main() {
	router := mux.NewRouter()


	// read graphql schema
	b, err := ioutil.ReadFile("app/graphql/schema.graphql")
	if err != nil {
		panic(err)
	}

	schema := graphql.MustParseSchema(string(b), &g.Query{})

	router.Use(func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Do stuff here
			log.Println(r.RequestURI)
			// Call the next handler, which can be another middleware in the chain, or the final handler.
			w.Header().Set( "Access-Control-Allow-Origin", "*")
			handler.ServeHTTP(w, r)
		})
	})


	router.Handle("/query", &relay.Handler{Schema: schema})

	fs := http.FileServer(http.Dir("public"))
	router.Handle("/", http.StripPrefix("/public/", fs))


	// todo: replace with env var / configs
	// todo: consider Viper
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", router))
}

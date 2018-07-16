package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"io/ioutil"
)

type query struct{}


func (_ *query) Chart1() *Chart1 { return &Chart1{} }

type Chart1 struct {}

func (m Chart1) Points() []*Point {
	return []*Point{
		{a: "C", b: 2},
		{a: "C", b: 7},
		{a: "C", b: 4},
		{a: "D", b: 1},
		{a: "D", b: 3},
	}
}

type Point struct {
	a string
	b int32
}

func (p *Point) A() string {return p.a}
func (p *Point) B() int32 {return p.b}

func main() {
	router := mux.NewRouter()


	// read graphql schema

	b, err := ioutil.ReadFile("app/schema.graphql")
	if err != nil {
		panic(err)
	}

	schema := graphql.MustParseSchema(string(b), &query{})

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

package dgraph

import (
	"log"
	"google.golang.org/grpc"
	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
	"context"
	)

func F() error {
	conn, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	dgraphClient := dgo.NewDgraphClient(api.NewDgraphClient(conn))

	op := &api.Operation{
		Schema: `name: string @index(exact) .`,
	}
	err = dgraphClient.Alter(context.Background(), op)
	return err
}
//router := mux.NewRouter()
//
//
//// read graphql schema
//b, err := ioutil.ReadFile("app/graphql/schema.graphql")
//if err != nil {
//panic(err)
//}
//
//schema := graphql.MustParseSchema(string(b), &g.Query{})
//
//router.Use(func(handler http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		// Do stuff here
//		log.Println(r.RequestURI)
//		// Call the next handler, which can be another middleware in the chain, or the final handler.
//		w.Header().Set( "Access-Control-Allow-Origin", "*")
//		handler.ServeHTTP(w, r)
//	})
//})
//
//
//router.Handle("/query", &relay.Handler{Schema: schema})
//
//fs := http.FileServer(http.Dir("public"))
//router.Handle("/", http.StripPrefix("/public/", fs))
//
//
//// todo: replace with env var / configs
//// todo: consider Viper
//log.Fatal(http.ListenAndServe("0.0.0.0:8000", router))
//}

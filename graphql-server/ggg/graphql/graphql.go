package graphql

import (
	"log"
	"net/http"

	graph "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"io/ioutil"
	"github.com/CreatCodeBuild/go-tutorials/graphql-server/core/usecases"
	"github.com/CreatCodeBuild/go-tutorials/graphql-server/core/adapters/testadapters"
	"github.com/CreatCodeBuild/go-tutorials/graphql-server/core/entities"
	"github.com/CreatCodeBuild/go-tutorials/graphql-server/core/adapters"
	"github.com/CreatCodeBuild/go-tutorials/graphql-server/ggg/db/badger"
	"fmt"
)

type query struct {
	userAdapter adapters.User
}

var adapter = testadapters.Adapters{}

func (q *query) GetMovies() ([]Movie, error) {

	movies, err := usecases.GetAllMovies(adapter.GetAllMovies)

	// map entity movie -> graphql movie
	gMovies := make([]Movie, len(movies))
	for i, movie := range movies {
		gMovies[i] = Movie{name: movie.Title}
	}

	return gMovies, err
}

type addMoviesArg struct {
	Title string
	UID   string
}

func (q *query) AddMovies(args addMoviesArg) ([]Movie, error) {
	err := usecases.AddMovie(args.UID, entities.Movie{Title: args.Title}, adapter.AddMovieToDB, q.userAdapter)
	if err != nil {
		return nil, err
	}

	return q.GetMovies()
}

type signUpArg struct {
	UID               string
	AuthorizedActions []string
}

func (q *query) SignUp(args signUpArg) (User, error) {
	user, err := usecases.UserSignUp(
		entities.User{
			UID:               args.UID,
			AuthorizedActions: args.AuthorizedActions,
		},
		q.userAdapter)
	if err != nil {
		fmt.Printf("%+v\n", err)
		return User{}, err
	}
	return User{
		uid:               user.UID,
		authorizedActions: user.AuthorizedActions,
	}, nil
}

type Movie struct {
	name string
}

func (m Movie) Name() string {
	return m.name
}

type User struct {
	uid               string
	authorizedActions []string
}

func (u User) UID() string {
	return u.uid
}

func (u User) AuthorizedActions() []string {
	return u.authorizedActions
}

////////////////// below is graphql specific logic ////////////////////////

func Init() {

	s, err := ioutil.ReadFile("./ggg/graphql/schema.graphql")
	if err != nil {
		panic(err)
	}

	badger, err := badger.NewBadger()
	if err != nil {
		panic(err)
	}

	schema := graph.MustParseSchema(
		string(s),
		&query{
			userAdapter: badger,
		},
	)

	http.Handle("/query", &relay.Handler{Schema: schema})
	serveGraphiQL()
	log.Fatal(http.ListenAndServe(":8888", nil))
}

func serveGraphiQL() {
	var page = []byte(`
<!DOCTYPE html>
<html>
	<head>
		<link href="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.11.11/graphiql.min.css" rel="stylesheet" />
		<script src="https://cdnjs.cloudflare.com/ajax/libs/es6-promise/4.1.1/es6-promise.auto.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/fetch/2.0.3/fetch.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/react/16.2.0/umd/react.production.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/react-dom/16.2.0/umd/react-dom.production.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.11.11/graphiql.min.js"></script>
	</head>
	<body style="width: 100%; height: 100%; margin: 0; overflow: hidden;">
		<div id="graphiql" style="height: 100vh;">Loading...</div>
		<script>
			function graphQLFetcher(graphQLParams) {
				return fetch("/query", {
					method: "post",
					body: JSON.stringify(graphQLParams),
					credentials: "include",
				}).then(function (response) {
					return response.text();
				}).then(function (responseBody) {
					try {
						return JSON.parse(responseBody);
					} catch (error) {
						return responseBody;
					}
				});
			}
			ReactDOM.render(
				React.createElement(GraphiQL, {fetcher: graphQLFetcher}),
				document.getElementById("graphiql")
			);
		</script>
	</body>
</html>
`)

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(page)
	}))
}

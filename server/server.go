package main

import (
	"log"
	"net/http"

	"bitbucket.org/ffxblue/golang-demo"
	"github.com/neelance/graphql-go"
	"github.com/neelance/graphql-go/relay"
	"google.golang.org/grpc"

	pb "bitbucket.org/ffxblue/golang-demo/generated/grpc"
	pbs "bitbucket.org/ffxblue/golang-demo/generated/spark"
)

var schema *graphql.Schema

const (
	gaAddress    = "localhost:50051"
	sparkAddress = "localhost:50052"
)

func init() {}

func main() {

	// grpc connection to analytics 1
	// Set up a connection to the server to the GA service
	gconn, err := grpc.Dial(gaAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer gconn.Close()
	gc := pb.NewGoogleAnalyticsClient(gconn)

	// grpc connection to analytics 2
	// Set up a connection to the server to the spark service
	sconn, err := grpc.Dial(sparkAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer sconn.Close()
	sc := pbs.NewSparkAnalyticsClient(sconn)

	// graphql
	schema, err = graphql.ParseSchema(demo.Schema, &demo.Resolver{GoogleAnalyticsClient: gc, SparkAnalyticsClient: sc})
	if err != nil {
		panic(err)
	}

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(page)
	}))

	http.Handle("/query", &relay.Handler{Schema: schema})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

var page = []byte(`
<!DOCTYPE html>
<html>
	<head>
		<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.7.8/graphiql.css" />
		<script src="https://cdnjs.cloudflare.com/ajax/libs/fetch/1.0.0/fetch.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/react/15.3.2/react.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/react/15.3.2/react-dom.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.7.8/graphiql.js"></script>
	</head>
	<body style="width: 100%; height: 100%; margin: 0; overflow: hidden;">
		<div id="graphiql" style="height: 100vh;">Loading...</div>
		<script>
			function graphQLFetcher(graphQLParams) {
				graphQLParams.variables = graphQLParams.variables ? JSON.parse(graphQLParams.variables) : null;
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

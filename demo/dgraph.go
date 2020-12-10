package demo

import (
	"flag"
	"log"
)

var (
	dgraph = flag.String("d", "localhost:9080", "Dgraph Alpha address")
)

func buildDgraph() {
	flag.Parse()
	conn, err := grpc.Dial(*dgraph, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := dgo.NewDgraphClient(api.NewDgraphClient(conn))

}

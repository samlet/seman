package graph_procs

import "testing"
import (
	"context"
	"flag"
	"fmt"
	"log"
	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
	"google.golang.org/grpc"
)

//⊕ [查询语言 - Dgraph Doc v1.0.14](https://docs.dgraph.io/query-language/)
var (
	dgraph = flag.String("d", "127.0.0.1:9080", "Dgraph Alpha address")
)

//+pre sagas/tests/dgraph/simple.py
func TestQuery(test *testing.T) {
	conn, err := grpc.Dial(*dgraph, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()


	dg := dgo.NewDgraphClient(api.NewDgraphClient(conn))


	/*
	resp, err := dg.NewTxn().Query(context.Background(), `{
  bladerunner(func: eq(name@en, "Blade Runner")) {
    uid
    name@en
    initial_release_date
    netflix_id
  }
}`)
	*/
	resp, err := dg.NewTxn().QueryWithVars(context.Background(), `query all($a: string) {
        all(func: eq(name, $a)) {
            uid
            name
            age
            married
            loc
            dob
            friend {
                name
                age
            }
            school {
                name
            }
        }
    }`, map[string]string{"$a": "Alice"})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Response: %s\n", resp.Json)
}


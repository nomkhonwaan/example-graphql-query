package graphql

import (
	"bufio"
	"encoding/csv"
	"log"
	"net/http"
	"os"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

var people People
var schema graphql.Schema
var err error

// init will read entire MOCK_DATA.csv and store total 1,000 records on people var
func init() {
	f, _ := os.Open("MOCK_DATA.csv")
	r := csv.NewReader(bufio.NewReader(f))

	records, _ := r.ReadAll()

	for i, record := range records {
		// Skip on header
		if i == 0 {
			continue
		}

		people = append(people, Person{
			ID:        record[0],
			FirstName: record[1],
			LastName:  record[2],
			Email:     record[3],
			Gender:    record[4],
		})
	}
}

// Serve GraphQL server on port 8080
func Serve() {
	schema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})
	if err != nil {
		log.Fatalf("An error has occured while pasrsing GrahpQL schema: %s", err)
	}

	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})

	http.Handle("/graphql", h)

	log.Println("Listening on port 8080")

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("An error has occured while listening on port 8080: %s", err)
	}
}

package main

import (
	"bufio"
	"encoding/csv"
	"log"
	"net/http"
	"os"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

type Person struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
}

type People []Person

func (p People) Filter(f func(Person) bool) People {
	var result People
	for _, person := range p {
		if f(person) {
			result = append(result, person)
		}
	}
	return result
}

var personType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Person",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.ID,
		},
		"firstName": &graphql.Field{
			Type: graphql.String,
		},
		"lastName": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"gender": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"people": &graphql.Field{
			Type: graphql.NewList(personType),
			Args: graphql.FieldConfigArgument{
				"gender": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				result := people

				if gender, ok := p.Args["gender"].(string); ok {
					result = people.Filter(func(elem Person) bool {
						return elem.Gender == gender
					})
				}

				return result, nil
			},
		},
		"person": &graphql.Field{
			Type: personType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if id, ok := p.Args["id"].(int); ok {
					return people[id-1], nil
				}
				return nil, nil
			},
		},
	},
})

var people People
var schema graphql.Schema
var err error

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

func main() {
	var schema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})
	if err != nil {
		log.Fatalf("An error has occured while parsing Schema config: %s", err)
	}

	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})
	http.Handle("/graphql", h)

	log.Println("Listening on port 8080")

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("An error has occured while starting HTTP server on port 8080: %s.", err)
	}
}

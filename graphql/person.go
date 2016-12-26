package graphql

import (
	"github.com/graphql-go/graphql"
)

// Person model
type Person struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
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

var personQuery = &graphql.Field{
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
}

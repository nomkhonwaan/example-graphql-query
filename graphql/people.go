package graphql

import (
	"github.com/graphql-go/graphql"
)

// People model, this model is a slice of Person model
type People []Person

func (people People) filter(fn func(Person) bool) People {
	var result People
	for _, person := range people {
		if fn(person) {
			result = append(result, person)
		}
	}

	return result
}

var peopleQuery = &graphql.Field{
	Type: graphql.NewList(personType),
	Args: graphql.FieldConfigArgument{
		"gender": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		if gender, ok := p.Args["gender"].(string); ok {
			return people.filter(func(person Person) bool {
				return person.Gender == gender
			}), nil
		}
		return people, nil
	},
}

package graph

import "github.com/Nikhil12894/gqlwithgo/graph/model"

//go:generate go run github.com/99designs/gqlgen
type Resolver struct {
	Vachils []*model.Vachil
}

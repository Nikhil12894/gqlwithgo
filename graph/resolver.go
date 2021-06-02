package graph

import "gorm.io/gorm"

//go:generate go run github.com/99designs/gqlgen
type Resolver struct {
	DB *gorm.DB
}

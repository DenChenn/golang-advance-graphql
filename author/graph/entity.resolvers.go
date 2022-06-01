package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/SwarzChen/golang-advance-graphql/author/service_error"

	"github.com/SwarzChen/golang-advance-graphql/author/graph/generated"
	"github.com/SwarzChen/golang-advance-graphql/author/graph/model"
)

func (r *entityResolver) FindAuthorByID(ctx context.Context, id string) (*model.Author, error) {
	for _, author := range AuthorData {
		if author.ID == id {
			return author, nil
		}
	}
	return nil, service_error.NoAuthorFoundError
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }

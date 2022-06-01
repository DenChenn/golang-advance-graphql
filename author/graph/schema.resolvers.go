package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/SwarzChen/golang-advance-graphql/author/service_error"

	"github.com/SwarzChen/golang-advance-graphql/author/graph/generated"
	"github.com/SwarzChen/golang-advance-graphql/author/graph/model"
)

func (r *queryResolver) Author(ctx context.Context, authorID string) (*model.Author, error) {
	for _, author := range AuthorData {
		if authorID == author.ID {
			return author, nil
		}
	}
	return nil, service_error.NoAuthorFoundError
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

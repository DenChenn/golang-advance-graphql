package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/SwarzChen/golang-advance-graphql/diary/service_error"

	"github.com/SwarzChen/golang-advance-graphql/diary/graph/generated"
	"github.com/SwarzChen/golang-advance-graphql/diary/graph/model"
)

func (r *queryResolver) Diary(ctx context.Context, diaryID string) (*model.Diary, error) {
	for _, diary := range DiaryData {
		if diary.ID == diaryID {
			return diary, nil
		}
	}
	return nil, service_error.NoDiaryFoundError
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

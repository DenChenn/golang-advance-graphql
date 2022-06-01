package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/SwarzChen/golang-advance-graphql/diary/graph/generated"
	"github.com/SwarzChen/golang-advance-graphql/diary/graph/model"
	"github.com/SwarzChen/golang-advance-graphql/diary/service_error"
)

func (r *entityResolver) FindAuthorByID(ctx context.Context, id string) (*model.Author, error) {
	// This resolve the extended version of author, that is, only deal with "diaries" and "number_of_diaries" field.
	var authorDiaryList []*model.Diary

	for _, diary := range DiaryData {
		if diary.Author.ID == id {
			authorDiaryList = append(authorDiaryList, diary)
		}
	}

	return &model.Author{
		ID:              id,
		Diaries:         authorDiaryList,
		NumberOfDiaries: len(authorDiaryList),
	}, nil
}

func (r *entityResolver) FindDiaryByID(ctx context.Context, id string) (*model.Diary, error) {
	// Inside this diary subgraph server, we can only resolve author's name inside diary entity.
	for _, diary := range DiaryData {
		if id == diary.ID {
			return diary, nil
		}
	}
	return nil, service_error.NoDiaryFoundError
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }

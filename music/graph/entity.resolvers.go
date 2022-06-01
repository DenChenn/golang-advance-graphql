package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/SwarzChen/golang-advance-graphql/music/graph/generated"
	"github.com/SwarzChen/golang-advance-graphql/music/graph/model"
)

func (r *entityResolver) FindDiaryByID(ctx context.Context, id string) (*model.Diary, error) {
	// We deal with extend version of diary here, that is, we only return "musics" field.
	var musicList []*model.Music
	for _, link := range DiaryMusicLinkData {
		if link.DiaryID == id {
			for _, targetMusic := range MusicData {
				if link.MusicID == targetMusic.ID {
					musicList = append(musicList, targetMusic)
				}
			}
		}
	}

	return &model.Diary{
		ID:     id,
		Musics: musicList,
	}, nil
}

func (r *entityResolver) FindMusicByID(ctx context.Context, id string) (*model.Music, error) {
	panic(fmt.Errorf("not implemented"))
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }

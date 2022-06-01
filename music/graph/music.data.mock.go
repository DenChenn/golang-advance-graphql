package graph

import "github.com/SwarzChen/golang-advance-graphql/music/graph/model"

var MusicData = []*model.Music{
	{
		ID: "m1",
		Author: &model.Author{
			ID: "a1",
		},
		Name:   "TWICE K-POP",
		Length: "4 min",
	},
	{
		ID: "m2",
		Author: &model.Author{
			ID: "a1",
		},
		Name:   "DJSnake ROCK",
		Length: "3 min",
	},
	{
		ID: "m3",
		Author: &model.Author{
			ID: "a2",
		},
		Name:   "Classic music",
		Length: "10 min",
	},
	{
		ID: "m4",
		Author: &model.Author{
			ID: "a2",
		},
		Name:   "Nice JAZZ",
		Length: "9 min",
	},
}

var DiaryMusicLinkData = []*model.DiaryMusicLink{
	{
		DiaryID: "d1",
		MusicID: "m1",
	},
	{
		DiaryID: "d1",
		MusicID: "m2",
	},
	{
		DiaryID: "d2",
		MusicID: "m1",
	},
	{
		DiaryID: "d2",
		MusicID: "m2",
	},
	{
		DiaryID: "d3",
		MusicID: "m3",
	},
	{
		DiaryID: "d3",
		MusicID: "m4",
	},
	{
		DiaryID: "d4",
		MusicID: "m3",
	},
	{
		DiaryID: "d4",
		MusicID: "m4",
	},
}

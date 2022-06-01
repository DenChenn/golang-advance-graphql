package graph

import "github.com/SwarzChen/golang-advance-graphql/diary/graph/model"

var DiaryData = []*model.Diary{
	{
		ID: "d1",
		Author: &model.Author{
			ID:   "a1",
			Name: "Crazy Coder",
		},
		Content: "Nice sunny day!",
	},
	{
		ID: "d2",
		Author: &model.Author{
			ID:   "a1",
			Name: "Crazy Coder",
		},
		Content: "I hate rainy days.",
	},
	{
		ID: "d3",
		Author: &model.Author{
			ID:   "a2",
			Name: "Realistic Coder",
		},
		Content: "Nice rainy day",
	},
	{
		ID: "d4",
		Author: &model.Author{
			ID:   "a2",
			Name: "Realistic Coder",
		},
		Content: "I hate sunny days",
	},
}

package graph

import "github.com/SwarzChen/golang-advance-graphql/author/graph/model"

var crazyCodeBadgeList = []*model.Badge{
	{
		Description: "Big k-pop fan.",
		Color:       "Blue",
	},
	{
		Description: "Big rock fan.",
		Color:       "Red",
	},
}

var seniorCoderBadgeList = []*model.Badge{
	{
		Description: "Big jazz fan.",
		Color:       "Blue",
	},
	{
		Description: "Big classic fan",
		Color:       "Green",
	},
}

var AuthorData = []*model.Author{
	{
		ID:     "a1",
		Name:   "Crazy Coder",
		Email:  "crazycoder@gmail.com",
		Badges: crazyCodeBadgeList,
	},
	{
		ID:     "a2",
		Name:   "Realistic Coder",
		Email:  "realisticcoder@gmail.com",
		Badges: seniorCoderBadgeList,
	},
}

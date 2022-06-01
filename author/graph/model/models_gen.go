// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Author struct {
	ID     string   `json:"id"`
	Name   string   `json:"name"`
	Email  string   `json:"email"`
	Badges []*Badge `json:"badges"`
}

func (Author) IsEntity() {}

type Badge struct {
	Description string `json:"description"`
	Color       string `json:"color"`
}
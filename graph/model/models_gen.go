// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Character struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Likes int    `json:"likes"`
}

type CharacterInput struct {
	Name  string `json:"name"`
	Likes int    `json:"likes"`
}

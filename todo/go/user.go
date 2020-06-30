package todo

import (
	// "fmt"

	"github.com/graph-gophers/graphql-go"
	// "github.com/graph-gophers/graphql-go/relay"
)

var users = []*user{
	{
		ID:            "2000",
		UserID:         "me",
		TotalCount:     10,
		CompletedCount: 90,
	},
}

var userData = make(map[graphql.ID]*user)

func init() {
	for _, u := range users {
		userData[u.UserID] = u
	}
}

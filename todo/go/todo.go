package todo

import (
	// "fmt"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

type todoConnection struct {
	PageInfo pageInfo
	Edges    *[]*todoEdge
}

type todoEdge struct {
	Node   todo
	Cursor string
}

type pageInfo struct {
	HasNextPage     bool
	HasPreviousPage bool
	StartCursor     string
	EndCursor       string
}

type todo struct {
	ID       graphql.ID
	Text     string
	Complete bool
}

type node interface {
	ID() graphql.ID
}

type user struct {
	ID             graphql.ID
	UserID         graphql.ID
	Todos          todoConnection
	TotalCount     int32
	CompletedCount int32
}

type Resolver struct{}

type userResolver struct {
	u *user
}

type nodeResolver struct {
	node
}

func (r *Resolver) User(args struct{ UserID graphql.ID }) *userResolver {
	if u := userData[args.UserID]; u != nil {
		return &userResolver{u}
	}
	return nil
}

func (r *userResolver) ID() graphql.ID {
	return relay.MarshalID("User", r.u.ID)
}

func (r *userResolver) UserID() graphql.ID {
	return r.u.UserID
}

func (r *userResolver) TotalCount() int32 {
	return r.u.TotalCount
}

func (r *userResolver) CompletedCount() int32 {
	return r.u.CompletedCount
}

func (r *nodeResolver) ToUser() (*nodeResolver, bool) {
	c, ok := r.node.(*nodeResolver)
	return c, ok
}

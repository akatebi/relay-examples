package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/akatebi/relay-examples/todo"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

var schema *graphql.Schema

func init() {
	data, error := ioutil.ReadFile("./schema.graphql")
	if error != nil {
		panic(error)
	}
	var Schema = string(data)
	schema = graphql.MustParseSchema(Schema, &todo.Resolver{})
}

func main() {
	var crap = http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Write(page)
		})
	http.Handle("/crap", crap)
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)
	http.Handle("/query", &relay.Handler{Schema: schema})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

var page = []byte(`<!DOCTYPE html><body>Crap</body><html></html>`)

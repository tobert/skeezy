package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"tux21b.org/v1/gocql"
)

var cass *gocql.Session

func main() {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "skeezy"
	cluster.Consistency = gocql.Quorum

	// using := here will mask the global 'cass' causing pointer exceptions
	var err error
	cass, err = cluster.CreateSession()
	if err != nil {
		panic(fmt.Sprintf("Error creating Cassandra session: %v", err))
	}
	defer cass.Close()

	r := mux.NewRouter()

	r.HandleFunc("/", getRoot)
	r.HandleFunc("/post/", listPosts).Methods("GET")
	r.HandleFunc("/post/{id}", getPost).Methods("GET")
	r.HandleFunc("/post/{id}", newPost).Methods("PUT")
	r.HandleFunc("/post/{id}", delPost).Methods("DELETE")

	r.HandleFunc("/comment/", listComments).Methods("GET")
	r.HandleFunc("/comment/{id}", getComment).Methods("GET")
	r.HandleFunc("/comment/{id}", newComment).Methods("PUT")
	r.HandleFunc("/comment/{id}", delComment).Methods("DELETE")

	http.ListenAndServe(":8080", r)
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("unimplemented route: getRoot()\n")
}

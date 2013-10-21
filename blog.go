package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"tux21b.org/v1/gocql"
)

//"tux21b.org/v1/gocql/uuid"

func main() {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "skeezy"
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()
	if err != nil {
		panic("Error creating Cassandra session.")
	}
	defer session.Close()

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
}

func getRoot(w http.ResponseWriter, r *http.Request) {
}

func getPosts(w http.ResponseWriter, r *http.Request) {
}

func getPost(w http.ResponseWriter, r *http.Request) {
}

func listPosts(w http.ResponseWriter, r *http.Request) {
}

func newPost(w http.ResponseWriter, r *http.Request) {
}

func delPost(w http.ResponseWriter, r *http.Request) {
}

func getComments(w http.ResponseWriter, r *http.Request) {
}

func getComment(w http.ResponseWriter, r *http.Request) {
}

func listComments(w http.ResponseWriter, r *http.Request) {
}

func newComment(w http.ResponseWriter, r *http.Request) {
}

func delComment(w http.ResponseWriter, r *http.Request) {
}

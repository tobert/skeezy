package main

import (
	"fmt"
	"net/http"
	"time"
	"tux21b.org/v1/gocql/uuid"
)

type comment struct {
	id      uuid.UUID
	post_id uuid.UUID
	dttm    time.Time
	author  string
	email   string
	body    string
}

func getComments(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("unimplemented route: getComments()\n")
}

func newComment(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("unimplemented route: newComment()\n")
}

func getComment(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("unimplemented route: getComment()\n")
}

func listComments(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("unimplemented route: listComments()\n")
}

func delComment(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("unimplemented route: delComment()\n")
}

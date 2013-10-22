package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"tux21b.org/v1/gocql/uuid"
)

type Post struct {
	Id      uuid.UUID `json:id`
	Body    string    `json:body`
	Dttm    time.Time `json:dttm`
	Authors []string  `json:authors`
	Tags    []string  `json:tags`
}

func getPost(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("unimplemented route: getPost()\n")
}

func newPost(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("unimplemented route: newPost()\n")
}

func listPosts(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("called listPosts\n")
	q := cass.Query(`SELECT id, body, dttm, authors, tags FROM posts`)
	p := Post{}
	q.Scan(&p.Id, &p.Body, &p.Dttm, &p.Authors, &p.Tags)
	json, _ := json.Marshal(p)
	w.Write(json)
}

func delPost(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("unimplemented route: delPost()\n")
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("unimplemented route: getPosts()\n")
}

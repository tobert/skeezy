package main

import (
	"encoding/json"
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
	q := cass.Query(`SELECT id, body, dttm, authors, tags FROM posts WHERE id=?`)
	p := Post{}
	q.Scan(&p.Id, &p.Body, &p.Dttm, &p.Authors, &p.Tags)
	json, _ := json.Marshal(p)
	w.Write(json)
}

func newPost(w http.ResponseWriter, r *http.Request) {
	// get JSON, parse it, write to C*
	p := Post{}
	q := cass.Query(`INSERT INTO posts (id, body, dttm, authors, tags) VALUES (?, ?, ?, ?, ?)`)
	q.Scan(&p.Id, &p.Body, &p.Dttm, &p.Authors, &p.Tags)
	json, _ := json.Marshal(p)
	w.Write(json)
}

func listPosts(w http.ResponseWriter, r *http.Request) {
	q := cass.Query(`SELECT id, body, dttm, authors, tags FROM posts`)
	// TODO: loop & build a list to marshal to json
	p := Post{}
	q.Scan(&p.Id, &p.Body, &p.Dttm, &p.Authors, &p.Tags)
	json, _ := json.Marshal(p)
	w.Write(json)
}

func delPost(w http.ResponseWriter, r *http.Request) {
	q := cass.Query(`DELETE FROM posts WHERE id=?`)
}

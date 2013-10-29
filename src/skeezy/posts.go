package skeezy

import (
	"encoding/json"
	"net/http"
	"tux21b.org/v1/gocql"
)

func ListPosts(cass *gocql.Session, w http.ResponseWriter, r *http.Request) {
	q := cass.Query(`SELECT id, body, created, authors, tags FROM posts`)
	// TODO: loop & build a list to marshal to json
	p := Post{}
	q.Scan(&p.Id, &p.Body, &p.Created, &p.Authors, &p.Tags)
	json, _ := json.Marshal(p)
	w.Write(json)
}

func GetPost(cass *gocql.Session, id string, w http.ResponseWriter, r *http.Request) {
	q := cass.Query(`SELECT id, body, created, authors, tags FROM posts WHERE id=?`, id)
	p := Post{}
	q.Scan(&p.Id, &p.Body, &p.Created, &p.Authors, &p.Tags)
	json, _ := json.Marshal(p)
	w.Write(json)
}

func NewPost(cass *gocql.Session, id string, w http.ResponseWriter, r *http.Request) {
	// get JSON, parse it, write to C*
	p := Post{}
	q := cass.Query(`INSERT INTO posts (id, body, created, authors, tags) VALUES (?, ?, ?, ?, ?)`)
	q.Scan(&p.Id, &p.Body, &p.Created, &p.Authors, &p.Tags)
	json, _ := json.Marshal(p)
	w.Write(json)
}

func DelPost(cass *gocql.Session, id string, w http.ResponseWriter, r *http.Request) {
	//q := cass.Query(`DELETE FROM posts WHERE id=?`)
}

package skeezy

import (
	"encoding/json"
	"github.com/gocql/gocql"
	"log"
	"net/http"
)

func ListPosts(cass *gocql.Session) PostList {
	plist := make(PostList, 0)

	iq := cass.Query(`SELECT id, body, created, authors, tags FROM posts`).Iter()
	for {
		p := Post{}
		if iq.Scan(&p.Id, &p.Body, &p.Created, &p.Authors, &p.Tags) {
			plist = append(plist, p)
		} else {
			break
		}
	}
	if err := iq.Close(); err != nil {
		log.Fatal(err)
	}

	return plist
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

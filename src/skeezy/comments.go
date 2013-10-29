package skeezy

import (
	"fmt"
	"log"
	"net/http"
	"tux21b.org/v1/gocql"
	"tux21b.org/v1/gocql/uuid"
)

func ListComments(cass *gocql.Session, id uuid.UUID) chan *Comment {
	cc := make(chan *Comment)

	go func() {
		iq := cass.Query(`SELECT id, parentId, created FROM comments WHERE postId=?`, id.Bytes()).Iter()
		for {
			c := Comment{}
			if iq.Scan(&c.Id, &c.ParentId, &c.Created) {
				cc <- &c
			} else {
				break
			}
		}
		if err := iq.Close(); err != nil {
			log.Fatal(err)
		}
		close(cc)
	}()

	return cc
}

func GetComment(cass *gocql.Session, id string, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{}\n")
}

func NewComment(cass *gocql.Session, id string, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{}\n")
}

func DelComment(cass *gocql.Session, id string, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{}\n")
}

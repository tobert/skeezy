package skeezy

import (
	"fmt"
	"log"
	"net/http"
	"tux21b.org/v1/gocql"
	"tux21b.org/v1/gocql/uuid"
)

func ListComments(cass *gocql.Session, id uuid.UUID) (chan Comment) {
	cc := make(chan Comment, 100)
	fmt.Printf("A comment list was requested ...\n")

	go func() {
		fmt.Printf("Querying database ....\n")
		iq := cass.Query(`SELECT id, parentId, created FROM comments WHERE postId=?`, id).Iter()
		for iq.Scan() {
			c := Comment{}
			iq.Scan(&c.Id, &c.ParentId, &c.Created)
			fmt.Printf("Comment: %v\n", c)
			cc <- c
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

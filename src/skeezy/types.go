package skeezy

import (
	"time"
	"github.com/gocql/gocql"
)

type User struct {
	Id       gocql.UUID `json:id`
	Username string    `json:username`
	Email    string    `json:email`
	Created  time.Time `json:created`
	Updated  gocql.UUID `json:updated`
}

type Post struct {
	Id      gocql.UUID `json:id`
	Body    string    `json:body`
	Created time.Time `json:created`
	Authors []string  `json:authors`
	Tags    []string  `json:tags`
	UserID  gocql.UUID `json:user_id`
}

type Comment struct {
	Id       gocql.UUID `json:id`
	ParentId gocql.UUID `json:parent_id`
	PostId   gocql.UUID `json:post_id`
	Created  time.Time `json:created`
	Author   string    `json:author`
	Email    string    `json:email`
	Body     string    `json:body`
}

package skeezy

import (
	"time"
	"tux21b.org/v1/gocql/uuid"
)

type User struct {
	Id       uuid.UUID `json:id`
	Username string    `json:username`
	Email    string    `json:email`
	Created  time.Time `json:created`
}

type Post struct {
	Id      uuid.UUID `json:id`
	Body    string    `json:body`
	Created time.Time `json:created`
	Authors []string  `json:authors`
	Tags    []string  `json:tags`
	UserID  uuid.UUID `json:user_id`
}

type Comment struct {
	Id       uuid.UUID `json:id`
	ParentId uuid.UUID `json:parent_id`
	PostId   uuid.UUID `json:post_id`
	Created  time.Time `json:created`
	Author   string    `json:author`
	Email    string    `json:email`
	Body     string    `json:body`
}

/*
func (c *Comment) MarshalJSON() ([]byte, error) {
	return []byte(`"T"`), nil
}
*/

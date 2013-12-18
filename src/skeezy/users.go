package skeezy

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"tux21b.org/v1/gocql"
	"tux21b.org/v1/gocql/uuid"
)

func ListUsers(cass *gocql.Session) []User {
	ulist := make([]User, 0)

	iq := cass.Query(`SELECT id, username, email, created, updated FROM users`).Iter()
	for {
		p := User{}
		if iq.Scan(&p.Id, &p.Username, &p.Email, &p.Created, &p.Updated) {
			ulist = append(ulist, p)
		} else {
			break
		}
	}
	if err := iq.Close(); err != nil {
		log.Fatal(err)
	}

	return ulist
}

func GetUser(cass *gocql.Session, id string, w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Getting %v\n", id)
	q := cass.Query(`SELECT id, username, email, created, updated FROM users WHERE id=?`, id)
	u := User{}
	q.Scan(&u.Id, &u.Username, &u.Email, &u.Created, &u.Updated)
	fmt.Printf("Got: %v\n", u)
	json, err := json.Marshal(u)
	if err != nil {
		fmt.Printf("Failed to fetch user %s: %s\n", id, err)
	}
	w.Write(json)
}

func NewUser(cass *gocql.Session, id string, w http.ResponseWriter, r *http.Request) {
	u := User{}
	body := r.FormValue("user")
	err := json.Unmarshal([]byte(body), &u)
	if err == nil {
		updated := uuid.TimeUUID()
		q := cass.Query(`INSERT INTO users (id, username, email, created, updated) VALUES (?, ?, ?, ?, ?) IF NOT EXISTS`)
		q.Scan(&u.Id, &u.Username, &u.Email, &u.Created, updated)

		// really should do a read and return that, but for now return what we were given
		json, _ := json.Marshal(u)
		w.Write(json)
	} else {
		w.Write([]byte(`{"error": "JSON parsing failed."}`))
	}
}

func UpdateUser(cass *gocql.Session, id string, w http.ResponseWriter, r *http.Request) {
	u := User{}
	body := r.FormValue("user")
	err := json.Unmarshal([]byte(body), &u)
	if err == nil {
		newUpdated := uuid.TimeUUID()
		q := cass.Query(`UPDATE users username=?, email=?, updated=? WHERE id=? IF updated=?`)
		q.Scan(&u.Username, &u.Email, newUpdated, &u.Id, &u.Updated)
		json, _ := json.Marshal(u) // bug
		w.Write(json)
	} else {
		w.Write([]byte(`{"error": "JSON parsing failed."}`))
	}
}

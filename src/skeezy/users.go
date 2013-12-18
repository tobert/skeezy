package skeezy

import (
	"encoding/json"
	"log"
	"net/http"
	"tux21b.org/v1/gocql"
)

func ListUsers(cass *gocql.Session, w http.ResponseWriter, r *http.Request) []User {
	ulist := make([]User, 1)

	iq := cass.Query(`SELECT id, username, email, created FROM users`).Iter()
	for {
		p := User{}
		if iq.Scan(&p.Id, &p.Username, &p.Email, &p.Created) {
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
	q := cass.Query(`SELECT id, username, email, created FROM users WHERE id=?`, id)
	u := User{}
	q.Scan(&u.Id, &u.Username, &u.Email, &u.Created)
	json, _ := json.Marshal(u)
	w.Write(json)
}

func NewUser(cass *gocql.Session, id string, w http.ResponseWriter, r *http.Request) {
	u := User{}
	body := r.FormValue("user")
	err := json.Unmarshal([]byte(body), &u)
	if err == nil {
		q := cass.Query(`INSERT INTO users (id, username, email, created) VALUES (?, ?, ?, ?) IF NOT EXISTS`)
		q.Scan(&u.Id, &u.Username, &u.Email, &u.Created)
		json, _ := json.Marshal(u)
		w.Write(json)
	} else {
		w.Write([]byte(`{"error": "JSON parsing failed."}`))
	}
}

func UpdateUser(cass *gocql.Session, id string, w http.ResponseWriter, r *http.Request) {
	u := User{}
	q := cass.Query(`INSERT INTO users (id, username, email, created) VALUES (?, ?, ?, ?) IF NOT EXISTS`)
	q.Scan(&u.Id, &u.Username, &u.Email, &u.Created)
	json, _ := json.Marshal(u)
	w.Write(json)
}

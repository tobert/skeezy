package main

import (
	"./src/skeezy"
	"encoding/json"
	"fmt"
	"net/http"
	"tux21b.org/v1/gocql"
	"tux21b.org/v1/gocql/uuid"
)

func main() {
	// connect to Cassandra
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "skeezy"
	cluster.Consistency = gocql.Quorum

	cass, err := cluster.CreateSession()
	if err != nil {
		panic(fmt.Sprintf("Error creating Cassandra session: %v", err))
	}
	defer cass.Close()

	// serve up static content
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./public/js/"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./public/css/"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./public/img/"))))
	http.Handle("/fonts/", http.StripPrefix("/fonts/", http.FileServer(http.Dir("./public/fonts/"))))

	// front page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/index.html")
	})

	// a list of post ids
	http.HandleFunc("/posts/", func(w http.ResponseWriter, r *http.Request) {
	/*
		cc := skeezy.ListPosts(cass, getId(r, "/posts/"))
		for post := range cc {
			js, _ := json.Marshal(post)
			w.Write(js)
		}
		*/
	})

	// deal with single posts, action depends on HTTP method
	http.HandleFunc("/post/", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Path[len("/post/"):]

		switch r.Method {
		case "GET":
			skeezy.GetPost(cass, id, w, r)
		case "PUT":
			skeezy.NewPost(cass, id, w, r)
		case "DELETE":
			skeezy.DelPost(cass, id, w, r)
		default:
			fmt.Fprintf(w, "Invalid method: '%s'\n", r.Method)
		}
	})

	// a list of comment ids
	http.HandleFunc("/comments/", func(w http.ResponseWriter, r *http.Request) {
		cc := skeezy.ListComments(cass, getId(r, "/comments/"))
		// return a JSON list, avoid making extra copies of the data in memory
		sep := []byte{'['}
		for comment := range cc {
			w.Write(sep)
			sep = []byte{',', '\n'}
			js, _ := json.Marshal(comment)
			w.Write(js)
		}
		w.Write([]byte{']', '\n'})
	})

	// deal with single comments, action depends on HTTP method
	http.HandleFunc("/comment/", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Path[len("/comment/"):]

		switch r.Method {
		case "GET":
			skeezy.GetComment(cass, id, w, r)
		case "PUT":
			skeezy.NewComment(cass, id, w, r)
		case "DELETE":
			skeezy.DelComment(cass, id, w, r)
		default:
			fmt.Fprintf(w, "Invalid method: '%s'\n", r.Method)
		}
	})

	// start the show
	http.ListenAndServe(":8080", nil)
}

func getId(r *http.Request, prefix string) uuid.UUID {
	idarg := r.URL.Path[len(prefix):]
	id, err := uuid.ParseUUID(idarg)
	if err != nil {
		fmt.Printf("Invalid ID: '%s'\n", idarg)
	}
	return id
}

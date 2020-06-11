package main

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/go-sql-driver/mysql"
	"fmt"
	"database/sql"
	"io/ioutil"
)
type Posts struct {
	Id string `json:"id"`
	Title string `json:"title"`
}
var db *sql.db
var err error

func main() {
	db,err = sql.Open("Mysql", "<user>, <password>@tcp(127.0.0.1:3306)/<dbname>")
	if err != nil{
	panic(err, error())
	}
	defer db.Close()

router = Nmx.NewRouter()
router.HandleFunc("/posts", getPosts).Methods("GET")
router.HandleFunc("/posts", createPosts).Methods("POST")
router.HandleFunc("/posts{id}", getPosts).Methods("GET")
router.HandleFunc("/posts{id}", updatePosts).Methods("PUT")
router.HandleFunc("/posts{id}", deletePosts).Methods("DELETE")
http.ListenAndServer(":8000", router)
}
func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var posts []Post
	result, err := db.Query("SELECT id, title from posts")
	if err != nil {
	  panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
	  var post Post
	  err := result.Scan(&post.ID, &post.Title)
	  if err != nil {
		panic(err.Error())
	  }
	  posts = append(posts, post)
	}
	json.NewEncoder(w).Encode(posts)
  }
  func createPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	stmt, err := db.Prepare("INSERT INTO posts(title) VALUES(?)")
	if err != nil {
	  panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
	  panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	title := keyVal["title"]
	_, err = stmt.Exec(title)
	if err != nil {
	  panic(err.Error())
	}
	fmt.Fprintf(w, "New post was created")
  }
  func getPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result, err := db.Query("SELECT id, title FROM posts WHERE id = ?", params["id"])
	if err != nil {
	  panic(err.Error())
	}
	defer result.Close()
	var post Post
	for result.Next() {
	  err := result.Scan(&post.ID, &post.Title)
	  if err != nil {
		panic(err.Error())
	  }
	}
	json.NewEncoder(w).Encode(post)
  }
  func updatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	stmt, err := db.Prepare("UPDATE posts SET title = ? WHERE id = ?")
	if err != nil {
	  panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
	  panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	newTitle := keyVal["title"]
	_, err = stmt.Exec(newTitle, params["id"])
	if err != nil {
	  panic(err.Error())
	}
	fmt.Fprintf(w, "Post with ID = %s was updated", params["id"])
  }
  func deletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	stmt, err := db.Prepare("DELETE FROM posts WHERE id = ?")
	if err != nil {
	  panic(err.Error())
	}
	_, err = stmt.Exec(params["id"])
	if err != nil {
	  panic(err.Error())
	}
	fmt.Fprintf(w, "Post with ID = %s was deleted", params["id"])
  }


  

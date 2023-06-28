package main

import (

	//"forum/backend/database" database will be created in backend folder
	//"forum/backend/handlers" handlers will be created in backend folder
	"fmt"
	"log"
	"net/http"
	//_ "github.com/mattn/go-sqlite3" //for later, when we open DB
)

func main() {
	path := http.FileServer(http.Dir("static")) //handling html pages
	http.Handle("/static/", http.StripPrefix("/static/", path))

	// http.HandleFunc("/", handlers.IndexHandler) //root page
	// http.HandleFunc("/login", handlers.LoginHandler)
	// http.HandleFunc("/loginauth", handlers.LoginAuthHandler)
	// http.HandleFunc("/logout", handlers.Logout)
	// http.HandleFunc("/register", handlers.RegisterHandler)
	// http.HandleFunc("/registerauth", handlers.RegisterAuthHandler)
	// http.HandleFunc("/addpost", handlers.AddPost)
	// http.HandleFunc("/createpost", handlers.CreateAPost)
	// http.HandleFunc("/filter", handlers.FilterHandle)
	// http.HandleFunc("/post", handlers.PostView)
	// http.HandleFunc("/rate", handlers.LikeDislikeHandler)

	fmt.Println("Server started at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

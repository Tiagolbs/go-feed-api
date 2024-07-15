package main

import (
	"fmt"
	"net/http"
)

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new post")
}

func (app *application) showPostHandler(w http.ResponseWriter, r *http.Request) {
	postID, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "show the details of post %d\n", postID)
}

func (app *application) listPostsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "list all posts")
}

func (app *application) editPostHandler(w http.ResponseWriter, r *http.Request) {
	postID, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "edit the details of post %d\n", postID)
}

func (app *application) deletePostHandler(w http.ResponseWriter, r *http.Request) {
	postID, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "delete post %d\n", postID)
}

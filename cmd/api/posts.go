package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/tiagolbs/go-feed-api/internal/data"
	"github.com/tiagolbs/go-feed-api/internal/validator"
)

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Content string `json:"content"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	post := &data.Post{
		Content: input.Content,
	}

	v := validator.New()
	if data.ValidatePost(v, post); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	fmt.Fprintf(w, "%+v\n", input)

}

func (app *application) showPostHandler(w http.ResponseWriter, r *http.Request) {
	postID, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	post := data.Post{
		ID:        postID,
		Content:   "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque mollis orci at erat semper, non lobortis lacus feugiat. Fusce venenatis eros eget libero dapibus pharetra.",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"post": post}, nil)
	if err != nil {
		app.logger.Println(err)
		app.serverErrorResponse(w, r, err)
	}
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

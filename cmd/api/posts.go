package main

import (
	"kodekraftr/learn-go-backend/internal/store"
	"net/http"
)

type createPostPayload struct {
	Content string   `json:"content"`
	Title   string   `json:"title"`
	Tags    []string `json:"tags"`
}

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {

	var payload createPostPayload

	if err := readJson(w, r, &payload); err != nil {
		writeJsonError(w, http.StatusBadRequest, err.Error())
		return
	}

	post := &store.Post{
		Content: payload.Content,
		Title:   payload.Title,
		UserID:  1,
		Tags:    payload.Tags,
	}

	ctx := r.Context()

	if err := app.store.Posts.Create(ctx, post); err != nil {
		writeJsonError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := writeJson(w, http.StatusCreated, post); err != nil {
		writeJsonError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

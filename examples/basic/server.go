package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/evertras/gonsen"
)

// The site prefix is important
//go:embed site/*
var siteFiles embed.FS

type server struct {
	inner *http.Server
}

func newServer() *server {
	// This server has a simple home page that lists some tasks and lets the user
	// mark them as complete

	// Data source
	repository := NewRepository()
	source := gonsen.NewSource(siteFiles)

	// Define the pages and the data types they use
	staticPageHome := gonsen.MustRenderStaticPage(source, "index.html")
	pageTaskList := gonsen.NewPage[[]Task](source, "list.html")
	pageTaskDetails := gonsen.NewPage[Task](source, "details.html")

	// Now build a simple standard mux... this could be any router framework,
	// but for simplicity we'll use the standard library here
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(staticPageHome)
	})

	assetFS, err := fs.Sub(siteFiles, "site")

	if err != nil {
		panic(err)
	}

	// Assets are all static files
	mux.Handle("/assets/", http.FileServer(http.FS(assetFS)))

	mux.HandleFunc("/task", pageTaskList.HandlerWithSource(func(r *http.Request) ([]Task, int) {
		tasks, err := repository.GetTasks()

		if err != nil {
			return nil, http.StatusInternalServerError
		}

		return tasks, http.StatusOK
	}))

	mux.HandleFunc("/task/", pageTaskDetails.HandlerWithSource(func(r *http.Request) (Task, int) {
		id, err := getTrailingID(r)

		if err != nil {
			log.Printf("Failed to get ID from request: %v", err)
			return Task{}, http.StatusBadRequest
		}

		task, err := repository.GetTask(int(id))

		if err != nil {
			// Assume it's "not found" out of laziness, but a real system should have
			// more checks than this...
			return task, http.StatusNotFound
		}

		return task, http.StatusOK
	}))

	mux.HandleFunc("/complete/", func(w http.ResponseWriter, r *http.Request) {
		id, err := getTrailingID(r)

		if err != nil {
			log.Printf("Failed to get ID from request: %v", err)
			w.WriteHeader(404)
			return
		}

		err = repository.MarkTaskComplete(id)

		if err != nil {
			log.Printf("Failed to mark task %d as complete: %v", id, err)
			w.WriteHeader(404)
		}

		log.Printf("Task %d marked as complete", id)

		w.WriteHeader(200)
	})

	return &server{
		inner: &http.Server{
			Addr:    "127.0.0.1:8080",
			Handler: mux,
		},
	}
}

func (s *server) ListenAndServe() error {
	log.Printf("Listening at http://%s", s.inner.Addr)
	return s.inner.ListenAndServe()
}

func getTrailingID(r *http.Request) (int, error) {
	// Getting this ID would be a bit easier with a more full-featured router,
	// but easy enough for demo purposes...
	splitPath := strings.Split(r.URL.Path, "/")
	idStr := splitPath[len(splitPath)-1]
	id, err := strconv.ParseInt(idStr, 10, 32)

	if err != nil {
		return 0, fmt.Errorf("failed to parse ID from fragment %q from path: %s", idStr, r.URL.Path)
	}

	return int(id), err
}

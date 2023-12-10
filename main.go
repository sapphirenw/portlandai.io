package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sapphirenw/gotmpl"
	"github.com/sapphirenw/portlandai.io/src/routes"
)

func main() {
	if err := gotmpl.XT.ParseDir("templates/", []string{".html"}); err != nil {
		fmt.Printf("There was an issue parsing the templates: %s\n", err)
		os.Exit(1)
	}

	r := chi.NewRouter()

	// middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Compress(5, "text/html", "text/css", "text/javascript"))
	r.Use(middleware.RedirectSlashes)

	// handle static files
	fsPublic := http.FileServer(http.Dir("./public"))
	r.Handle("/public/*", http.StripPrefix("/public", fsPublic))

	// static assets
	r.Get("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		file, err := os.ReadFile("./public/images/favicon.ico")
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(404)
			w.Write([]byte("Could not find"))
			return
		}
		w.Header().Set("Content-Type", "image/x-icon")
		w.Write(file)
	})
	r.Get("/apple-touch-icon.png", func(w http.ResponseWriter, r *http.Request) {
		file, err := os.ReadFile("./public/images/apple-touch-icon.png")
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(404)
			w.Write([]byte("Could not find"))
			return
		}
		w.Header().Set("Content-Type", "image/png")
		w.Write(file)
	})

	// routes
	// r.NotFound(routes.NotFound)
	// r.Get("/error", routes.Error)
	r.Get("/", routes.Index)
	// r.Get("/search", routes.Search)

	fmt.Println("Listening on :3000")
	fmt.Println(http.ListenAndServe(":3000", r))
}

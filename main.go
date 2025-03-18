package main

import (
	"embed"
	"net/http"

	"github.com/a-h/templ"
)

//go:embed static/*
var static embed.FS

//go:generate npx --yes tailwindcss -i ./static/css/style.css -o ./static/css/tailwind.css --minify
//go:generate templ generate

func main() {
	server := new(http.Server)
	server.Addr = ":8080"

	pagesHandler := http.NewServeMux()

	pagesHandler.Handle("/", templ.Handler(templ.NopComponent))
	pagesHandler.Handle("/static/", http.FileServer(http.FS(static)))

	server.Handler = pagesHandler

	_ = server.ListenAndServe()
}

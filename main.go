package main

import (
	"embed"
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/dominicmeyer/Tanuki-Ledger/internal/components"
)

//go:embed static/*
var static embed.FS

//go:generate npx --yes tailwindcss -i ./static/css/style.css -o ./static/css/tailwind.css --minify
//go:generate templ generate

func main() {
	timeoutSeconds := 3

	server := new(http.Server)
	server.Addr = ":8080"
	server.ReadHeaderTimeout = time.Duration(timeoutSeconds) * time.Second

	pagesHandler := http.NewServeMux()

	pagesHandler.Handle("/", templ.Handler(components.Index()))
	pagesHandler.Handle("/static/", http.FileServer(http.FS(static)))

	server.Handler = pagesHandler

	_ = server.ListenAndServe()
}

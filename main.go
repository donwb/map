package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"goji.io/pat"

	"goji.io"
	"golang.org/x/net/context"
)

func main() {
	mux := goji.NewMux()
	port := ":" + os.Getenv("PORT")

	// static routes
	mux.HandleFuncC(pat.Get("/js/*"), func(c context.Context, w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "public/"+r.URL.Path[1:])
	})

	mux.HandleFuncC(pat.Get("/bootstrap/*"), func(c context.Context, w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "public/"+r.URL.Path[1:])
	})

	mux.HandleFuncC(pat.Get("/css/*"), func(c context.Context, w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "public/"+r.URL.Path[1:])
	})

	mux.HandleFuncC(pat.Get("/font-awesome/*"), func(c context.Context, w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "public/"+r.URL.Path[1:])
	})

	mux.HandleFuncC(pat.Get("/ico/*"), func(c context.Context, w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "public/"+r.URL.Path[1:])
	})

	mux.HandleFuncC(pat.Get("/img/*"), func(c context.Context, w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "public/"+r.URL.Path[1:])
	})

	mux.HandleFuncC(pat.Get("/"), rootRoute)
	mux.HandleFuncC(pat.Get("/ping"), pingRoute)

	fmt.Println("Ready to go...")

	http.ListenAndServe(port, mux)
}

func rootRoute(c context.Context, w http.ResponseWriter, r *http.Request) {
	filename := "public/index.html"
	body, _ := ioutil.ReadFile(filename)
	fmt.Fprintf(w, "%s", body)
}

func pingRoute(c context.Context, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", "pong")
}

func panicIf(err error) {
	if err != nil {
		log.Println(err)
		panic(err)
	}
}

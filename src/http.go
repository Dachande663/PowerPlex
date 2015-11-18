package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"path/filepath"
)

// Start a HTTP webserver
func startHttp() {

	router := httprouter.New()
	router.NotFound = get404Http

	router.GET("/", getIndexHttp)
	router.GET("/assets/*file", getAssetHttp)
	// router.GET("/apps/*file", http.FileServer(http.Dir("public")))
	router.ServeFiles("/apps/*filepath", http.Dir("public"))
	router.GET("/ws", startWebsocket)

	log.Fatal(http.ListenAndServe(":8000", router))

}

// Get system homepage
func getIndexHttp(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	serveHttpAsset(w, "resources/index.html")
}

// Get an encoded asset
func getAssetHttp(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	serveHttpAsset(w, string("resources/assets")+p.ByName("file"))
}

// 404 page
func get404Http(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	serveHttpAsset(w, "resources/404.html")
}

// Serve an http asset encoded into the app
func serveHttpAsset(w http.ResponseWriter, file string) {

	data, err := Asset(file)

	if err != nil {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Asset Not Found")
		println(err.Error())
		return
	}

	ext := filepath.Ext(file)

	switch ext {
	case ".html":
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
	case ".css":
		w.Header().Set("Content-Type", "text/css")
	case ".js":
		w.Header().Set("Content-Type", "application/javascript")
	case ".json":
		w.Header().Set("Content-Type", "application/json")
	}

	w.Write(data)

}

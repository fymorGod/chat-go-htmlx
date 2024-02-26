package main

import "net/http"

func server(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	http.ServeFile(w, r, "templates/index.html")
}

func main() {
	hub := NewHub()
	go hub.Run()

	http.HandleFunc("/", server)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serverWs(hub, w, r)
	})

	http.ListenAndServe(":8000", nil)
}

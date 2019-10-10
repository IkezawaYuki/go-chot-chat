package main

import (
	"fmt"
	"text/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
)

type templateHandler struct {
	once sync.Once
	filename string
	templ *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("view", t.filename)))
	})
	t.templ.Execute(w, nil)
}


func main(){

	files := http.FileServer(http.Dir(config.Static))
	fmt.Println(files)
	http.Handle("/static/", http.StripPrefix("/static/", files))

	http.Handle("/", &templateHandler{filename: "message.html"})
	http.Handle("/login", &templateHandler{filename: "login.html"})

	r := newRoom()
	http.Handle("/room", r)

	if err := http.ListenAndServe(":8080", nil); err != nil{
		log.Fatal("ListenAndServe:", err)
	}
}

package main

import (
	"flag"
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
	t.templ.Execute(w, r)
}


func main(){
	var addr = flag.String("addr", ":8080", "アプリケーションのアドレス")
	flag.Parse()

	files := http.FileServer(http.Dir(config.Static))
	fmt.Println(files)
	http.Handle("/static/", http.StripPrefix("/static/", files))

	http.Handle("/", &templateHandler{filename: "message.html"})
	http.Handle("/login", &templateHandler{filename: "login.html"})

	r := newRoom()
	http.Handle("/room", r)

	go r.run()

	log.Println("Webサーバーを開始します。ポート：", *addr)

	if err := http.ListenAndServe(*addr, nil); err != nil{
		log.Fatal("ListenAndServe:", err)
	}
}

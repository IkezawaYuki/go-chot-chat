package main

import (
	"flag"
	"fmt"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/facebook"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/gomniauth/providers/google"
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

const pass = "aiueo"

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
	gomniauth.SetSecurityKey(pass)
	gomniauth.WithProviders(
		facebook.New(security.ID, security.Key, "http://localhost:8080/auth/callback/facebook"),
		github.New(security.ID, security.Key,"http://localhost:8080/auth/callback/github"),
		google.New(security.ID, security.Key,"http://localhost:8080/auth/callback/google"),
		)

	files := http.FileServer(http.Dir(config.Static))
	fmt.Println(files)
	http.Handle("/static/", http.StripPrefix("/static/", files))

	//http.Handle("/", &templateHandler{filename: "message.html"})
	http.Handle("/chat", MustAuth(&templateHandler{filename:"messages.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/", loginHandler)

	r := newRoom()
	//r.tracer = trace.New(os.Stdout)

	http.Handle("/room", r)

	go r.run()

	log.Println("Webサーバーを開始します。ポート：", *addr)

	if err := http.ListenAndServe(*addr, nil); err != nil{
		log.Fatal("ListenAndServe:", err)
	}
}

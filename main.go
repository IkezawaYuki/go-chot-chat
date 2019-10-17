package main

import (
	"flag"
	"fmt"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/facebook"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/gomniauth/providers/google"
	"github.com/stretchr/objx"
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

const pass = "aiueo"

var avatars Avatar = TryAvatars{
	UseFileSystemAvatar,
	UseAuthAvatar,
	UseGravatar,
}

type templateHandler struct {
	once sync.Once
	filename string
	templ *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	t.once.Do(func() {
		fmt.Println(t.filename)
		t.templ = template.Must(template.ParseFiles(filepath.Join("view", t.filename)))
	})

	data := map[string]interface{}{
		"Host": r.Host,
	}
	if authCookie, err := r.Cookie("auth"); err == nil{
		data["UserData"] = objx.MustFromBase64(authCookie.Value)
	}

	t.templ.Execute(w, data)
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

	http.Handle("/chat", MustAuth(&templateHandler{filename:"message.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/", loginHandler)
	http.Handle("/upload", &templateHandler{filename: "upload.html"})
	http.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request){
		http.SetCookie(w, &http.Cookie{
			Name: "auth",
			Value: "",
			Path: "/",
			MaxAge: -1,
		})
		w.Header()["Location"] = []string{"/chat"}
		w.WriteHeader(http.StatusTemporaryRedirect)
	})
	http.Handle("/avatars/",
		http.StripPrefix("/avatars/",
			http.FileServer(http.Dir("./avatars"))))

	//r := newRoom(UseAuthAvatar)
	//r := newRoom(UseGravatar)
	r := newRoom(UseFileSystemAvatar)

	http.Handle("/room", r)

	go r.run()

	log.Println("Webサーバーを開始します。ポート：", *addr)

	if err := http.ListenAndServe(*addr, nil); err != nil{
		log.Fatal("ListenAndServe:", err)
	}
}

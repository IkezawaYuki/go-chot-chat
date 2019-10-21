package mail

import (
	"bytes"
	"fmt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"
	"google.golang.org/appengine/mail"
)

func Confirm(w http.ResponseWriter, r *http.Request){
	ctx := appengine.NewContext(r)
	addr := r.FormValue("email")
	url := createConfirmationURL(r)
	msg := &mail.Message{
		Sender:"example.com",
		To: []string{addr},
		Subject: "Confirm your registration",
		Body: fmt.Sprintf(confirmMessage, url),
	}
	if err := mail.Send(ctx, msg); err != nil{
		log.Errorf(ctx, "Couldn't send email: %v", err)
	}
}

func createConfirmationURL(r *http.Request)string{
	return ""
}


const confirmMessage = `
アカウント作成、ありがとうございます。
下記のリンクをクリックしてアカウントを確定させてください。

%s
`

func IncomingMail(w http.ResponseWriter, r *http.Request){
	ctx := appengine.NewContext(r)
	defer r.Body.Close()
	var b bytes.Buffer
	if _, err := b.ReadFrom(r.Body); err != nil{
		log.Errorf(ctx, "Error reading body: %v", err)
		return
	}
	log.Infof(ctx, "Recievid mail: %v", b)
}





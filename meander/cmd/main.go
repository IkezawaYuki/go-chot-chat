package cmd

import (
	"encoding/json"
	"go-chot-chat/meander"
	"net/http"
	"runtime"
)

func main(){
	runtime.GOMAXPROCS(runtime.NumCPU())
	http.HandleFunc("/journeys", func(writer http.ResponseWriter, request *http.Request) {
		respond(writer, request, meander.Journeys)
	})
	http.ListenAndServe(":8080", http.DefaultServeMux)
}

func respond(w http.ResponseWriter, r *http.Request, data []interface{}) error{
	return json.NewEncoder(w).Encode(data)
}
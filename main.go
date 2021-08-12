package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Resp struct {
	Status bool     `json:status`
	Fruits []string `json:fruits`
}

func ListUser(w http.ResponseWriter, r *http.Request) {
	f := []string{"lemon", "melon", "apple"}

	rep, err := json.Marshal(Resp{
		Status: true,
		Fruits: f,
	})
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-type", "Application/json")
	_, err = w.Write(rep)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
	return
}

func Hola(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/chao" {
		w.Write([]byte("'{'status': false, 'hola': 'holaaaaaaaaaaaaaaaaaaa', 'foo': 'aaaaaaa'}'"))
		w.WriteHeader(200)
		return
	}

	if r.URL.Path == "/foo" {
		w.Write([]byte("'{'status': false, 'hola': 'wwwww', 'foo': 'wwwww'}'"))
		w.WriteHeader(200)
		return
	}

	w.Write([]byte("'{'status': true, 'hola': 'chao'}'"))
	w.WriteHeader(200)
	return
}

func main() {
	http.HandleFunc("/", ListUser)
	http.HandleFunc("/hola", Hola)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil))
}

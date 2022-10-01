package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func main() {
	viper.SetDefault("PORT", "80")
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "/.")
	})
	r.HandleFunc("/hello/{name}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fmt.Fprintf(w, "hello %s", vars["name"])
	})
	log.Println("listening on :" + viper.GetString("PORT"))
	log.Fatal(http.ListenAndServe(":"+viper.GetString("PORT"), r))
}

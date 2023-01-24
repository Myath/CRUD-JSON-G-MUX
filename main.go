package main

import (
	"CRUD-JSON-WITH-GORILLA-MUX/handler"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func main() {
	config := viper.NewWithOptions(
		viper.EnvKeyReplacer(
			strings.NewReplacer(".", "_"),
		),
	)
	config.SetConfigFile("env/config")
    config.SetConfigType("ini")
	config.AutomaticEnv()

	if err := config.ReadInConfig(); err != nil{
		log.Fatalf("error loading configuration: %v", err)
	}

	p := config.GetInt("server.port")

	r := mux.NewRouter()

	r.HandleFunc("/students", handler.StudentsList).Methods(http.MethodGet)

	r.HandleFunc("/students/create", handler.CreateStudent).Methods(http.MethodGet)

	r.HandleFunc("/students/store", handler.StudentStore).Methods(http.MethodPost)

	r.HandleFunc("/student/{id:[0-9]+}/edit", handler.StudentEdit).Methods(http.MethodGet)

	r.HandleFunc("/student/{id:[0-9]+}/update", handler.StudentUpdate).Methods(http.MethodPost)

	r.HandleFunc("/student/{id:[0-9]+}/delete", handler.DeleteStudent).Methods(http.MethodGet)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", p), r); err != nil {
		log.Fatalf("%#v", err)
	}
}

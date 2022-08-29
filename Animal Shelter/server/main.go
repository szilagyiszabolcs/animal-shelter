package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func SendResponse(w http.ResponseWriter, i any, wrapper ...string) {
	data, err := json.Marshal(i)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(wrapper) > 0 {
		data = append([]byte("{\""+wrapper[0]+"\":"), data...)
		data = append(data, []byte("}")...)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func DecodeRequest(w http.ResponseWriter, r *http.Request, i any) bool {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(i)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return false
	}
	return true
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql", "root:cica123@tcp(localhost:3306)/ANIMAL_SHELTER")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MySQL")
	defer db.Close()

	mux := mux.NewRouter()
	header := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "X-Content-Type-Options"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PATCH", "DELETE", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	mux.HandleFunc("/animals", Controller_Animals).Methods("GET", "POST")
	mux.HandleFunc("/animals/{id:[0-9]+}", Controller_Animals_Id).Methods("GET", "PATCH", "DELETE")

	mux.HandleFunc("/shelters", Controller_Shelters).Methods("GET", "POST")
	mux.HandleFunc("/shelters/{id:[0-9]+}", Controller_Shelters_Id).Methods("GET", "PATCH", "DELETE")

	mux.HandleFunc("/species", Controller_Species).Methods("GET", "POST")

	http.ListenAndServe(":7777", handlers.CORS(header, methods, origins)(mux))
}

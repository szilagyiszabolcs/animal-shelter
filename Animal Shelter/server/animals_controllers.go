package main

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Animal struct {
	Id          int
	ShelterId   int
	SpeciesId   int
	Name        string
	Breed       string
	Gender      string
	Age         int
	Description string
	Height      int
	Weight      float32
}

func Controller_Animals(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		result, err := GetAnimals()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		SendResponse(w, result, "animals")

	case http.MethodPost:
		var animal Animal
		if !DecodeRequest(w, r, &animal) {
			return
		}

		err := AddAnimal(animal)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		SendResponse(w, struct{}{})
	}
}

func Controller_Animals_Id(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	switch r.Method {
	case http.MethodGet:
		result, err := GetAnimal(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		SendResponse(w, result, "animal")

	case http.MethodPatch:
		var animal Animal
		if !DecodeRequest(w, r, &animal) {
			return
		}

		err := UpdateAnimal(animal, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		SendResponse(w, struct{}{})

	case http.MethodDelete:
		err := DeleteAnimal(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		SendResponse(w, struct{}{})
	}
}

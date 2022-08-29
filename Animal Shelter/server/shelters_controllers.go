package main

import "net/http"

type Shelter struct {
	Id           int
	Name         string
	Zip          int
	City         string
	Street       string
	StreetNumber int
	Email        string
	PhoneNumber  int
	Webpage      string
}

func Controller_Shelters(w http.ResponseWriter, r *http.Request) {

}

func Controller_Shelters_Id(w http.ResponseWriter, r *http.Request) {

}

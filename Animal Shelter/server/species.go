package main

import "net/http"

type Species struct {
	Id   int
	Name string
}

type SuggestedSpecies struct {
	Name string
}

func Controller_Species(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		result, err := GetSpecies()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		SendResponse(w, result, "species")

	case http.MethodPost:
		var suggestedSpecies SuggestedSpecies
		if !DecodeRequest(w, r, &suggestedSpecies) {
			return
		}

		err := SuggestSpecies(suggestedSpecies)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		SendResponse(w, struct{}{})
	}
}

func GetSpecies() ([]Species, error) {
	result := make([]Species, 0)
	rows, err := db.Query("SELECT * FROM SPECIES")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var species Species
	for rows.Next() {
		rows.Scan(&species.Id, &species.Name)
		result = append(result, species)
	}

	return result, nil
}

func SuggestSpecies(name SuggestedSpecies) error {
	_, err := db.Exec("INSERT INTO SUGGESTED_SPECIES VALUES(?)", name.Name)
	return err
}

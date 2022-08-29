package main

import "errors"

func GetAnimals() ([]Animal, error) {
	result := make([]Animal, 0)

	rows, err := db.Query("SELECT `Id`, `ShelterId`, `SpeciesId`, `Name`, `Breed`, `Gender`, `Age`, `Description`, `Height`, `Weight` FROM ANIMALS")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var animal Animal
	for rows.Next() {
		rows.Scan(&animal.Id, &animal.ShelterId, &animal.SpeciesId, &animal.Name, &animal.Breed, &animal.Gender, &animal.Age, &animal.Description, &animal.Height, &animal.Weight)
		result = append(result, animal)
	}

	return result, nil
}

func AddAnimal(animal Animal) error {
	_, err := db.Exec("INSERT INTO ANIMALS (`ShelterId`, `SpeciesId`, `Name`, `Breed`, `Gender`, `Age`, `Description`, `Height`, `Weight`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)", animal.ShelterId, animal.SpeciesId, animal.Name, animal.Breed, animal.Gender, animal.Age, animal.Description, animal.Height, animal.Weight)
	return err
}

func GetAnimal(id int) (*Animal, error) {
	animal := &Animal{}

	rows, err := db.Query("SELECT `Id`, `ShelterId`, `SpeciesId`, `Name`, `Breed`, `Gender`, `Age`, `Description`, `Height`, `Weight` FROM animals WHERE animals.`Id` = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		rows.Scan(&animal.Id, &animal.ShelterId, &animal.SpeciesId, &animal.Name, &animal.Breed, &animal.Gender, &animal.Age, &animal.Description, &animal.Height, &animal.Weight)
	} else {
		return nil, errors.New("item not found")
	}

	return animal, nil
}

func UpdateAnimal(animal Animal, id int) error {
	_, err := db.Exec("UPDATE ANIMALS SET `Name`=?, `Breed`=?, `Gender`=?, `Age`=?,`Description`=?, `Height`=?, `Weight`=? WHERE `Id`=?", animal.Name, animal.Breed, animal.Gender, animal.Age, animal.Description, animal.Height, animal.Weight, id)
	return err
}

func DeleteAnimal(id int) error {
	_, err := db.Exec("DELETE FROM ANIMALS WHERE `Id`=?", id)
	return err
}

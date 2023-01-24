package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Note use of StructTag https://pkg.go.dev/reflect#StructTag which allow metadata to be attached
// to a field. Used here to provide information to the unmarshal opertion on how to transform field.
type Data struct {
	Status  string
	Code    int
	Total   int
	Persons []Person `json:"data"`
}

type Person struct {
	Id        int
	FirstName string
	LastName  string
	Gender    string
	Email     string
	Phone     string
	Address   Address
}

type Address struct {
	Id         int
	Street     string
	StreetName string
	City       string
	Country    string
	CountyCode string `json:"county_code"`
}

func main() {

	resp, error := http.Get("https://fakerapi.it/api/v1/persons?_quantity=3&_birthday_start=2005-01-01")

	if error != nil {
		log.Fatal("Error while making request: ", error)
	}

	defer resp.Body.Close()

	bodyBytes, error := io.ReadAll(resp.Body)

	if error != nil {
		log.Fatal("Error reading response: ", error)
	}

	var data Data
	error = json.Unmarshal(bodyBytes, &data)

	if error != nil {
		log.Fatal("Error unmarshalling response: ", error)
	}

	fmt.Println(string(bodyBytes))
	fmt.Println(data)

	fmt.Println("Response reported", data.Total, "people returned. A total of", len(data.Persons), "people were unmarshalled.")
}

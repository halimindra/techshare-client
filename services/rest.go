package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"orami.com/techshare/models"
)

func PrintPersonREST(client *http.Client, serverAddr string, id int64) {
	url := fmt.Sprintf("%v/people/%d", serverAddr, id)
	request, err := http.NewRequest("GET", url, bytes.NewBuffer([]byte{}))
	if err != nil {
		log.Fatal(err)
	}

	// Set request header
	request.Header = map[string][]string{
		"Content-Type": {"application/json"},
	}

	// Do client request
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	// Decode
	var person models.Person
	err = json.NewDecoder(response.Body).Decode(&person)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("-------------- GetPerson --------------")
	log.Printf("%+v\n", person)
}

func PrintPeopleREST(client *http.Client, serverAddr string, limit int64) {
	url := fmt.Sprintf("%v/people?limit=%d", serverAddr, limit)
	request, err := http.NewRequest("GET", url, bytes.NewBuffer([]byte{}))
	if err != nil {
		log.Fatal(err)
	}

	// Set request header
	request.Header = map[string][]string{
		"Content-Type": {"application/json"},
	}

	// Do client request
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	// Decode
	var people []models.Person
	err = json.NewDecoder(response.Body).Decode(&people)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("-------------- List People --------------")
	for _, person := range people {
		log.Printf("%+v\n", person)
	}
}

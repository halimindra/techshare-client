package main

type Person struct {
	Address struct {
		City    string `json:"city"`
		Country string `json:"country"`
	} `json:"address"`
	ID int `json:"id"`
}

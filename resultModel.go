package main

type Result struct {
	Name string `json:"name"`
	Location string `json:"location"`
	Email string `json:"email"`
	Age int	`json:"age"`
	Phone string `json:"phone1"`
	Cell string	`json:"phone2"`
	Pictures []string `json:"pictures"`
}
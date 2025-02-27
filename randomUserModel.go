package main

import "time"

type APIResponse struct {
	Results []ResponseBody
}

type ResponseBody struct {
	Name Name
	Location Location
	Email string
	Dob Dob
	Phone string
	Cell string
	Picture Picture
}

type Name struct {
	Title string
	First string
	Last string
}

type Location struct {
	Street Street
	City string
	State string
	Country string
}

type Street struct {
	Number int
	Name string
}

type Dob struct {
	Date time.Time
	Age int
}

type Picture struct {
	Large string
	Medium string
	Thumbnail string
}
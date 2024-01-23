package main

type medicine struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	ExpirationDate string `json:"expirationDate"`
	Total          int    `json:"total"`
}

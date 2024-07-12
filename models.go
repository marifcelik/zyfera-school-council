package main

type grade struct {
	Code  string `json:"code"`
	Value int    `json:"value"`
}

type Student struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Surname string  `json:"surname"`
	Number  string  `json:"stdNumber"`
	Grades  []grade `json:"grades"`
}

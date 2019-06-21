package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const BaseUrl = "https://swapi.co/api/"

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/people/", getPerson)
	http.HandleFunc("/people", getPeople)
	fmt.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

type Planet struct {
	Name       string `json:"name"`
	Terrain    string `json:"terrain"`
	Population string `json:"population"`
}

type Person struct {
	Name         string `json:"name"`
	HomeworldUrl string `json:"homeworld"`
	Homeworld    Planet
}

type AllPeople struct {
	People []Person `json:"results"`
}

func (p *Person) getHome() {
	resp, err := http.Get(p.HomeworldUrl)
	if err != nil {
		fmt.Println(err)
	}

	bytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bytes, &p.Homeworld)
}

func getPerson(w http.ResponseWriter, r *http.Request) {
	var p Person
	id := strings.TrimPrefix(r.URL.Path, "/people/")
	url := fmt.Sprintf("https://swapi.co/api/people/%s", id)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	bytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bytes, &p)
	p.getHome()
	fmt.Println(p)
}

func getPeople(w http.ResponseWriter, r *http.Request) {
	var p AllPeople
	res, err := http.Get(BaseUrl + "people")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Failed to request people")
		return
	}

	bytes, err := ioutil.ReadAll(res.Body)
	json.Unmarshal(bytes, &p)

	for _, pers := range p.People {
		pers.getHome()
		fmt.Println(pers)
	}

	fmt.Println(p)
}

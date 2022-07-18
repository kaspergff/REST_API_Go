package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

var users []User

func main() {
	// load JSON
	content,err := ioutil.ReadFile("data.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	
	err2 := json.Unmarshal(content, &users)
	if err2 != nil {
		fmt.Println("ERROR Unmarshal")
		fmt.Println(err2.Error())
	}

	// setup router
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/user", createUser).Methods("POST")
	router.HandleFunc("/user", getOneUser).Methods("GET")
	router.HandleFunc("/users", getAllUsers).Methods("GET")


	
	log.Fatal(http.ListenAndServe(":8080", router))

}

type User struct {
	ID         		int `json:"ID"`
	Name       		string `json:"name"`
	Phone 		 		string `json:"phone"`
	Country 	 		string `json:"country"`
	Alphanumeric 	string `json:"alphanumeric"`
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "ERROR Creating a new user")
	}
	
	json.Unmarshal(reqBody, &newUser)
	users = append(users, newUser)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newUser)
}

func getOneUser(w http.ResponseWriter, r *http.Request){
	userID := mux.Vars(r)["ID"]
	for _, singleUser := range users {
		t := strconv.Itoa(singleUser.ID)
		if t == userID {
			json.NewEncoder(w).Encode(singleUser)
		}
	}
}



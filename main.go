package main
import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type Car struct {
	ID string `json:"id"`
	Make string `json:"make"`
	Model string `json:"model"`
	Year int `json:"year"`
}

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "root:nanak908@tcp(127.0.0.1:3306)/homedepot")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	router := mux.NewRouter()

	router.HandleFunc("/cars", getCars).Methods("GET")
	router.HandleFunc("/cars/{id}", getCar).Methods("GET")
	router.HandleFunc("/cars/{id}", createCar).Methods("POST")

	http.ListenAndServe(":8000", router)
}

func getCars(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var cars []Car
	result, err := db.Query("SELECT id, make, model, year FROM cars")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var car Car
		err := result.Scan(&car.ID, &car.Make, &car.Model, &car.Year)
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println(car)
		cars = append(cars, car)
	}
	log.Println("Received all cars")
	json.NewEncoder(w).Encode(cars)
}

func createCar(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare("INSERT INTO cars(id, make, model, year) VALUES(?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		log.Fatal(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	makeField := keyVal["make"]
	model := keyVal["model"]
	year := keyVal["year"]
	_, err = stmt.Exec(id, makeField, model, year)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Created car")
	fmt.Fprintf(w, "New car was created")
}

func getCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result, err := db.Query("SELECT make, model, year FROM cars WHERE id = ?", params["id"])
	if err != nil {
		log.Fatal(err.Error())
	}
	defer result.Close()
	var car Car
	for result.Next() {
		err := result.Scan(&car.Make, &car.Model, &car.Year)
		if err != nil {
			log.Fatal(err.Error())
		}
		car.ID = params["id"]
	}
	log.Println("Received Car")
	json.NewEncoder(w).Encode(car)
}

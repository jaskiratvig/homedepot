package main

import (
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"net/http"
	"strings"
	"testing"
)

func TestCreateCar(t *testing.T) {
	url := "http://localhost:8000/cars/1234"

	var reader io.Reader
	reader = strings.NewReader(`{"make": "Lexus", "model": "Dancing", "year": 2015}`) //Convert string to reader

	request, err := http.NewRequest("POST", url, reader)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err) //Something is wrong while sending request
	}

	assert.Equal(t, 200, res.StatusCode)
}

func TestGetCar(t *testing.T) {
	url := "http://localhost:8000/cars/1234"

	var reader io.Reader
	reader = strings.NewReader("") //Convert string to reader

	request, err := http.NewRequest("GET", url, reader)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err) //Something is wrong while sending request
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf(string(bodyBytes))

	assert.Equal(t, 200, res.StatusCode)
	assert.NotEmpty(t, string(bodyBytes))
}

func TestGetCars(t *testing.T) {
	url := "http://localhost:8000/cars"

	var reader io.Reader
	reader = strings.NewReader("") //Convert string to reader

	request, err := http.NewRequest("GET", url, reader)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err) //Something is wrong while sending request
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf(string(bodyBytes))

	assert.Equal(t, 200, res.StatusCode)
	assert.NotEmpty(t, string(bodyBytes))
}

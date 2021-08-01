package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestDoubleHandler(t *testing.T) {

	testTable := []struct{
		name string
		value string
		double int
		err string
	}{
		{name:"Double of 5", value:"5", double:10},
		{name:"Missing value", value:"", err:"Missing value"},
		{name:"Not a number", value:"x", err:"Not a number"},
	}

	for _, testCase := range testTable {

		t.Run(testCase.name, func(t *testing.T){

		//HERE YOU SHOULD CHANGE THE VALUE THAT YOU WANT TO TEST double?v=<ValueToTest> Default=5
		request, err := http.NewRequest("GET", "http://localhost:8080/double?v="+testCase.value, nil)

		if err != nil {
			t.Fatalf("Could not create http request: %v", err)
		}

		recorder := httptest.NewRecorder()

		doubleHandler(recorder, request)

		response := recorder.Result()
		defer response.Body.Close()

		result, err := ioutil.ReadAll(response.Body)

		if err != nil {

			t.Fatalf("Could not read response: %v", err)

		}

		if testCase.err != "" {

			if response.StatusCode != 400 {

				t.Errorf("Expected status BadRequest, got; %v", response.Status)

			}

			if msg := string(bytes.TrimSpace(result)); msg != testCase.err {
				t.Errorf("Expected message %q; got %q", testCase.err, msg)
			}

			return

		}

		if response.StatusCode != http.StatusOK {

			t.Errorf("Expected status 200 (OK), got %v",response.StatusCode)

		}



		data, err := strconv.Atoi(string(bytes.TrimSpace(result)))

		if err != nil {

			t.Errorf("Expected integer, got %s", result)

		}

		//Here we test the double of given value. Change this if you change the default value. Default = 10.
		if data != 10 {

			t.Fatalf("Expected double of 5 (10); got %v", data)

		}

		})
	}
}

func TestRouting (t *testing.T) {

	server := httptest.NewServer(handler())
	defer server.Close()

	response, err := http.Get(fmt.Sprintf("%s/double?v=5", server.URL))

	defer response.Body.Close()

	result, err := ioutil.ReadAll(response.Body)

	if err != nil {

		t.Fatalf("Could not read response: %v", err)

	}

	if err != nil {

		t.Fatalf("Could not set GET request: %v", err)

	}

	if response.StatusCode != http.StatusOK {

		t.Errorf("Expected status 200 (OK), got %v",response.StatusCode)

	}



	data, err := strconv.Atoi(string(bytes.TrimSpace(result)))

	if err != nil {

		t.Errorf("Expected integer, got %s", result)

	}

	//Here we test the double of given value. Change this if you change the default value. Default = 10.
	if data != 10 {

		t.Fatalf("Expected double of 5 (10); got %v", data)

	}



}


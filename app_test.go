package customerAPI

import (
	"bytes"
	"customerAPI/models"
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestGetUsers(t *testing.T) {
	successfulGet := func(t *testing.T) {
		var customers []*models.Customer
		client := http.Client{}
		req, err := http.NewRequest("GET", TestAPIUrl, nil)
		if err != nil {
			t.Fatal("failed to create http request:", err)
		}
		req.SetBasicAuth("username", "password")

		resp, err := client.Do(req)
		if err != nil {
			t.Fatal("get request failed:", err)
		}
		if resp.StatusCode >= 400 {
			t.Fatal("expected non-error status code, got:", resp.StatusCode)
		}
		dec := json.NewDecoder(resp.Body)
		err = dec.Decode(&customers)
		if err != nil {
			t.Fatal("failed to decode json:", err)
		}
		if len(customers) == 0 {
			t.Fatal("expected non-empty response")
		}
	}
	failureNoAuth := func(t *testing.T) {
		expectedCode := 403
		client := http.Client{}
		req, err := http.NewRequest("GET", TestAPIUrl, nil)
		if err != nil {
			t.Fatal("failed to create http request:", err)
		}
		resp, err := client.Do(req)
		if err != nil {
			t.Fatal("get request failed:", err)
		}
		if resp.StatusCode != expectedCode {
			t.Fatalf("no authentication request failed, expected: %d, got: %d", expectedCode, resp.StatusCode)
		}
	}
	t.Run("Successful retrieve users", successfulGet)
	t.Run("Unauthorized when no auth", failureNoAuth)
}

func TestGetUser(t *testing.T) {
	successfulGet := func(t *testing.T) {
		expected := &models.Customer{
			FirstName:     "Kevin",
			LastName:      "Hart",
			Email:         "funny@aol.com",
			Company:       "SNL",
			TermsAccepted: boolPointer(false),
		}
		var customer *models.Customer
		client := http.Client{}
		req, err := http.NewRequest("GET", TestAPIUrl+"/2", nil)
		if err != nil {
			t.Fatal("failed to create http request:", err)
		}
		req.SetBasicAuth("username", "password")

		resp, err := client.Do(req)
		if err != nil {
			t.Fatal("get request failed:", err)
		}
		if resp.StatusCode >= 400 {
			t.Fatal("expected non-error status code, got:", resp.StatusCode)
		}
		dec := json.NewDecoder(resp.Body)
		err = dec.Decode(&customer)
		if err != nil {
			t.Fatal("failed to decode json:", err)
		}

		if !reflect.DeepEqual(expected, customer) {
			t.Fatalf("users returned unexpected data, expected: %v, got: %v", expected, customer)
		}
	}
	failureNoAuth := func(t *testing.T) {
		expectedCode := 403
		client := http.Client{}
		req, err := http.NewRequest("GET", TestAPIUrl+"/2", nil)
		if err != nil {
			t.Fatal("failed to create http request:", err)
		}
		resp, err := client.Do(req)
		if err != nil {
			t.Fatal("get request failed:", err)
		}
		if resp.StatusCode != expectedCode {
			t.Fatalf("no authentication request failed, expected: %d, got: %d", expectedCode, resp.StatusCode)
		}
	}
	failureNotFound := func(t *testing.T) {
		expectedCode := 404
		client := http.Client{}
		req, err := http.NewRequest("GET", TestAPIUrl+"/10000", nil)
		if err != nil {
			t.Fatal("failed to create http request:", err)
		}
		req.SetBasicAuth("username", "password")

		resp, err := client.Do(req)
		if err != nil {
			t.Fatal("get request failed:", err)
		}
		if resp.StatusCode != expectedCode {
			t.Fatalf("no authentication request failed, expected: %d, got: %d", expectedCode, resp.StatusCode)
		}
	}

	t.Run("Successful retrieve user", successfulGet)
	t.Run("Unauthorized when no auth", failureNoAuth)
	t.Run("Not found when id doesnt exist", failureNotFound)
}

func TestPostUser(t *testing.T) {
	successfulPost := func(t *testing.T) {
		expectedCode := 201
		data := &models.Customer{
			FirstName:     "Bill",
			LastName:      "Nye",
			Email:         "science@guy.com",
			TermsAccepted: boolPointer(true),
		}
		jsonData, err := json.Marshal(&data)
		if err != nil {
			t.Fatal("failed to parse json:", err)
		}
		buf := bytes.NewBuffer(jsonData)
		client := http.Client{}
		req, err := http.NewRequest("POST", TestAPIUrl, buf)
		if err != nil {
			t.Fatal("failed to create http request:", err)
		}
		req.SetBasicAuth("username", "password")
		resp, err := client.Do(req)
		if err != nil {
			t.Fatal("post request failed:", err)
		}
		if resp.StatusCode != expectedCode {
			t.Fatalf("adding a user failed, expected status code: %d, got: %d", expectedCode, resp.StatusCode)
		}
	}
	failureNoAuth := func(t *testing.T) {
		expectedCode := 403
		client := http.Client{}
		req, err := http.NewRequest("POST", TestAPIUrl, nil)
		if err != nil {
			t.Fatal("failed to create http request:", err)
		}
		resp, err := client.Do(req)
		if err != nil {
			t.Fatal("post request failed:", err)
		}
		if resp.StatusCode != expectedCode {
			t.Fatalf("no authentication request failed, expected: %d, got: %d", expectedCode, resp.StatusCode)
		}
	}
	failureValidation := func(t *testing.T) {
		expectedCode := 400
		testData := []*models.Customer{
			{
				FirstName:     "NoLastName",
				Email:         "real@email.com",
				TermsAccepted: boolPointer(true),
			},
			{
				LastName:"NoFirstName",
				Email:"still@real.com",
				TermsAccepted:boolPointer(true),
			},
			{
				FirstName:"Exists",
				LastName:"As well",
				Email:"fake",
				TermsAccepted:boolPointer(true),
			},
			{
				FirstName:"Didn't read",
				LastName:"Terms and Conditions",
				Email:"valid@gmail.com",
			},
		}
		for _, data := range testData {
			jsonData, err := json.Marshal(&data)
			if err != nil {
				t.Fatal("failed to parse json:", err)
			}
			buf := bytes.NewBuffer(jsonData)
			client := http.Client{}
			req, err := http.NewRequest("POST", TestAPIUrl, buf)
			if err != nil {
				t.Fatal("failed to create http request:", err)
			}
			req.SetBasicAuth("username", "password")
			resp, err := client.Do(req)
			if err != nil {
				t.Fatal("post request failed:", err)
			}
			if resp.StatusCode != expectedCode {
				t.Fatalf("adding a user failed, expected status code: %d, got: %d", expectedCode, resp.StatusCode)
			}
		}
	}

	t.Run("Successful add user", successfulPost)
	t.Run("Unauthorized when no auth", failureNoAuth)
	t.Run("Bad request when invalid data", failureValidation)
}

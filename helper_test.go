package customerAPI

import (
	"customerAPI/models"
	"customerAPI/storage"
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
	"testing"
)

const (
	TestingPort  = 8081
	TestDataPath = "testdata/"
	TestDataset  = TestDataPath + "test_ds.json"
)

var TestAPIUrl = "http://127.0.0.1:" + strconv.Itoa(TestingPort) + "/v1/users"

func TestMain(m *testing.M) {
	db := storage.NewMapStorage()
	var data []*models.Customer
	out, err := ioutil.ReadFile(TestDataset)
	if err != nil {
		os.Exit(1)
	}
	if err := json.Unmarshal(out, &data); err != nil {
		os.Exit(2)
	}
	for _, item := range data {
		if err := db.CreateCustomer(item); err != nil {
			os.Exit(3)
		}
	}

	app := NewApp(db)
	app.Init()
	go app.Run(TestingPort)

	code := m.Run()
	os.Exit(code)
}

func boolPointer(value bool) *bool {
	return &value
}
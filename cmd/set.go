package cmd

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// GetValue returns the value of a key
func GetValue(key string) (string, error) {
	url := "http://localhost:3000/api/value/" + key

	resp, err := http.Get(url)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	res := string(body)
	return res, err
}

// CreateKV creates a new key value pair
func CreateKV(key, value string) error {
	url := "http://localhost:3000/api/set/value/"
	values := map[string]string{"key": key, "value": value}
	jsonValue, _ := json.Marshal(values)

	_, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

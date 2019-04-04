package routes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/liquidslr/storeservice/db"
)

// DBClient interface implementation for database client
var DBClient db.DbClient

type requestBody struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// GetKeyValue http handler
func GetKeyValue(w http.ResponseWriter, r *http.Request) {
	var key = mux.Vars(r)["key"]
	account, err := DBClient.GetValue(key)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(http.StatusOK)
	w.Write(account)
}

// SetKeyValue http handler
func SetKeyValue(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	req := requestBody{}
	json.Unmarshal(bodyBytes, &req)
	err = DBClient.SetValue(req.Key, req.Value)

	if err == nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Key Value created successfully"))
	}
}

// GetAllValues http handler
func GetAllValues(w http.ResponseWriter, r *http.Request) {
	pairs := DBClient.GetAll()

	values, err := json.Marshal(pairs)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(values)
}

package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	tm "github.com/buger/goterm"
)

// Subscribe to changes in key value pairs
func Subscribe() {
	value, err := Watch()
	if err != nil {
		fmt.Println("Some error occurred")
	}

	var prettyJSON bytes.Buffer
	error := json.Indent(&prettyJSON, []byte(value), "", "\t")
	if error != nil {
		log.Println("JSON parse error: ", error)
		return
	}

	tm.Clear()
	// By moving cursor to top-left position we ensure that console output
	// will be overwritten each time, instead of adding new.
	tm.MoveCursor(1, 1)
	tm.Println("Key Values", string(prettyJSON.Bytes()))
	tm.Flush() // Call it every time at the end of rendering
	time.Sleep(time.Second * 1)

}

// Watch looks for changes in key value pairs
func Watch() (string, error) {
	url := "http://localhost:3000/api/get/all/"
	resp, err := http.Get(url)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	res := string(body)
	return res, err
}

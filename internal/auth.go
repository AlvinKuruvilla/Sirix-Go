package internal

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

//Auth a struct representing the data needed for authentication
type Auth struct {
	username string
	password string
	url      string
	client   http.Client
}

func authenticate(a Auth, w http.ResponseWriter, r *http.Request, jsonByteData []byte) {
	//todo: do explicit status code checking if possible
	respBody := bytes.NewBuffer(jsonByteData)
	resp, err := http.Post(a.url, "application/json", respBody)
	if resp.StatusCode != http.StatusOK {
		log.Fatal("There was an error wheb requesting the url", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	log.Printf(sb)
}

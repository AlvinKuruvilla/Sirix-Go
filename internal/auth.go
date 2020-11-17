package internal

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

type scheme int

const (
	//HTTP usng the HTTP protocol
	HTTP scheme = iota
	//HTTPS using the HTTPS protocol
	HTTPS
)

var schemeStr = [...]string{
	"HTTPS",
	"HTTP",
}

//Auth a struct representing the data needed for authentication
type Auth struct {
	username  string
	password  string
	baseURL  string
	client    http.Client
	tokenData TokeData
	s         scheme
	host      string
}

//NewAuth creates a new Auth struct
func NewAuth(
	//todo: do we need to include certificate authority and scheme like in the rust code?
	username string,
	password string,
	baseURL string,
	client http.Client,
	s scheme,
	host string) Auth {
	emptyTokenData := TokeData{"", 0, 0, 0, 0, "", "", "", ""}
	a := Auth{
		username,
		password,
		baseURL,
		client,
		emptyTokenData,
		s,
		host}
	return a
}

func authenticate(a Auth, w http.ResponseWriter, r *http.Request, jsonByteData []byte) {
	//todo: do explicit status code checking if possible
	respBody := bytes.NewBuffer(jsonByteData)
  resp, err := http.Post(a.baseURL, "application/json", respBody)
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

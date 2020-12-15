package internal

import (
	"bytes"
	"net/http"

	"github.com/emirpasic/gods/maps/hashmap"
)

type scheme int

var headerMap = hashmap.New()

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
	baseURL   string
	client    http.Client
	tokenData TokeData
	s         scheme
	host      string
}

//NewAuth creates a new Auth struct
func NewAuth(
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
		host,
	}
	return a
}

//NOTE: right now this is synchronous
func (a Auth) Authenticate(w http.ResponseWriter, r *http.Request, jsonByteData []byte) {
	reqBody := bytes.NewBuffer(jsonByteData)
	req, err := http.NewRequest("POST", a.baseURL, reqBody)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

}
